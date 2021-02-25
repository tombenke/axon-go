package main

import (
	"fmt"
	"github.com/tombenke/axon-go/common/actor/processor"
	"github.com/tombenke/axon-go/common/msgs/base"
)

// getProcessorFun is the message processor function of the actor node
func getProcessorFun(config Config, influxDb InfluxDb) func(ctx processor.Context) error {

	return func(ctx processor.Context) error {
		data := ctx.GetInputMessage("input").(*base.Float64).Body.Data
		fmt.Printf("INPUT: %v\n", data)
		influxDb.WritePoint(data)
		return nil
	}
}
