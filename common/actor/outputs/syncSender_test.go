package outputs

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/tombenke/axon-go/common/messenger"
	messengerImpl "github.com/tombenke/axon-go/common/messenger/nats"
	"github.com/tombenke/axon-go/common/msgs"
	"github.com/tombenke/axon-go/common/msgs/orchestra"
	at "github.com/tombenke/axon-go/common/testing"
	"sync"
	"testing"
)

func TestSyncSender(t *testing.T) {
	// Connect to messaging
	m := messengerImpl.NewMessenger(messengerCfg)
	defer m.Close()

	// Use a WaitGroup to wait for the processes of the testbed to complete their mission
	wg := sync.WaitGroup{}

	// Create a trigger channel to start the test
	triggerCh := make(chan bool)

	// Start the processes of the test-bed
	doneChkCh := make(chan bool)
	reportCh, testCompletedCh, chkStoppedCh := at.ChecklistProcess(syncChecklist, doneChkCh, &wg, logger)

	doneOrchCh := make(chan bool)
	orchStoppedCh := startMockOrchestrator(t, reportCh, doneOrchCh, &wg, logger, m)

	doneRcvCh := make(chan bool)
	rcvStoppedCh := startMockMessageReceivers(getOutputsData(), reportCh, doneRcvCh, &wg, logger, m)

	doneProcCh := make(chan bool)
	outputsCh, procStoppedCh := startMockProcessor(triggerCh, reportCh, doneProcCh, &wg, logger)

	// Start the sender process
	doneSndCh := make(chan bool)
	senderStoppedCh := SyncSender(actorName, outputsCh, doneSndCh, &wg, m, logger)

	// Start testing
	logger.Infof("Send trigger to start testing")
	triggerCh <- true

	// Wait until test is completed, then stop the processes
	logger.Infof("Wait until test is completed")
	<-testCompletedCh

	logger.Infof("Stops Mock Orchestrator")
	close(doneOrchCh)
	logger.Infof("Wait Mock Orchestrator to stop")
	<-orchStoppedCh
	logger.Infof("Mock Orchestrator stopped")

	logger.Infof("Stops Sender")
	close(doneSndCh)
	logger.Infof("Wait Sender to stop")
	<-senderStoppedCh
	logger.Infof("Sender stopped")

	logger.Infof("Stops Mock Processor")
	close(doneProcCh)
	logger.Infof("Wait Mock Processor to stop")
	<-procStoppedCh
	logger.Infof("Mock Processor stopped")

	logger.Infof("Stops Stops Mock Receiver")
	close(doneRcvCh)
	logger.Infof("Wait Mock Receiver to stop")
	<-rcvStoppedCh
	logger.Infof("Mock Receiver stopped")

	logger.Infof("Stops Checklist")
	close(doneChkCh)
	logger.Infof("Wait Checklist to stop")
	<-chkStoppedCh
	logger.Infof("Checklist stopped")

	// Wait for the message to come in
	wg.Wait()
}

// startMockOrchestrator starts a standalone process that emulates the behaviour of an external orchestrator application.
// Orchestrator waits for an incoming message via the `processing-completed` messaging channel,
// then sends a trigger message to the SyncSender process via the `send-outputs` messaging channel.
// The Mock Orchestrator reports every relevant event to the Checklist process.
// Mock Orchestrator will shut down if it receives a message via the `doneCh` channel.
func startMockOrchestrator(t *testing.T, reportCh chan string, doneCh chan bool, wg *sync.WaitGroup, logger *logrus.Logger, m messenger.Messenger) chan bool {
	processingCompletedCh := make(chan []byte)
	processingCompletedSubs := m.ChanSubscribe("processing-completed", processingCompletedCh)

	sendingCompletedCh := make(chan []byte)
	sendingCompletedSubs := m.ChanSubscribe("sending-completed", sendingCompletedCh)

	orchStoppedCh := make(chan bool)

	wg.Add(1)
	go func() {
		defer processingCompletedSubs.Unsubscribe()
		defer close(processingCompletedCh)
		defer sendingCompletedSubs.Unsubscribe()
		defer close(sendingCompletedCh)
		defer close(orchStoppedCh)
		defer wg.Done()

		for {
			select {
			case <-doneCh:
				logger.Infof("MockOrchestrator shuts down.")
				return

			case messageBytes := <-processingCompletedCh:
				logger.Infof("MockOrchestrator received 'processing-completed' message.")
				// Check if the right actorName was sent in the message
				processingCompletedMsg := orchestra.NewProcessingCompletedMessage("")
				processingCompletedMsg.Decode(msgs.JSONRepresentation, messageBytes)
				assert.Equal(t, processingCompletedMsg.(*orchestra.ProcessingCompleted).Body.Data, actorName)

				logger.Infof("MockOrchestrator sends 'send-results' message.")
				sendResultsMsg := orchestra.NewSendResultsMessage()
				m.Publish("send-results", sendResultsMsg.Encode(msgs.JSONRepresentation))

			case messageBytes := <-sendingCompletedCh:
				logger.Infof("MockOrchestrator received 'sending-completed' message.")
				// Check if the right actorName was sent in the message
				sendingCompletedMsg := orchestra.NewSendingCompletedMessage("")
				sendingCompletedMsg.Decode(msgs.JSONRepresentation, messageBytes)
				assert.Equal(t, sendingCompletedMsg.(*orchestra.SendingCompleted).Body.Data, actorName)
				reportCh <- checkSendingCompleted
			}
		}
	}()

	logger.Infof("Mock Orchestrator started.")
	return orchStoppedCh
}
