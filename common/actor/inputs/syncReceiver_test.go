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

	// Create a channel to shut down the processes if needed
	doneCh := make(chan bool)

	// Start the receiver process
	SyncReceiver(inputsCfg, doneCh, &wg, m, logger)

	// Wait until test is completed, then stop the processes
	close(doneCh)

	// Wait for the message to come in
	wg.Wait()
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

	// Create a channel to shut down the processes if needed
	doneCh := make(chan bool)

	// Start the processes of the test-bed
	reportCh, testCompletedCh := at.ChecklistProcess(checklistDefaultsOnly, doneCh, &wg, logger)

	// Create a trigger channel to start the test
	triggerOrchCh := startMockOrchestrator(reportCh, doneCh, &wg, m, logger)

	// Start the receiver process
	inputsCh := SyncReceiver(inputsCfg, doneCh, &wg, m, logger)

	startMockProcessor(inputsCh, reportCh, doneCh, &wg, logger)

	// Start testing
	time.Sleep(10 * time.Millisecond)
	triggerOrchCh <- true

	// Wait until test is completed, then stop the processes
	<-testCompletedCh
	close(doneCh)

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

	// Create a channel to shut down the processes if needed
	doneCh := make(chan bool)

	// Start the processes of the test-bed
	reportCh, testCompletedCh := at.ChecklistProcess(checklistFull, doneCh, &wg, logger)

	// Create a trigger channel to start the test
	triggerOrchCh := startMockOrchestrator(reportCh, doneCh, &wg, m, logger)

	// Start the receiver process
	inputsCh := SyncReceiver(inputsCfg, doneCh, &wg, m, logger)

	startMockProcessor(inputsCh, reportCh, doneCh, &wg, logger)

	// Start testing
	sendInputMessages(inputsCfg, syncInputs, reportCh, m, logger)
	time.Sleep(10 * time.Millisecond)
	triggerOrchCh <- true

	// Wait until test is completed, then stop the processes
	<-testCompletedCh
	close(doneCh)

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
func startMockOrchestrator(reportCh chan string, doneCh chan bool, wg *sync.WaitGroup, m messenger.Messenger, logger *logrus.Logger) chan bool {
	triggerOrchCh := make(chan bool)

	wg.Add(1)
	go func() {
		defer close(triggerOrchCh)
		defer wg.Done()

		for {
			select {
			case <-doneCh:
				logger.Infof("Mock Orchestrator shuts down.")
				return

			case <-triggerOrchCh:
				receiveAndProcessMsg := orchestra.NewReceiveAndProcessMessage(float64(1.0))
				m.Publish("receive-and-process", receiveAndProcessMsg.Encode(msgs.JSONRepresentation))
				logger.Infof("Mock Orchestrator sent 'receive-and-process' message.")
				reportCh <- checkSendReceiveAndProcess
			}
		}
	}()

	logger.Infof("Mock Orchestrator started.")
	return triggerOrchCh
}
