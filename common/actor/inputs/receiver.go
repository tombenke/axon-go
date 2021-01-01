// Package inputs provides the functions to receive input messages, collect them and forward to the processor
package inputs

import (
	"github.com/sirupsen/logrus"
	"github.com/tombenke/axon-go/common/config"
	"github.com/tombenke/axon-go/common/io"
	"github.com/tombenke/axon-go/common/messenger"
	"sync"
)

// Receiver receives inputs from the connecting actors processor function via the `outputsCh`
// that it sends to the processor for further processing.
// The inputs structures hold every details about the ports, the message itself,
// and the subject to receive from.
// This function starts the receiver routine as a standalone process,
// and returns a channel that the process uses to forward the incoming inputs.
func Receiver(inputsCfg config.Inputs, doneCh chan bool, appWg *sync.WaitGroup, m messenger.Messenger, logger *logrus.Logger) chan io.Inputs {

	// Setup communication channel with the processor
	inputsCh := make(chan io.Inputs)

	appWg.Add(1)
	go func() {
		logger.Infof("Receiver started.")
		defer close(inputsCh)
		defer appWg.Done()

		// Setup communication channels with the orchestrator
		receiveAndProcessCh := make(chan []byte)
		receiveAndProcessSubs := m.ChanSubscribe("receive-and-process", receiveAndProcessCh)
		defer receiveAndProcessSubs.Unsubscribe()
		defer close(receiveAndProcessCh)

		// Create Input ports, and initialize with default messages
		inputs := setupInputPorts(inputsCfg, logger)

		// Creates an inputs multiplexer channel for observers to send their inputs via one channel
		inputsMuxCh := make(chan io.Input)
		defer close(inputsMuxCh)

		// Starts the input port observers
		startInPortsObservers(inputs, inputsMuxCh, doneCh, appWg, m, logger)
		//appWg.Add(numObservers)

		for {
			select {
			case <-doneCh:
				logger.Infof("Receiver shuts down.")
				return

			case input := <-inputsMuxCh:
				logger.Infof("Receiver received an input message from '%s' channel on '%s' port", input.Channel, input.Name)
				inputs.SetMessage(input.Name, input.Message)
				// TODO: immediately forward to processor if not in synchronized mode

			case <-receiveAndProcessCh:
				logger.Infof("Receiver received 'receive-and-process' message from orchestrator")
				inputsCh <- inputs
				// TODO: use only in synchronized mode
			}
		}
	}()

	return inputsCh
}

// setupInputPorts creates inputs ports, and initilizes them with their default messages
func setupInputPorts(inputsCfg config.Inputs, logger *logrus.Logger) io.Inputs {

	logger.Infof("Receiver sets up input ports")
	// Create input ports
	inputs := io.NewInputs(inputsCfg)

	// Set every input ports' message to its default
	for p := range inputs {
		defaultMessage := inputs[p].DefaultMessage
		(&inputs).SetMessage(p, defaultMessage)
	}

	return inputs
}

// startInPortsObservers starts one message observer for every port,
// and returns with the number of observers started.
func startInPortsObservers(inputs io.Inputs, inputsMuxCh chan io.Input, doneCh chan bool, wg *sync.WaitGroup, m messenger.Messenger, logger *logrus.Logger) int {
	for p := range inputs {
		wg.Add(1)
		newPortObserver(inputs[p], inputsMuxCh, doneCh, wg, m, logger)
	}
	return len(inputs)
}

// newPortObserver subscribes to an input channel with a go routine that observes the incoming messages.
// When a message arrives through the channel, the go routine forwards that through the `inCh` towards the aggregator.
// The newPortObserver creates and returns with the `inCh` channel that the aggregator can consume.
func newPortObserver(input io.Input, inputsMuxCh chan io.Input, doneCh chan bool, wg *sync.WaitGroup, m messenger.Messenger, logger *logrus.Logger) {
	inMsgCh := make(chan []byte)
	inMsgSubs := m.ChanSubscribe(input.Channel, inMsgCh)

	go func() {
		logger.Infof("Receiver started new observer of '%s' channel for '%s' port", input.Channel, input.Name)
		defer inMsgSubs.Unsubscribe()
		defer close(inMsgCh)
		defer wg.Done()

		for {
			select {
			case <-doneCh:
				logger.Infof("Input message observer for '%s' port shuts down.", input.Name)
				return

			case inputMsg := <-inMsgCh:
				logger.Infof("Input message observer received a message via '%s' channel for '%s' port", input.Channel, input.Name)
				input.Message.Decode(input.Representation, inputMsg)
				inputsMuxCh <- input
			}
		}
	}()

	return
}
