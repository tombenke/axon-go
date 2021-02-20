package main

import (
	"github.com/tombenke/axon-go/common/actor/processor"
	"github.com/tombenke/axon-go/common/msgs/base"
	"time"
)

func nowAsUnixWithPrecision(precision string) int64 {
	nowNs := time.Now().UnixNano()
	switch precision {
	case "ns":
		return nowNs
	case "u", "us":
		return nowNs / 1e3
	case "ms":
		return nowNs / 1e6
	case "s":
		return nowNs / 1e9
	}
	return nowNs
}

// getProcessorFun is the message processor function of the actor node
func getProcessorFun(config Config) func(ctx processor.Context) error {
	return func(ctx processor.Context) error {
		// Determine the actual time with the expected precision
		nowNs := nowAsUnixWithPrecision(config.Precision)

		// Create the output message
		cronMsg := base.NewAnyMessage(map[string]interface{}{
			"time": nowNs,
			"meta": map[string]interface{}{
				"timePrecision": "ns",
			}},
		)

		// Set the output port value
		ctx.SetOutputMessage("output", cronMsg)
		return nil
	}
}
