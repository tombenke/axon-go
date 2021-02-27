package main

import (
	"fmt"
	"github.com/tombenke/axon-go-common/actor/processor"
	"github.com/tombenke/axon-go-common/msgs/base"
)

// getProcessorFun is the message processor function of the actor node
func getProcessorFun(config Config, influxDb InfluxDb) func(ctx processor.Context) error {

	return func(ctx processor.Context) error {
		for inputName := range ctx.Inputs {
			fmt.Printf("INPUT: %v\n", inputName)
			data := ctx.GetInputMessage(inputName).(*base.Float64).Body.Data
			influxDb.WritePoint(inputName, "data", data)
		}
		return nil
	}
}
