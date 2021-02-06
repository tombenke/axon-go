// Package node contains the implementation of the `Node` component
// that is the core element of every actor-node application
package node

import (
	"sync"

	"github.com/tombenke/axon-go/common/actor/inputs"
	"github.com/tombenke/axon-go/common/actor/outputs"
	"github.com/tombenke/axon-go/common/actor/processor"
	"github.com/tombenke/axon-go/common/actor/status"
	"github.com/tombenke/axon-go/common/config"
	"github.com/tombenke/axon-go/common/io"
	"github.com/tombenke/axon-go/common/log"
	"github.com/tombenke/axon-go/common/messenger"
	messengerImpl "github.com/tombenke/axon-go/common/messenger/nats"
)

// Node represents the common core object of an actor-node application
type Node struct {
	config    config.Node
	messenger messenger.Messenger
	name      string
	procFun   func(processor.Context) error
	done      chan bool
}

// NewNode creates and returns with a new `Node` object
// which represents the common core component of an actor-node application
func NewNode(config config.Node, procFun func(processor.Context) error) Node {
	node := Node{
		config:  config,
		name:    config.Name,
		procFun: procFun,
		done:    make(chan bool),
	}

	// Connect to messaging
	node.config.Messenger.Logger = log.Logger
	node.config.Messenger.ClientID = node.name
	node.config.Messenger.ClientName = node.name
	node.config.Messenger.ClusterID = "test-cluster"
	node.messenger = messengerImpl.NewMessenger(node.config.Messenger)

	return node
}

// Start starts the core engine of an actor-node application
func (n Node) Start(nodeWg *sync.WaitGroup) {

	logger := log.Logger
	logger.Infof("Start '%s' actor node", n.config.Name)

	// Create channels to control the shut down of the components
	doneStatusCh := make(chan bool)
	doneInputsRcvCh := make(chan bool)
	doneProcessorCh := make(chan bool)
	doneOutputsCh := make(chan bool)

	// Declare the channels through which the components notify that they have stopped
	var inputsRcvStoppedCh chan bool
	var processorStoppedCh chan bool
	var outputsStoppedCh chan bool

	// Declare the channels for communication among the componens
	var inputsCh chan io.Inputs
	var outputsCh chan io.Outputs

	// Start the status component to communicate with the orchestrator
	statusStoppedCh := status.Status(n.config.Name, doneStatusCh, nodeWg, n.messenger, log.Logger)

	if n.config.Orchestration.Synchronization {
		// Start the core components in synchronous mode
		inputsCh, inputsRcvStoppedCh = inputs.SyncReceiver(n.config.Ports.Inputs, doneInputsRcvCh, nodeWg, n.messenger, log.Logger)
		outputsCh, processorStoppedCh = processor.StartProcessor(n.procFun, n.config.Ports.Outputs, doneProcessorCh, nodeWg, inputsCh, log.Logger)
		outputsStoppedCh = outputs.SyncSender(n.name, outputsCh, doneOutputsCh, nodeWg, n.messenger, log.Logger)
	} else {
		// Start the core components in asynchronous mode
		inputsCh, inputsRcvStoppedCh = inputs.AsyncReceiver(n.config.Ports.Inputs, doneInputsRcvCh, nodeWg, n.messenger, log.Logger)
		outputsCh, processorStoppedCh = processor.StartProcessor(n.procFun, n.config.Ports.Outputs, doneProcessorCh, nodeWg, inputsCh, log.Logger)
		outputsStoppedCh = outputs.AsyncSender(n.name, outputsCh, doneOutputsCh, nodeWg, n.messenger, log.Logger)
	}

	// Start waiting for the shutdown signal
	nodeWg.Add(1)
	go func() {
		logger.Infof("Node started.")
		defer logger.Infof("Node stopped.")
		defer nodeWg.Done()

		<-n.done
		logger.Infof("Node is shutting down")
		// Stop status
		close(doneStatusCh)
		<-statusStoppedCh

		// The components of the processing pipeline must be shut down in reverse order
		// otherwise the channel close might cause problems

		// Stop outputs
		close(doneOutputsCh)
		<-outputsStoppedCh

		// Stop processor
		close(doneProcessorCh)
		<-processorStoppedCh

		// Stop inputs receiver
		close(doneInputsRcvCh)
		<-inputsRcvStoppedCh

		n.messenger.Close()
	}()
}

// Shutdown stops the Node process
func (n Node) Shutdown() {
	n.done <- true
}
