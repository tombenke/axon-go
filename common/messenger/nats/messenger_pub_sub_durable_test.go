package nats

import (
	"github.com/stretchr/testify/require"
	"sync"
	"testing"
)

// Test the Asynchronous / Observer pattern: publish/subscribe
func TestPubSubDurable(t *testing.T) {
	// Connect to NATS
	m := NewMessenger(testConfig)
	defer m.Close()

	// Use a WaitGroup to wait for the message to arrive
	wg := sync.WaitGroup{}
	wg.Add(1)

	// Subscribe to the source subject with the message processing function
	testChannelDurable := "test_channel_durable"
	testMsgContent := []byte("Some text to send...")
	m.SubscribeDurable(testChannelDurable, func(content []byte) {
		defer wg.Done()
		require.EqualValues(t, content, testMsgContent)
	})

	// Send a message
	err := m.PublishDurable(testChannelDurable, testMsgContent)
	require.Nil(t, err)

	// Wait for the message to come in
	wg.Wait()
}
