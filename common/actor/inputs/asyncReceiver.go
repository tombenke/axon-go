// Package inputs provides the functions to receive input messages, collect them and forward to the processor
package inputs

import (
	"github.com/sirupsen/logrus"
	"github.com/tombenke/axon-go/common/config"
	"github.com/tombenke/axon-go/common/io"
	"github.com/tombenke/axon-go/common/messenger"
	"sync"
)

// AsyncReceiver receives inputs from the connecting actors processor function via the `outputsCh`
// that it sends to the processor for further processing.
// The inputs structures hold every details about the ports, the message itself,
// and the subject to receive from.
// This function starts the receiver routine as a standalone process,
// and returns a channel that the process uses to forward the incoming inputs.
func AsyncReceiver(inputsCfg config.Inputs, doneCh chan bool, appWg *sync.WaitGroup, m messenger.Messenger, logger *logrus.Logger) (chan io.Inputs, chan bool) {
	receiverStoppedCh := make(chan bool)

	// Setup communication channel with the processor
	inputsCh := make(chan io.Inputs)

	appWg.Add(1)
	go func() {
		logger.Infof("Receiver started in async mode.")
		defer logger.Infof("Receiver stopped.")
		defer close(inputsCh)
		defer close(receiverStoppedCh)
		defer appWg.Done()

		// Create wait-group for the channel observer sub-processes
		obsWg := sync.WaitGroup{}
		obsDoneCh := make(chan bool)

		// Create Input ports, and initialize with default messages
		inputs := asyncSetupInputPorts(inputsCfg, logger)

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
				logger.Infof("Receiver closed the 'obsDoneCh'.")
				logger.Infof("Receiver starts waiting for observers to stop")
				obsWg.Wait()
				logger.Infof("Receiver's observers stopped")
				return

			case input := <-inputsMuxCh:
				logger.Infof("Receiver got message to '%s' port", input.Name)
				inputs.SetMessage(input.Name, input.Message)
				// Immediately forward to the processor if not in synchronized mode
				inputsCh <- inputs
				logger.Infof("Receiver sent 'inputs' to 'inputsCh'")
			}
		}
	}()

	return inputsCh, receiverStoppedCh
}

// asyncSetupInputPorts creates inputs ports, and initilizes them with their default messages
func asyncSetupInputPorts(inputsCfg config.Inputs, logger *logrus.Logger) io.Inputs {

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
