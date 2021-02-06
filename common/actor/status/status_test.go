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

	actorName := "test-actor"

	// Start the processes of the test-bed
	doneChkCh := make(chan bool)
	reportCh, testCompletedCh, checklistStoppedCh := at.ChecklistProcess(checklist, doneChkCh, &wg, logger)

	doneOrcCh := make(chan bool)
	orcStoppedCh := startMockOrchestrator(reportCh, triggerCh, doneOrcCh, &wg, logger, m)

	// Start the status process
	doneStatusCh := make(chan bool)
	statusStoppedCh := Status(actorName, doneStatusCh, &wg, m, logger)

	// Start testing
	triggerCh <- true

	// Wait until test is completed, then stop the processes
	logger.Infof("Wait until test is completed")
	<-testCompletedCh

	logger.Infof("Stops Orchestrator")
	close(doneOrcCh)
	logger.Infof("Wait Orchestrator to stop")
	<-orcStoppedCh
	logger.Infof("Orchestrator stopped")

	logger.Infof("Stops Status")
	close(doneStatusCh)
	logger.Infof("Wait Status to stop")
	<-statusStoppedCh
	logger.Infof("Status stopped")

	logger.Infof("Stops Checklist")
	close(doneChkCh)
	logger.Infof("Wait Checklist to stop")
	<-checklistStoppedCh
	logger.Infof("Checklist stopped")

	// Wait for the message to come in
	wg.Wait()
}

// startMockOrchestrator starts a standalone process that emulates
// the behaviour of an external orchestrator application.
// Orchestrator waits for an incoming trigger to start the test process via sending a status request,
// then waits for receiving the status response.
// The Mock Orchestrator reports every relevant event to the Checklist process.
// Mock Orchestrator will shut down if it receives a message via the `doneCh` channel.
func startMockOrchestrator(reportCh chan string, triggerCh chan bool, doneCh chan bool, wg *sync.WaitGroup, logger *logrus.Logger, m messenger.Messenger) chan bool {
	statusReportCh := make(chan []byte)
	statusReportSubs := m.ChanSubscribe("status-report", statusReportCh)
	orcStoppedCh := make(chan bool)

	wg.Add(1)
	go func() {
		defer func() {
			logger.Infof("MockOrchestrator stopped.")
			if err := statusReportSubs.Unsubscribe(); err != nil {
				panic(err)
			}
			close(statusReportCh)
			close(orcStoppedCh)
			wg.Done()
		}()

		for {
			select {
			case <-doneCh:
				logger.Infof("MockOrchestrator shuts down.")
				return

			case <-triggerCh:
				logger.Infof("MockOrchestrator received 'start-trigger'.")
				logger.Infof("MockOrchestrator sends 'status-request' message.")
				statusRequestMsg := orchestra.NewStatusRequestMessage()
				if err := m.Publish("status-request", statusRequestMsg.Encode(msgs.JSONRepresentation)); err != nil {
					panic(err)
				}
				// TODO: Make orchestra message representations and channel names configurable
				reportCh <- checkSendStatusRequest

			case <-statusReportCh:
				logger.Infof("MockOrchestrator received 'status-report' message.")
				reportCh <- checkStatusReportReceived
			}
		}
	}()
	logger.Infof("Mock Orchestrator started.")

	return orcStoppedCh
}
