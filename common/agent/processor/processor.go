// Package processor provides the implementation of the `Processor` process, and its helper functions.
package processor

import (
	"github.com/tombenke/axon-go/common/io"
	"github.com/tombenke/axon-go/common/messenger"
	"sync"
)

// StartProcessor starts the `Processor` core process.
func StartProcessor(m messenger.Messenger, done chan bool, appWg *sync.WaitGroup, inputsCh chan io.Inputs, procFun func(Context) error) (outputsCh chan io.Outputs, err error) {
	outputsCh = make(chan io.Outputs)
	(*appWg).Add(1)
	go Processor(procFun, done, appWg, inputsCh, outputsCh)
	return outputsCh, nil
}

// Processor is the implementation of the core process that executes the so called `procFun` function with a context.
// The context provides an interface to the `procFun` to access to the messages of the input ports,
// as well as to access to the output ports that will emit the results of the computation.
func Processor(procFun func(Context) error, done chan bool, appWg *sync.WaitGroup, inputsCh chan io.Inputs, outputsCh chan io.Outputs) {
	defer appWg.Done()
	for {
		select {
		case inputs := <-inputsCh:
			processInputs(inputs, procFun, outputsCh)

		case <-done:
			close(outputsCh)
			return
		}
	}
}

func processInputs(inputs io.Inputs, procFun func(Context) error, outputsCh chan io.Outputs) {
	// TODO: Build Context
	//outputs := make(io.Outputs, 1)
}
