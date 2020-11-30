package messenger

import (
	"github.com/stretchr/testify/require"
	"sync"
	"testing"
)

var (
	testChannelDurable = "test_channel_durable"
)

// Test the Asynchronous / Observer pattern: publish/subscribe
func TestPubSubDurable(t *testing.T) {
	// Connect to NATS
	m := NewMessenger(testMessengerConfig)
	defer m.Close()

	// Use a WaitGroup to wait for the message to arrive
	wg := sync.WaitGroup{}
	wg.Add(1)

	// Subscribe to the source subject with the message processing function
	m.SubscribeDurable(testChannelDurable, func(content []byte) {
		defer wg.Done()
		require.EqualValues(t, content, testMsgContent)
	})

	// Send a message
	m.PublishDurable(testChannelDurable, testMsgContent)

	// Wait for the message to come in
	wg.Wait()
}

// Test the Asynchronous / Observer pattern: publish/subscribe
func TestPubSubDurableWithAck(t *testing.T) {
	// Connect to NATS
	m := NewMessenger(testMessengerConfig)
	defer m.Close()

	// Use a WaitGroup to wait for the message to arrive
	wg := sync.WaitGroup{}
	wg.Add(2)

	// Subscribe to the source subject with the message processing function
	m.SubscribeDurableWithAck(testChannelDurable, func(content []byte, ackCb func() error) {
		defer wg.Done()
		ackCb()
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
