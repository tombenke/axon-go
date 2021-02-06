package nats

import (
	"github.com/stretchr/testify/require"
	"sync"
	"testing"
)

// Test the Asynchronous / Observer pattern: publish/subscribe
func TestPubSubDurableWithAck(t *testing.T) {
	// Connect to NATS
	m := NewMessenger(testConfig)
	defer m.Close()

	// Use a WaitGroup to wait for the message to arrive
	wg := sync.WaitGroup{}
	wg.Add(2)

	// Subscribe to the source subject with the message processing function
	testChannelDurable := "test_channel_durable"
	testMsgContent := []byte("Some text to send...")
	m.SubscribeDurableWithAck(testChannelDurable, func(content []byte, ackCb func() error) {
		defer wg.Done()
		err := ackCb()
		require.Nil(t, err)
		require.EqualValues(t, content, testMsgContent)
	})

	// Send a message
	_, err := m.PublishAsyncDurable(testChannelDurable, testMsgContent, func(guid string, ackErr error) {
		defer wg.Done()
		require.Nil(t, ackErr)
	})
	require.Nil(t, err)

	// Wait for the message to come in
	wg.Wait()
}
