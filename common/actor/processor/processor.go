// Package processor provides the implementation of the `Processor` process, and its helper functions.
package processor

import (
	"github.com/sirupsen/logrus"
	"github.com/tombenke/axon-go/common/config"
	"github.com/tombenke/axon-go/common/io"
	"sync"
)

// StartProcessor starts the `Processor` core process, then returns an `io.Outputs` channel that forwards the
// results of processing to the outputs.
// Processor is the implementation of the core process that executes the so called `procFun` function with a context.
// The context provides an interface to the `procFun` to access to the messages of the input ports,
// as well as to access to the output ports that will emit the results of the computation.
func StartProcessor(procFun func(Context) error, outputsCfg config.Outputs, doneCh chan bool, appWg *sync.WaitGroup, inputsCh chan io.Inputs, logger *logrus.Logger) (chan io.Outputs, chan bool) {
	outputsCh := make(chan io.Outputs)
	procStoppedCh := make(chan bool)

	(*appWg).Add(1)
	go func() {
		logger.Infof("Processor started.")
		defer logger.Infof("Processor stopped.")
		defer close(outputsCh)
		defer close(procStoppedCh)
		defer appWg.Done()

		// Setup the output ports
		outputs := io.NewOutputs(outputsCfg)

		for {
			select {
			case <-doneCh:
				logger.Infof("Processor shuts down.")
				return

			case inputs := <-inputsCh:
				logger.Infof("Processor got inputs")
				processInputs(inputs, outputs, procFun, outputsCh, logger)
			}
		}
	}()

	return outputsCh, procStoppedCh
}

func processInputs(inputs io.Inputs, outputs io.Outputs, procFun func(Context) error, outputsCh chan io.Outputs, logger *logrus.Logger) {
	context := NewContext(logger, inputs, outputs)

	logger.Infof("Processor calls processor-function")
	err := procFun(context)
	if err != nil {
		panic(err)
	}

	logger.Infof("Processor sends the results")
	outputsCh <- context.Outputs
}
