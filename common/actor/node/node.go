// Package node contains the implementation of the `Node` component
// that is the core element of every actor-node application
package node

import (
	"sync"

	//"github.com/sirupsen/logrus"
	"github.com/tombenke/axon-go/common/config"
	"github.com/tombenke/axon-go/common/log"
	"github.com/tombenke/axon-go/common/messenger"
	messengerImpl "github.com/tombenke/axon-go/common/messenger/nats"
)

// Node represents the common core object of an actor-node application
type Node struct {
	Messenger messenger.Config
	Name      string
	//  Inputs    Inputs
	//	Outputs   Outputs
}

// NewNode creates and returns with a new `Node` object
// which represents the common core component of an actor-node application
func NewNode(config config.Node) Node {
	node := Node{
		Messenger: config.Messenger,
		Name:      config.Name,
	}
	return node
}

// Start starts the core engine of an actor-node application
func (n Node) Start() {

	logger := log.Logger
	logger.Infof("Start '%s' actor node", n.Name)

	// Connect to messaging
	n.Messenger.Logger = log.Logger
	n.Messenger.ClientID = n.Name
	n.Messenger.ClientName = n.Name
	n.Messenger.ClusterID = "test-cluster"
	m := messengerImpl.NewMessenger(n.Messenger)
	defer m.Close()

	// Use a WaitGroup to wait for the processes of the node
	wg := sync.WaitGroup{}

	/* TODO
	// Create a trigger channel to start the test
	triggerCh := make(chan bool)

	nodeName := "test-node"

	// Start the processes of the node
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
	*/

	// Wait for the message to come in
	wg.Wait()
}
