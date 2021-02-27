package main

import (
	"github.com/tombenke/axon-go-common/actor/processor"
)

// getProcessorFun is the message processor function of the actor node
func getProcessorFun(config Config) func(ctx processor.Context) error {

	jsvm := NewJSVM(config.ScriptFile)

	return func(ctx processor.Context) error {
		return jsvm.Run(ctx)
	}
}
