package messenger

import (
	"github.com/stretchr/testify/require"
	"sync"
	"testing"
)

var (
	testSubject = "test_subject"
)

// Test the Asynchronous / Observer pattern: publish/subscribe
func TestPubSub(t *testing.T) {
	// Connect to NATS
	m := NewMessenger(testMessengerConfig)
	defer m.Close()

	// Use a WaitGroup to wait for the message to arrive
	wg := sync.WaitGroup{}
	wg.Add(1)

	// Subscribe to the source subject with the message processing function
	m.Subscribe(testSubject, func(content []byte) {
		defer wg.Done()
		require.EqualValues(t, content, testMsgContent)
	})

	// Send a message
	m.Publish(testSubject, testMsgContent)

	// Wait for the message to come in
	wg.Wait()
}
