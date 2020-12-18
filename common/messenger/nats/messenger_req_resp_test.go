package nats

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"sync"
	"testing"
	"time"
)

// Test the Synchronous / Consumer pattern: request/response
func TestReqResp(t *testing.T) {
	// Connect to NATS
	m := NewMessenger(testConfig)
	defer m.Close()

	// Use a WaitGroup to wait for the message to arrive
	wg := sync.WaitGroup{}
	wg.Add(1)

	// Subscribe to the source subject with the message processing function
	testSubject := "test_subject"
	testMsgContent := []byte("Some text to send...")
	testRespContent := []byte("Some text to send back as response...")
	m.Response(testSubject, func(content []byte) ([]byte, error) {
		defer wg.Done()
		require.EqualValues(t, content, testMsgContent)
		return testRespContent, nil
	})

	// Send a message
	resp, err := m.Request(testSubject, testMsgContent, 50*time.Millisecond)
	assert.Nil(t, err)
	require.EqualValues(t, resp, testRespContent)

	// Wait for the message to come in
	wg.Wait()
}

// Test the Synchronous / Consumer pattern: request/response with server side error
func TestReqRespServerErr(t *testing.T) {
	// Connect to NATS
	m := NewMessenger(testConfig)
	defer m.Close()

	// Use a WaitGroup to wait for the message to arrive
	wg := sync.WaitGroup{}
	wg.Add(1)

	// Subscribe to the source subject with the message processing function
	testSubject := "test_subject"
	testMsgContent := []byte("Some text to send...")
	testRespErr := errors.New("Server error")
	m.Response(testSubject, func(content []byte) ([]byte, error) {
		defer wg.Done()
		require.EqualValues(t, content, testMsgContent)
		return nil, testRespErr
	})

	// Send a message
	resp, err := m.Request(testSubject, testMsgContent, 50*time.Millisecond)
	assert.Nil(t, err)
	require.EqualValues(t, resp, testRespErr.Error())

	// Wait for the message to come in
	wg.Wait()
}

// Test the Synchronous / Consumer pattern: request/response with timeout error
func TestReqRespTimeoutErr(t *testing.T) {
	// Connect to NATS
	m := NewMessenger(testConfig)
	defer m.Close()

	// Use a WaitGroup to wait for the message to arrive
	wg := sync.WaitGroup{}
	wg.Add(1)

	// Subscribe to the source subject with the message processing function
	testSubject := "test_subject"
	testMsgContent := []byte("Some text to send...")
	m.Response(testSubject, func(content []byte) ([]byte, error) {
		defer wg.Done()
		require.EqualValues(t, content, testMsgContent)
		time.Sleep(100 * time.Millisecond)
		return []byte(``), nil
	})

	// Send a message
	_, err := m.Request(testSubject, testMsgContent, 50*time.Millisecond)
	assert.NotNil(t, err)
	assert.Equal(t, err, errors.New("nats: timeout"), "should be equal")

	// Wait for the message to come in
	wg.Wait()
}
