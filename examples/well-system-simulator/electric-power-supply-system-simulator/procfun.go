package main

import (
	"github.com/tombenke/axon-go/common/actor/processor"
	"github.com/tombenke/axon-go/common/msgs/base"
)

// ProcessorFun is the message processor function of the actor node
func ProcessorFun(ctx processor.Context) error {
	maxPower := ctx.GetInputMessage("max-power").(*base.Float64).Body.Data
	powerNeed := ctx.GetInputMessage("power-need").(*base.Float64).Body.Data

	var powerOutput float64
	if powerNeed > maxPower {
		powerOutput = maxPower
	} else {
		powerOutput = powerNeed
	}

	ctx.SetOutputMessage("power-output", base.NewFloat64Message(powerOutput))
	return nil
}
