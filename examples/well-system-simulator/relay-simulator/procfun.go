package main

import (
	"github.com/tombenke/axon-go/common/actor/processor"
	"github.com/tombenke/axon-go/common/msgs/base"
)

func ProcessorFun(ctx processor.Context) error {
	relayState := ctx.GetInputMessage("relay-state").(*base.Bool).Body.Data
	powerNeedIn := ctx.GetInputMessage("power-need").(*base.Float64).Body.Data

	powerNeedOut := float64(0)
	powerOutput := float64(0)

	// Provide power it got if relay is ON, otherwise provide no power
	if relayState {
		powerOutput = powerNeedIn
		// Forward power need to the power source
		powerNeedOut = powerNeedIn
	}

	ctx.SetOutputMessage("power-need", base.NewFloat64Message(powerNeedOut))
	ctx.SetOutputMessage("power-output", base.NewFloat64Message(powerOutput))
	return nil
}
