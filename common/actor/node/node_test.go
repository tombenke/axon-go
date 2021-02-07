package node_test

import (
	//"fmt"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/tombenke/axon-go/common/actor/node"
	"github.com/tombenke/axon-go/common/actor/processor"
	"github.com/tombenke/axon-go/common/config"
	"github.com/tombenke/axon-go/common/log"
	"github.com/tombenke/axon-go/common/messenger"
	messengerImpl "github.com/tombenke/axon-go/common/messenger/nats"
	"github.com/tombenke/axon-go/common/msgs"
	"github.com/tombenke/axon-go/common/msgs/base"
	"github.com/tombenke/axon-go/common/msgs/orchestra"
	at "github.com/tombenke/axon-go/common/testing"
	"sync"
	"testing"
)

const (
	actorName = "node-test-actor"

	checkSendStatusRequest    = "orchestrator sent status request"
	checkStatusReportReceived = "orchestrator received status report"
)

var checklist = []string{
	checkSendStatusRequest,
	checkStatusReportReceived,
}

var messengerCfg = messenger.Config{
	Urls:       "localhost:4222",
	UserCreds:  "",
	ClientName: "status-test-client",
	ClusterID:  "test-cluster",
	ClientID:   "status-test-client",
	Logger:     log.Logger,
}

func init() {
	log.SetLevelStr("debug")
}

func TestNodeStartStop(t *testing.T) {

	log.Logger.Infof("\nTestNodeStartStop ==========")
	n := node.NewNode(makeNodeTestConfig(), ProcessorFun)
	assert.NotNil(t, n)
	nwg := sync.WaitGroup{}
	n.Start(&nwg)
	n.Shutdown()
	nwg.Wait()
}

func TestNodeStatus(t *testing.T) {
	log.Logger.Infof("\nTestNodeStatus ==========")
	// Create and start Node to test
	n := node.NewNode(makeNodeTestConfig(), ProcessorFun)
	assert.NotNil(t, n)
	wg := sync.WaitGroup{}
	n.Start(&wg)

	runStatusTest(t, n)
	wg.Wait()
}

func runStatusTest(t *testing.T, n node.Node) {
	// Connect to messaging
	m := messengerImpl.NewMessenger(messengerCfg)
	defer m.Close()
	wg := sync.WaitGroup{}

	// Create a trigger channel to start the test
	triggerCh := make(chan bool)

	// Start the processes of the test-bed
	doneChkCh := make(chan bool)
	reportCh, testCompletedCh, checklistStoppedCh := at.ChecklistProcess(checklist, doneChkCh, &wg, log.Logger)

	doneOrcCh := make(chan bool)
	orcStoppedCh := startMockOrchestrator(reportCh, triggerCh, doneOrcCh, &wg, log.Logger, m)

	// Start testing
	triggerCh <- true

	// Wait until test is completed, then stop the processes
	log.Logger.Infof("Wait until test is completed")
	<-testCompletedCh

	log.Logger.Infof("Stops Orchestrator")
	close(doneOrcCh)
	log.Logger.Infof("Wait Orchestrator to stop")
	<-orcStoppedCh
	log.Logger.Infof("Orchestrator stopped")

	log.Logger.Infof("Shut down the Node (incl. Status)")
	n.Shutdown()

	log.Logger.Infof("Stops Checklist")
	close(doneChkCh)
	log.Logger.Infof("Wait Checklist to stop")
	<-checklistStoppedCh
	log.Logger.Infof("Checklist stopped")
}

// makeNodeTestConfig returns with a built-in configuration for the node test
func makeNodeTestConfig() config.Node {
	// Create the new, empty node with its name and configurability parameters
	node := config.NewNode(actorName, actorName, false, true)

	// Add I/O ports
	node.AddInputPort("reference-water-level", "base/Float64", "application/json", "", `{ "Body": { "Data": 0.75 } }`)
	node.AddInputPort("water-level", "base/Float64", "application/json", "well-water-level", `{ "Body": { "Data": 0.0 } }`)
	node.AddOutputPort("water-level-state", "base/Bool", "application/json", "well-water-upper-level-state")

	return node
}

// ProcessorFun is the message processor function of the actor node
func ProcessorFun(ctx processor.Context) error {
	waterLevel := ctx.GetInputMessage("water-level").(*base.Float64).Body.Data
	referenceWaterLevel := ctx.GetInputMessage("reference-water-level").(*base.Float64).Body.Data

	waterLevelState := waterLevel >= referenceWaterLevel

	ctx.SetOutputMessage("water-level-state", base.NewBoolMessage(waterLevelState))
	return nil
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
		defer logger.Infof("MockOrchestrator stopped.")
		defer func() {
			if err := statusReportSubs.Unsubscribe(); err != nil {
				panic(err)
			}
		}()
		defer close(statusReportCh)
		defer close(orcStoppedCh)
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
