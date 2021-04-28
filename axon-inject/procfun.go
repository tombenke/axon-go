package main

import (
	"github.com/tombenke/axon-go-common/actor/processor"
)

// getProcessorFun is the message processor function of the actor node
func getProcessorFun(config Config) func(ctx processor.Context) error {
	return func(ctx processor.Context) error {
		injectedMessage := ctx.GetInputMessage("inject")

		// Set the output port value
		ctx.SetOutputMessage("output", injectedMessage)
		return nil
	}
}
