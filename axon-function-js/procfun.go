package main

import (
	"encoding/json"
	"fmt"
	"github.com/tombenke/axon-go/common/actor/processor"
	"github.com/tombenke/axon-go/common/msgs/base"
)

// getProcessorFun is the message processor function of the actor node
func getProcessorFun(config Config) func(ctx processor.Context) error {
	return func(ctx processor.Context) error {
		msg := ctx.GetInputMessage("input").(*base.Any)
		msgStr, err := json.MarshalIndent(msg, "", "  ")
		if err != nil {
			panic(err)
		}

		fmt.Printf("\n%s\n", msgStr)
		ctx.SetOutputMessage("output", msg)
		return nil
	}
}
