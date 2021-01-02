package status

import (
	"github.com/sirupsen/logrus"
	"github.com/tombenke/axon-go/common/messenger"
	messengerImpl "github.com/tombenke/axon-go/common/messenger/nats"
	"github.com/tombenke/axon-go/common/msgs"
	"github.com/tombenke/axon-go/common/msgs/orchestra"
	at "github.com/tombenke/axon-go/common/testing"
	"sync"
	"testing"
)

const (
	checkSendStatusRequest    = "orchestrator sent status request"
	checkStatusReportReceived = "orchestrator received status report"
)

var checklist = []string{
	checkSendStatusRequest,
	checkStatusReportReceived,
}

var logger = logrus.New()

var messengerCfg = messenger.Config{
	Urls:       "localhost:4222",
	UserCreds:  "",
	ClientName: "status-test-client",
	ClusterID:  "test-cluster",
	ClientID:   "status-test-client",
	Logger:     logger,
}

func TestStatus(t *testing.T) {
	// Connect to messaging
	m := messengerImpl.NewMessenger(messengerCfg)
	defer m.Close()

	// Use a WaitGroup to wait for the processes of the testbed to complete their mission
	wg := sync.WaitGroup{}

	// Create a trigger channel to start the test
	triggerCh := make(chan bool)

	// Create a channel to shut down the processes if needed
	doneCh := make(chan bool)

	actorName := "test-actor"

	// Start the processes of the test-bed
	reportCh, testCompletedCh := at.ChecklistProcess(checklist, doneCh, &wg, logger)
	startMockOrchestrator(reportCh, triggerCh, doneCh, &wg, logger, m)

	// Start the sender process
	wg.Add(1)
	go Status(actorName, doneCh, &wg, m, logger)

	// Start testing
	triggerCh <- true

	// Wait until test is completed, then stop the processes
	<-testCompletedCh
	close(doneCh)

	// Wait for the message to come in
	wg.Wait()
}

// startMockOrchestrator starts a standalone process that emulates
// the behaviour of an external orchestrator application.
// Orchestrator waits for an incoming trigger to start the test process via sending a status request,
// then waits for receiving the status response.
// The Mock Orchestrator reports every relevant event to the Checklist process.
// Mock Orchestrator will shut down if it receives a message via the `doneCh` channel.
func startMockOrchestrator(reportCh chan string, triggerCh chan bool, doneCh chan bool, wg *sync.WaitGroup, logger *logrus.Logger, m messenger.Messenger) {
	statusReportCh := make(chan []byte)
	statusReportSubs := m.ChanSubscribe("status-report", statusReportCh)

	wg.Add(1)
	go func() {
		defer statusReportSubs.Unsubscribe()
		defer close(statusReportCh)
		defer wg.Done()

		for {
			select {
			case <-doneCh:
				logger.Infof("MockOrchestrator shuts down.")
				return

			case <-triggerCh:
				logger.Infof("MockOrchestrator received 'start-trigger'.")
				logger.Infof("MockOrchestrator sends 'status-request' message.")
				statusRequestMsg := orchestra.NewStatusRequestMessage()
				m.Publish("status-request", statusRequestMsg.Encode(msgs.JSONRepresentation))
				// TODO: Make orchestra message representations and channel names configurable
				reportCh <- checkSendStatusRequest

			case <-statusReportCh:
				logger.Infof("MockOrchestrator received 'status-report' message.")
				reportCh <- checkStatusReportReceived
			}
		}
	}()
	logger.Infof("Mock Orchestrator started.")
}
