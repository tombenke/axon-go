package inputs

import (
	"github.com/sirupsen/logrus"
	"github.com/tombenke/axon-go/common/messenger"
	messengerImpl "github.com/tombenke/axon-go/common/messenger/nats"
	"github.com/tombenke/axon-go/common/msgs"
	"github.com/tombenke/axon-go/common/msgs/orchestra"
	at "github.com/tombenke/axon-go/common/testing"
	"sync"
	"testing"
	"time"
)

// TestSyncReceiverStartStop sets up the input ports, then stops.
func TestSyncReceiverStartStop(t *testing.T) {
	// Connect to messaging
	m := messengerImpl.NewMessenger(messengerCfg)
	defer m.Close()

	// Use a WaitGroup to wait for the processes of the testbed to complete their mission
	wg := sync.WaitGroup{}

	// Create a channel for the RESET
	resetCh := make(chan bool)

	// Create a channel to shut down the processes if needed
	doneCh := make(chan bool)

	// Start the receiver process
	SyncReceiver(syncInputsCfg, resetCh, doneCh, &wg, m, logger)

	// Wait until test is completed, then stop the processes
	close(doneCh)

	// Wait for the message to come in
	wg.Wait()
	close(resetCh)
}

// TestReceiveDefaultsOnly sets up the input ports, then gets a receive-and-process message,
// but receive no input messages via the ports, so it uses the default values defined to the ports.
// It sends the result inputs to the processor.
func TestSyncReceiverDefaultsOnly(t *testing.T) {
	// Connect to messaging
	m := messengerImpl.NewMessenger(messengerCfg)
	defer m.Close()

	// Use a WaitGroup to wait for the processes of the testbed to complete their mission
	wg := sync.WaitGroup{}

	// Start the processes of the test-bed
	doneCheckCh := make(chan bool)
	reportCh, testCompletedCh, chkStoppedCh := at.ChecklistProcess(checklistDefaultsOnly, doneCheckCh, &wg, logger)

	// Create a trigger channel to start the test
	doneOrchCh := make(chan bool)
	triggerOrchCh, orchStoppedCh := startMockOrchestrator(reportCh, doneOrchCh, &wg, m, logger)

	// Create a channel for the RESET
	resetCh := make(chan bool)

	// Start the receiver process
	doneRcvCh := make(chan bool)
	inputsCh, rcvStoppedCh := SyncReceiver(syncInputsCfg, resetCh, doneRcvCh, &wg, m, logger)

	doneProcCh := make(chan bool)
	procStoppedCh := startMockProcessor(inputsCh, reportCh, doneProcCh, &wg, logger)

	// Start testing
	time.Sleep(10 * time.Millisecond)
	triggerOrchCh <- true

	// Wait until test is completed, then stop the processes
	logger.Infof("Wait until test is completed")
	<-testCompletedCh

	logger.Infof("Stops Mock Orchestrator")
	close(doneOrchCh)
	logger.Infof("Wait Mock Orchestrator to stop")
	<-orchStoppedCh
	logger.Infof("Mock Orchestrator stopped")

	logger.Infof("Stops Mock Processor")
	close(doneProcCh)
	logger.Infof("Wait Mock Processor to stop")
	<-procStoppedCh
	logger.Infof("Mock Processor stopped")

	logger.Infof("Stops Receiver")
	close(doneRcvCh)
	logger.Infof("Wait Receiver to stop")
	<-rcvStoppedCh
	close(resetCh)
	logger.Infof("Receiver stopped")

	logger.Infof("Stops Checklist")
	close(doneCheckCh)
	logger.Infof("Wait Checklist to stop")
	<-chkStoppedCh
	logger.Infof("Checklist stopped")

	// Wait for the message to come in
	wg.Wait()
}

// TestReceiveInputs sets up the input ports, and gets inputs to each ports, then a receive-and-process message,
// It uses the incoming messages that it sends as the result inputs to the processor.
func TestSyncReceiverInputs(t *testing.T) {
	// Connect to messaging
	m := messengerImpl.NewMessenger(messengerCfg)
	defer m.Close()

	// Use a WaitGroup to wait for the processes of the testbed to complete their mission
	wg := sync.WaitGroup{}

	// Start the processes of the test-bed
	doneCheckCh := make(chan bool)
	reportCh, testCompletedCh, chkStoppedCh := at.ChecklistProcess(checklistFull, doneCheckCh, &wg, logger)

	// Create a trigger channel to start the test
	doneOrcCh := make(chan bool)
	triggerOrchCh, orchStoppedCh := startMockOrchestrator(reportCh, doneOrcCh, &wg, m, logger)

	// Create a channel for the RESET
	resetCh := make(chan bool)

	// Start the receiver process
	doneRcvCh := make(chan bool)
	inputsCh, rcvStoppedCh := SyncReceiver(syncInputsCfg, resetCh, doneRcvCh, &wg, m, logger)

	doneProcCh := make(chan bool)
	procStoppedCh := startMockProcessor(inputsCh, reportCh, doneProcCh, &wg, logger)

	// Start testing
	sendInputMessages(syncInputsCfg, syncInputs, reportCh, m, logger)
	time.Sleep(10 * time.Millisecond)
	triggerOrchCh <- true

	// Wait until test is completed, then stop the processes
	logger.Infof("Wait until test is completed")
	<-testCompletedCh

	logger.Infof("Stops Mock Orchestrator")
	close(doneOrcCh)
	logger.Infof("Wait Mock Orchestrator to stop")
	<-orchStoppedCh
	logger.Infof("Mock Orchestrator stopped")

	logger.Infof("Stops Mock Processor")
	close(doneProcCh)
	logger.Infof("Wait Mock Processor to stop")
	<-procStoppedCh
	logger.Infof("Mock Processor stopped")

	logger.Infof("Stops Receiver")
	close(doneRcvCh)
	logger.Infof("Wait Receiver to stop")
	<-rcvStoppedCh
	close(resetCh)
	logger.Infof("Receiver stopped")

	logger.Infof("Stops Checklist")
	close(doneCheckCh)
	logger.Infof("Wait Checklist to stop")
	<-chkStoppedCh
	logger.Infof("Checklist stopped")

	// Wait for the message to come in
	wg.Wait()
}

// TestReceiveInputsBulk sets up the input ports, and gets inputs to each ports, then a receive-and-process message,
// In some ports more than one inputs arrive, so it uses the latest one arrived to the port,
// that it sends to the processor.
func TestSyncReceiverInputsBulk(t *testing.T) {
	// TODO Implement TestSyncReceiverInputsBulk test-case
}

// startMockOrchestrator starts a standalone process that emulates
// the behaviour of an external orchestrator application.
// Orchestrator waits for a trigger to send `receive-and-process` message to the input aggregator.
// The Mock Orchestrator reports every relevant event to the Checklist process.
// Mock Orchestrator will shut down if it receives a message via the `doneCh` channel.
func startMockOrchestrator(reportCh chan string, doneCh chan bool, wg *sync.WaitGroup, m messenger.Messenger, logger *logrus.Logger) (chan bool, chan bool) {
	triggerOrchCh := make(chan bool)
	orchStoppedCh := make(chan bool)

	wg.Add(1)
	go func() {
		defer close(triggerOrchCh)
		defer close(orchStoppedCh)
		logger.Infof("Mock Orchestrator stopped.")
		defer wg.Done()

		for {
			select {
			case <-doneCh:
				logger.Infof("Mock Orchestrator shuts down.")
				return

			case <-triggerOrchCh:
				receiveAndProcessMsg := orchestra.NewReceiveAndProcessMessage(float64(1.0))
				if err := m.Publish("receive-and-process", receiveAndProcessMsg.Encode(msgs.JSONRepresentation)); err != nil {
					panic(err)
				}
				logger.Infof("Mock Orchestrator sent 'receive-and-process' message.")
				reportCh <- checkSendReceiveAndProcess
			}
		}
	}()

	logger.Infof("Mock Orchestrator started.")
	return triggerOrchCh, orchStoppedCh
}
