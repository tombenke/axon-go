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
	AsyncReceiver(asyncInputsCfg, doneCh, &wg, m, logger)

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
	// doneCh := make(chan bool)

	// Start the processes of the test-bed
	doneCheckCh := make(chan bool)
	reportCh, testCompletedCh, chkStoppedCh := at.ChecklistProcess(checklistAsync, doneCheckCh, &wg, logger)

	// Start the receiver process
	doneRcvCh := make(chan bool)
	inputsCh, rcvStoppedCh := AsyncReceiver(asyncInputsCfg, doneRcvCh, &wg, m, logger)

	doneProcCh := make(chan bool)
	procStoppedCh := startMockProcessor(inputsCh, reportCh, doneProcCh, &wg, logger)

	// Give chance for observers to start before send messages through external messaging mw.
	time.Sleep(100 * time.Millisecond)

	// Start testing
	sendInputMessages(asyncInputsCfg, asyncInputs, reportCh, m, logger)

	// Wait until test is completed, then stop the processes
	<-testCompletedCh

	logger.Infof("Stops Mock Processor")
	close(doneProcCh)
	logger.Infof("Wait Mock Processor to stop")
	<-procStoppedCh
	logger.Infof("Mock Processor stopped")

	logger.Infof("Stops Stops Receiver")
	close(doneRcvCh)
	logger.Infof("Wait Receiver to stop")
	<-rcvStoppedCh
	logger.Infof("Receiver stopped")

	logger.Infof("Stops Checklist")
	close(doneCheckCh)
	logger.Infof("Wait Checklist to stop")
	<-chkStoppedCh
	logger.Infof("Checklist stopped")

	// Wait for the message to come in
	wg.Wait()
}
