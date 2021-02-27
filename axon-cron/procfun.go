package main

import (
	"github.com/tombenke/axon-go-common/actor/processor"
	"github.com/tombenke/axon-go-common/msgs/base"
	"github.com/tombenke/axon-go-common/msgs/common"
	"time"
)

// getProcessorFun is the message processor function of the actor node
func getProcessorFun(config Config) func(ctx processor.Context) error {
	return func(ctx processor.Context) error {
		cronMsg := base.NewEmptyMessageAt(time.Now().UnixNano(), common.TimePrecision(config.Precision))

		// Set the output port value
		ctx.SetOutputMessage("output", cronMsg)
		return nil
	}
}
