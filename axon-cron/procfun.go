package main

import (
	"github.com/tombenke/axon-go/common/actor/processor"
	"github.com/tombenke/axon-go/common/msgs/base"
)

// ProcessorFun is the message processor function of the actor node
func ProcessorFun(ctx processor.Context) error {
	/*
		waterLevel := ctx.GetInputMessage("water-level").(*base.Float64).Body.Data
		referenceWaterLevel := ctx.GetInputMessage("reference-water-level").(*base.Float64).Body.Data

		waterLevelState := waterLevel >= referenceWaterLevel

		ctx.SetOutputMessage("water-level-state", base.NewBoolMessage(waterLevelState))
	*/
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
