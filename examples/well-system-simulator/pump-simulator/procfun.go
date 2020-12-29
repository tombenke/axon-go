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
	   const pump_capacity = 2. // [m3/h]
	   const power_consumption = 800. // [W]
	   const msec_per_hour = 60 * 60 * 1000

	   // Pump needs power to work
	   ctx.SetOutputData("power-need", power_consumption)

	   if (ctx.GetInputData("power-input") >= power_consumption) {
	       // Pump is ON
	       // Requires water from water source max. as much as the pump's capacity
	       const pump_capacity_dt = pump_capacity / msec_per_hour * ctx.GetInputData("dt")
	       ctx.SetOutputData("water-need", ctx.GetInputData("water-need") > pump_capacity_dt ? pump_capacity_dt : ctx.GetInputData("water-need"))

	       // Forwards the incoming water
	       ctx.SetOutputData("water-output", ctx.GetInputData("water-input"))
	   } else {
	       // Pump is OFF
	       // Requires no water
	       ctx.SetOutputData("water-need", 0.)

	       // Forwards no water
	       ctx.SetOutputData("water-output", 0.)
	   }

	*/
	return nil
}
