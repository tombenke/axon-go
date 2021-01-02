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
func StartProcessor(procFun func(Context) error, outputsCfg config.Outputs, doneCh chan bool, appWg *sync.WaitGroup, inputsCh chan io.Inputs, logger *logrus.Logger) chan io.Outputs {
	outputsCh := make(chan io.Outputs)

	(*appWg).Add(1)
	go Processor(procFun, outputsCfg, doneCh, appWg, inputsCh, outputsCh, logger)

	return outputsCh
}

// Processor is the implementation of the core process that executes the so called `procFun` function with a context.
// The context provides an interface to the `procFun` to access to the messages of the input ports,
// as well as to access to the output ports that will emit the results of the computation.
func Processor(procFun func(Context) error, outputsCfg config.Outputs, doneCh chan bool, appWg *sync.WaitGroup, inputsCh chan io.Inputs, outputsCh chan io.Outputs, logger *logrus.Logger) {
	logger.Infof("Processor started.")
	defer close(outputsCh)
	defer appWg.Done()

	// Setup the output ports
	outputs := io.NewOutputs(outputsCfg)

	for {
		select {
		case <-doneCh:
			logger.Infof("Processor shuts down.")
			return

		case inputs := <-inputsCh:
			processInputs(inputs, outputs, procFun, outputsCh, logger)
		}
	}
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
