package nats

import (
	"github.com/stretchr/testify/require"
	"github.com/tombenke/axon-go/common/messenger"
	"sync"
	"testing"
)

// Test the Asynchronous / Observer pattern: publish/subscribe
func TestPubSub(t *testing.T) {
	// Connect to NATS
	m := NewMessenger(testConfig)
	defer m.Close()

	// Use a WaitGroup to wait for the message to arrive
	wg := sync.WaitGroup{}
	wg.Add(1)

	// Subscribe to the source subject with the message processing function
	testSubject := "test_subject"
	testMsgContent := []byte("Some text to send...")
	var s messenger.Subscriber
	s = m.Subscribe(testSubject, func(content []byte) {
		defer wg.Done()
		require.EqualValues(t, content, testMsgContent)
		err := s.Unsubscribe()
		require.Nil(t, err)
	})

	// Send a message
	err := m.Publish(testSubject, testMsgContent)
	require.Nil(t, err)

	// Wait for the message to come in
	wg.Wait()
}
