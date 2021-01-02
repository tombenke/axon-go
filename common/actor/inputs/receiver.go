// Package inputs provides the functions to receive input messages, collect them and forward to the processor
package inputs

import (
	"github.com/sirupsen/logrus"
	"github.com/tombenke/axon-go/common/config"
	"github.com/tombenke/axon-go/common/io"
	"github.com/tombenke/axon-go/common/messenger"
	"github.com/tombenke/axon-go/common/msgs"
	"github.com/tombenke/axon-go/common/msgs/orchestra"
	"sync"
)

// Receiver receives inputs from the connecting actors processor function via the `outputsCh`
// that it sends to the processor for further processing.
// The inputs structures hold every details about the ports, the message itself,
// and the subject to receive from.
// This function starts the receiver routine as a standalone process,
// and returns a channel that the process uses to forward the incoming inputs.
func Receiver(syncMode bool, inputsCfg config.Inputs, doneCh chan bool, appWg *sync.WaitGroup, m messenger.Messenger, logger *logrus.Logger) chan io.Inputs {

	// Setup communication channel with the processor
	inputsCh := make(chan io.Inputs)

	appWg.Add(1)
	go func() {
		logger.Infof("Receiver started.")
		defer close(inputsCh)
		defer appWg.Done()

		// Create wait-group for the channel observer sub-processes
		obsWg := sync.WaitGroup{}
		obsDoneCh := make(chan bool)

		// Setup communication channels with the orchestrator
		receiveAndProcessCh := make(chan []byte)
		receiveAndProcessSubs := m.ChanSubscribe("receive-and-process", receiveAndProcessCh)
		defer receiveAndProcessSubs.Unsubscribe()
		defer close(receiveAndProcessCh)

		// Create Input ports, and initialize with default messages
		inputs := setupInputPorts(syncMode, inputsCfg, logger)

		// Creates an inputs multiplexer channel for observers to send their inputs via one channel
		inputsMuxCh := make(chan io.Input)
		defer close(inputsMuxCh)

		// Starts the input port observers
		startInPortsObservers(inputs, inputsMuxCh, obsDoneCh, &obsWg, m, logger)

		for {
			select {
			case <-doneCh:
				logger.Infof("Receiver shuts down.")
				close(obsDoneCh)
				obsWg.Wait()
				return

			case input := <-inputsMuxCh:
				logger.Infof("Receiver got message to '%s' port", input.Name)
				inputs.SetMessage(input.Name, input.Message)
				// TODO: Immediately forward to the processor if not in synchronized mode
				// TODO: In synchronized mode set the message for the _timestamp and _dt virtual ports

			case messageBytes := <-receiveAndProcessCh:
				logger.Infof("Receiver received 'receive-and-process' message from orchestrator")
				receiveAndProcessMsg := orchestra.NewReceiveAndProcessMessage(float64(0))
				receiveAndProcessMsg.Decode(msgs.JSONRepresentation, messageBytes)
				if syncMode {
					inputs.SetMessage("_RAP", receiveAndProcessMsg)
				}
				inputsCh <- inputs
				// TODO: use only in synchronized mode
			}
		}
	}()

	return inputsCh
}

// setupInputPorts creates inputs ports, and initilizes them with their default messages
func setupInputPorts(syncMode bool, inputsCfg config.Inputs, logger *logrus.Logger) io.Inputs {

	logger.Infof("Receiver sets up input ports")

	// Extends the input ports with '_RAP'
	if syncMode {
		if findPortCfgByName(inputsCfg, "_RAP") {
			panic("Can not define an input port with the '_RAP' reserved name.")
		}

		rapInPort := config.In{
			IO: config.IO{
				Name:           "_RAP",
				Type:           "orchestra/ReceiveAndProcess",
				Representation: "application/json",
				Channel:        "_RAP",
			},
			Default: "",
		}

		inputsCfg = append(inputsCfg, rapInPort)
	}

	// Create input ports
	inputs := io.NewInputs(inputsCfg)

	// Set every input ports' message to its default
	for p := range inputs {
		defaultMessage := inputs[p].DefaultMessage
		(&inputs).SetMessage(p, defaultMessage)
	}

	return inputs
}

// findPortCgfByName returns true if it finds an input port named to `portName` in the `inputsCfg` array,
// otherwise it returns false.
func findPortCfgByName(inputsCfg config.Inputs, portName string) bool {
	for p := range inputsCfg {
		if inputsCfg[p].Name == portName {
			return true
		}
	}

	return false
}

// startInPortsObservers starts one message observer for every port,
// and returns with the number of observers started.
func startInPortsObservers(inputs io.Inputs, inputsMuxCh chan io.Input, doneCh chan bool, wg *sync.WaitGroup, m messenger.Messenger, logger *logrus.Logger) {
	for p := range inputs {
		newPortObserver(inputs[p], inputsMuxCh, doneCh, wg, m, logger)
	}
}

// newPortObserver subscribes to an input channel with a go routine that observes the incoming messages.
// When a message arrives through the channel, the go routine forwards that through the `inCh` towards the aggregator.
// The newPortObserver creates and returns with the `inCh` channel that the aggregator can consume.
func newPortObserver(input io.Input, inputsMuxCh chan io.Input, doneCh chan bool, wg *sync.WaitGroup, m messenger.Messenger, logger *logrus.Logger) {
	inMsgCh := make(chan []byte)
	inMsgSubs := m.ChanSubscribe(input.Channel, inMsgCh)

	wg.Add(1)
	go func() {
		logger.Infof("Receiver started new observer on '%s' port", input.Name)
		defer inMsgSubs.Unsubscribe()
		defer close(inMsgCh)
		defer wg.Done()

		for {
			select {
			case <-doneCh:
				logger.Infof("Input message observer on '%s' port shuts down.", input.Name)
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
