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

// SyncReceiver receives inputs from the connecting actors processor function via the `outputsCh`
// that it sends to the processor for further processing.
// The inputs structures hold every details about the ports, the message itself,
// and the subject to receive from.
// This function starts the receiver routine as a standalone process,
// and returns a channel that the process uses to forward the incoming inputs.
func SyncReceiver(inputsCfg config.Inputs, resetCh chan bool, doneCh chan bool, appWg *sync.WaitGroup, m messenger.Messenger, logger *logrus.Logger) (chan io.Inputs, chan bool) {
	receiverStoppedCh := make(chan bool)

	// Setup communication channel with the processor
	inputsCh := make(chan io.Inputs)

	appWg.Add(1)
	go func() {
		logger.Debugf("Receiver started in sync mode.")
		defer logger.Debugf("Receiver stopped.")
		defer close(inputsCh)
		defer close(receiverStoppedCh)
		defer appWg.Done()

		// Create wait-group for the channel observer sub-processes
		obsWg := sync.WaitGroup{}
		obsDoneCh := make(chan bool)

		// Setup communication channels with the orchestrator
		receiveAndProcessCh := make(chan []byte)
		receiveAndProcessSubs := m.ChanSubscribe("receive-and-process", receiveAndProcessCh)
		defer func() {
			if err := receiveAndProcessSubs.Unsubscribe(); err != nil {
				panic(err)
			}
			close(receiveAndProcessCh)
		}()

		// Create Input ports, and initialize with default messages
		inputs := syncSetupInputPorts(inputsCfg, logger)

		// Creates an inputs multiplexer channel for observers to send their inputs via one channel
		inputsMuxCh := make(chan io.Input)
		defer close(inputsMuxCh)

		// Starts the input port observers
		startInPortsObservers(inputs, inputsMuxCh, obsDoneCh, &obsWg, m, logger)

		for {
			select {
			case <-doneCh:
				logger.Debugf("Receiver shuts down.")
				close(obsDoneCh)
				logger.Debugf("Receiver closed the 'obsDoneCh'.")
				logger.Debugf("Receiver starts waiting for observers to stop")
				obsWg.Wait()
				logger.Debugf("Receiver's observers stopped")
				return

			case <-resetCh:
				logger.Debugf("Receiver got RESET signal")
				receiveAndProcessMsg := orchestra.NewReceiveAndProcessMessage(float64(0))
				inputs.SetMessage("_RAP", receiveAndProcessMsg)
				inputsCh <- inputs
				logger.Debugf("Receiver sent 'inputs' to 'inputsCh'")

			case input := <-inputsMuxCh:
				logger.Debugf("Receiver got message to '%s' port", input.Name)
				inputs.SetMessage(input.Name, input.Message)

			case messageBytes := <-receiveAndProcessCh:
				logger.Debugf("Receiver received 'receive-and-process' message from orchestrator")
				receiveAndProcessMsg := orchestra.NewReceiveAndProcessMessage(float64(0))
				if err := receiveAndProcessMsg.Decode(msgs.JSONRepresentation, messageBytes); err != nil {
					panic(err)
				}
				inputs.SetMessage("_RAP", receiveAndProcessMsg)
				inputsCh <- inputs
				logger.Debugf("Receiver sent 'inputs' to 'inputsCh'")
			}
		}
	}()

	return inputsCh, receiverStoppedCh
}

// setupInputPorts creates inputs ports, and initilizes them with their default messages
func syncSetupInputPorts(inputsCfg config.Inputs, logger *logrus.Logger) io.Inputs {

	logger.Debugf("Receiver sets up input ports")

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
