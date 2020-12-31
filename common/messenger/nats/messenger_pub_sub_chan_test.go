package nats

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"github.com/tombenke/axon-go/common/messenger"
	"sync"
	"testing"
)

// Test the Asynchronous / Observer pattern: publish/subscribe
func TestPubSubChan(t *testing.T) {
	// Connect to NATS
	m := NewMessenger(testConfig)
	defer m.Close()

	// Use a WaitGroup to wait for the message to arrive
	wg := sync.WaitGroup{}
	wg.Add(1)

	// Subscribe to the source subject with the message processing function
	testSubject := "test_subject"
	testMsgContent := []byte("Some text to send...")
	numMessagesToSend := 5
	var s messenger.Subscriber
	ch := make(chan []byte)
	s = m.ChanSubscribe(testSubject, ch)

	go func() {
		defer wg.Done()
		for i := numMessagesToSend; i > 0; i-- {
			select {
			case content := <-ch:
				fmt.Printf("channel consumer got: '%v'\n", content)
				require.EqualValues(t, content, testMsgContent)
			}
		}
		s.Unsubscribe()
	}()

	// Send a message
	for i := numMessagesToSend; i > 0; i-- {
		m.Publish(testSubject, testMsgContent)
	}

	// Wait for the message to come in
	wg.Wait()
}
