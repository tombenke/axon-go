package main

import (
	"github.com/tombenke/axon-go/common/actor/processor"
	"github.com/tombenke/axon-go/common/msgs/base"
)

func ProcessorFun(ctx processor.Context) error {

	// Config parameters
	well_backfill_capacity := ctx.GetInputMessage("well-backfill-capacity").(*base.Float64).Body.Data // 0.6 [m3/h]
	well_cross_section := ctx.GetInputMessage("well-cross-section").(*base.Float64).Body.Data         // Math.pow(0.16 / 2., 2) * Math.PI [m2]
	max_water_level := ctx.GetInputMessage("max-water-level").(*base.Float64).Body.Data               // -30 [m]
	min_water_level := ctx.GetInputMessage("min-water-level").(*base.Float64).Body.Data               // -36 [m]
	msec_per_hour := float64(60 * 60 * 1000)

	// Inputs
	water_need := ctx.GetInputMessage("water-need").(*base.Float64).Body.Data
	dt := ctx.GetInputMessage("dt").(*base.Float64).Body.Data / msec_per_hour
	water_level := ctx.GetInputMessage("water-level").(*base.Float64).Body.Data // max_water_level (the well is full by default)

	// Do backfill up to the max level
	backfill_height := (well_backfill_capacity / well_cross_section) * dt // [m]
	water_level = water_level + backfill_height
	if water_level > max_water_level {
		water_level = max_water_level
	}

	// Determine the real consumption according to the needed one and to the actual level
	consumption_height := water_need / well_cross_section // [m]
	if water_level-consumption_height < min_water_level {
		consumption_height = water_level - min_water_level
	}
	water_output := consumption_height * well_cross_section // [m3]

	// Determine the new water level removing the water consumption from it
	water_level = water_level - consumption_height

	// Outputs
	ctx.SetOutputMessage("water-output", base.NewFloat64Message(water_output))
	ctx.SetOutputMessage("water-level", base.NewFloat64Message(water_level))
	return nil
}
