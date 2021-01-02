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

	// Create a channel to shut down the processes if needed
	doneCh := make(chan bool)

	// Create a trigger channel to start the test
	triggerCh := make(chan bool)

	// Start the processes of the test-bed
	reportCh, testCompletedCh := at.ChecklistProcess(syncChecklist, doneCh, &wg, logger)
	startMockOrchestrator(t, reportCh, doneCh, &wg, logger, m)
	startMockMessageReceivers(getOutputsData(), reportCh, doneCh, &wg, logger, m)
	outputsCh := startMockProcessor(triggerCh, reportCh, doneCh, &wg, logger)

	// Start the sender process
	wg.Add(1)
	go SyncSender(actorName, outputsCh, doneCh, &wg, m, logger)

	// Start testing
	triggerCh <- true

	// Wait until test is completed, then stop the processes
	<-testCompletedCh
	close(doneCh)

	// Wait for the message to come in
	wg.Wait()
}

// startMockOrchestrator starts a standalone process that emulates the behaviour of an external orchestrator application.
// Orchestrator waits for an incoming message via the `processing-completed` messaging channel,
// then sends a trigger message to the SyncSender process via the `send-outputs` messaging channel.
// The Mock Orchestrator reports every relevant event to the Checklist process.
// Mock Orchestrator will shut down if it receives a message via the `doneCh` channel.
func startMockOrchestrator(t *testing.T, reportCh chan string, doneCh chan bool, wg *sync.WaitGroup, logger *logrus.Logger, m messenger.Messenger) {
	processingCompletedCh := make(chan []byte)
	processingCompletedSubs := m.ChanSubscribe("processing-completed", processingCompletedCh)

	sendingCompletedCh := make(chan []byte)
	sendingCompletedSubs := m.ChanSubscribe("sending-completed", sendingCompletedCh)

	wg.Add(1)
	go func() {
		defer processingCompletedSubs.Unsubscribe()
		defer close(processingCompletedCh)
		defer sendingCompletedSubs.Unsubscribe()
		defer close(sendingCompletedCh)
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
}
