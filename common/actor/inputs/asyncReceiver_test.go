package inputs

import (
	messengerImpl "github.com/tombenke/axon-go/common/messenger/nats"
	at "github.com/tombenke/axon-go/common/testing"
	"sync"
	"testing"
	"time"
)

// TestAsyncReceiverStartStop sets up the input ports, then stops.
func TestAsyncReceiverStartStop(t *testing.T) {
	// Connect to messaging
	m := messengerImpl.NewMessenger(messengerCfg)
	defer m.Close()

	// Use a WaitGroup to wait for the processes of the testbed to complete their mission
	wg := sync.WaitGroup{}

	// Create a channel to shut down the processes if needed
	doneCh := make(chan bool)

	// Start the receiver process
	AsyncReceiver(inputsCfg, doneCh, &wg, m, logger)

	// Wait until test is completed, then stop the processes
	close(doneCh)

	// Wait for the message to come in
	wg.Wait()
}

// TestReceiveInputs sets up the input ports, and gets inputs to each ports, then a receive-and-process message,
// It uses the incoming messages that it sends as the result inputs to the processor.
func TestAsyncReceiverInputs(t *testing.T) {
	// Connect to messaging
	m := messengerImpl.NewMessenger(messengerCfg)
	defer m.Close()

	// Use a WaitGroup to wait for the processes of the testbed to complete their mission
	wg := sync.WaitGroup{}

	// Create a channel to shut down the processes if needed
	doneCh := make(chan bool)

	// Start the processes of the test-bed
	reportCh, testCompletedCh := at.ChecklistProcess(checklistAsync, doneCh, &wg, logger)

	// Start the receiver process
	inputsCh := AsyncReceiver(inputsCfg, doneCh, &wg, m, logger)

	startMockProcessor(inputsCh, reportCh, doneCh, &wg, logger)

	// Give chance for observers to start before send messages through external messaging mw.
	time.Sleep(100 * time.Millisecond)

	// Start testing
	sendInputMessages(inputsCfg, syncInputs, reportCh, m, logger)

	// Wait until test is completed, then stop the processes
	<-testCompletedCh
	close(doneCh)

	// Wait for the message to come in
	wg.Wait()
}
