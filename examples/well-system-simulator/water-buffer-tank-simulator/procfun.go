package main

import (
	"github.com/tombenke/axon-go/common/actor/processor"
	"github.com/tombenke/axon-go/common/msgs/base"
)

func ProcessorFun(ctx processor.Context) error {

	msec_per_hour := float64(60 * 60 * 1000)
	// Inputs
	dt := ctx.GetInputMessage("dt").(*base.Float64).Body.Data / msec_per_hour
	ctx.Logger.Infof("dt: %f", dt)

	/* TODO
	   const tank_cross_section = 1. // [m2]
	   const max_water_level = 0.95 // [m]

	   // determine the refill amount
	   const refill_height = ctx.GetInputData("water-input") / tank_cross_section

	   // determine the output amount
	   const output_need_height = ctx.SetOutputData("water-output-need") / tank_cross_section // [m]
	   let water_output = ctx.GetInputData("water-output-need")

	   // Determine the new level using the balance
	   const balance = refill_height - output_need_height
	   let water_level = ctx.GetInputData("water-buffer-tank-level") + balance

	   // Compensate the output and level according to the tank limits
	   if water_level > max_water_level {
	       water_level = max_water_level
	   } else if (water_level < 0.) {
	       water_output = (output_need_height + water_level) * tank_cross_section
	       water_level = 0.
	   }

	   ctx.SetOutputData("water-output") = water_output
	   ctx.SetOutputData("water-buffer-tank-level") = water_level

	*/
	return nil
}
