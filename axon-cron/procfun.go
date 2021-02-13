package main

import (
	"fmt"
	"github.com/tombenke/axon-go/common/actor/processor"
	"github.com/tombenke/axon-go/common/msgs/base"
)

// getProcessorFun is the message processor function of the actor node
func getProcessorFun(config Config) func(ctx processor.Context) error {
	return func(ctx processor.Context) error {
		fmt.Printf("CronDef: %s\n", config.CronDef)
		// TODO: Set the `axon.cron` output value
		cronMsg := base.NewAnyMessage(map[string]interface{}{
			"time": int64(1612693423000521770),
			"meta": map[string]interface{}{
				"timePrecision": "ns",
			}},
		)
		ctx.SetOutputMessage("cron", cronMsg)
		return nil
	}
}
