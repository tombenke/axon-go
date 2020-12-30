package main

import (
	"github.com/tombenke/axon-go/common/actor/processor"
	"github.com/tombenke/axon-go/common/msgs/base"
)

// ProcessorFun is the message processor function of the actor node
func ProcessorFun(ctx processor.Context) error {

	msecPerHour := float64(60 * 60 * 1000)
	// Inputs
	dt := ctx.GetInputMessage("dt").(*base.Float64).Body.Data / msecPerHour
	ctx.Logger.Infof("dt: %f", dt)

	/* TODO
	   const tank_cross_section = 1. // [m2]
	   const maxWaterLevel = 0.95 // [m]

	   // determine the refill amount
	   const refill_height = ctx.GetInputData("water-input") / tank_cross_section

	   // determine the output amount
	   const output_need_height = ctx.SetOutputData("water-output-need") / tank_cross_section // [m]
	   let waterOutput = ctx.GetInputData("water-output-need")

	   // Determine the new level using the balance
	   const balance = refill_height - output_need_height
	   let waterLevel = ctx.GetInputData("water-buffer-tank-level") + balance

	   // Compensate the output and level according to the tank limits
	   if waterLevel > maxWaterLevel {
	       waterLevel = maxWaterLevel
	   } else if (waterLevel < 0.) {
	       waterOutput = (output_need_height + waterLevel) * tank_cross_section
	       waterLevel = 0.
	   }

	   ctx.SetOutputData("water-output") = waterOutput
	   ctx.SetOutputData("water-buffer-tank-level") = waterLevel

	*/
	return nil
}
