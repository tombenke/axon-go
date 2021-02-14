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
	doneCh    chan bool
	resetCh   chan bool

	doneStatusCh    chan bool
	doneInputsRcvCh chan bool
	doneProcessorCh chan bool
	doneOutputsCh   chan bool

	// Declare the channels for communication among the componens
	inputsCh  chan io.Inputs
	outputsCh chan io.Outputs

	// Declare the channels through which the components notify that they have stopped
	inputsRcvStoppedCh chan bool
	processorStoppedCh chan bool
	outputsStoppedCh   chan bool
	statusStoppedCh    chan bool
	wg                 *sync.WaitGroup
}

// NewNode creates and returns with a new `Node` object
// which represents the common core component of an actor-node application
func NewNode(config config.Node, procFun func(processor.Context) error) Node {
	node := Node{
		config:  config,
		name:    config.Name,
		procFun: procFun,
		doneCh:  make(chan bool),
		resetCh: make(chan bool),

		// Create channels to control the shut down of the components
		doneStatusCh:    make(chan bool),
		doneInputsRcvCh: make(chan bool),
		doneProcessorCh: make(chan bool),
		doneOutputsCh:   make(chan bool),
		wg:              &sync.WaitGroup{},
	}

	// Configure the global logger of the application according to the configuration
	log.SetLevelStr(config.LogLevel)
	log.SetFormatterStr(config.LogFormat)

	// Connect to messaging
	node.config.Messenger.Logger = log.Logger
	node.config.Messenger.ClientID = node.name
	node.config.Messenger.ClientName = node.name
	node.config.Messenger.ClusterID = "test-cluster"
	node.messenger = messengerImpl.NewMessenger(node.config.Messenger)

	log.Logger.Debugf("Start '%s' actor node's internal components", node.config.Name)
	// Start the status component to communicate with the orchestrator
	node.statusStoppedCh = status.Status(node.config.Name, node.doneStatusCh, node.wg, node.messenger, log.Logger)

	// Start the core components of the Node
	if node.config.Orchestration.Synchronization {
		// Start the core components in synchronous mode
		node.inputsCh, node.inputsRcvStoppedCh = inputs.SyncReceiver(node.config.Ports.Inputs, node.resetCh, node.doneInputsRcvCh, node.wg, node.messenger, log.Logger)
		node.outputsCh, node.processorStoppedCh = processor.StartProcessor(node.procFun, node.config.Ports.Outputs, node.doneProcessorCh, node.wg, node.inputsCh, log.Logger)
		node.outputsStoppedCh = outputs.SyncSender(node.name, node.outputsCh, node.doneOutputsCh, node.wg, node.messenger, log.Logger)
	} else {
		// Start the core components in asynchronous mode
		node.inputsCh, node.inputsRcvStoppedCh = inputs.AsyncReceiver(node.config.Ports.Inputs, node.resetCh, node.doneInputsRcvCh, node.wg, node.messenger, log.Logger)
		node.outputsCh, node.processorStoppedCh = processor.StartProcessor(node.procFun, node.config.Ports.Outputs, node.doneProcessorCh, node.wg, node.inputsCh, log.Logger)
		node.outputsStoppedCh = outputs.AsyncSender(node.name, node.outputsCh, node.doneOutputsCh, node.wg, node.messenger, log.Logger)
	}
	return node
}

// Start starts the core engine of an actor-node application
func (n Node) Start() {

	log.Logger.Infof("Start '%s' actor node", n.config.Name)

	// Start waiting for the shutdown signal
	n.wg.Add(1)
	go func() {
		log.Logger.Debugf("Node started.")
		defer log.Logger.Debugf("Node stopped.")
		defer n.wg.Done()

		<-n.doneCh
		log.Logger.Debugf("Node is shutting down")

		// Stop status
		close(n.doneStatusCh)
		<-n.statusStoppedCh

		// The components of the processing pipeline must be shut down in reverse order
		// otherwise the channel close might cause problems

		// Stop outputs
		close(n.doneOutputsCh)
		<-n.outputsStoppedCh

		// Stop processor
		close(n.doneProcessorCh)
		<-n.processorStoppedCh

		// Stop inputs receiver
		close(n.doneInputsRcvCh)
		<-n.inputsRcvStoppedCh

		// Close the RESET mechanism
		close(n.resetCh)

		n.messenger.Close()
	}()

	// RESET the Node
	n.Reset()
}

// Wait waits until the internal components of the Node terminates
func (n Node) Wait() {
	n.wg.Wait()
}

// Reset triggers the RESET process in the components of the Node
func (n Node) Reset() {
	//n.resetCh <- true
}

// Shutdown stops the Node process
func (n Node) Shutdown() {
	close(n.doneCh)
}

// Next Injects the `inputs` messages into the inputs channel, like it were received by the input ports.
func (n Node) Next(inputs io.Inputs) {
	log.Logger.Debugf("Node.Next() is called\n")
	n.inputsCh <- inputs
}

func (n Node) NewInputs() io.Inputs {
	return io.NewInputs(n.config.Ports.Inputs)
}
