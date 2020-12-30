package main

import (
	"github.com/tombenke/axon-go/common/actor/processor"
	"github.com/tombenke/axon-go/common/msgs/base"
)

// ProcessorFun is the message processor function of the actor node
func ProcessorFun(ctx processor.Context) error {

	// Config parameters
	wellBackfillCapacity := ctx.GetInputMessage("well-backfill-capacity").(*base.Float64).Body.Data // 0.6 [m3/h]
	wellCrossSection := ctx.GetInputMessage("well-cross-section").(*base.Float64).Body.Data         // Math.pow(0.16 / 2., 2) * Math.PI [m2]
	maxWaterLevel := ctx.GetInputMessage("max-water-level").(*base.Float64).Body.Data               // -30 [m]
	minWaterLevel := ctx.GetInputMessage("min-water-level").(*base.Float64).Body.Data               // -36 [m]
	msecPerHour := float64(60 * 60 * 1000)

	// Inputs
	waterNeed := ctx.GetInputMessage("water-need").(*base.Float64).Body.Data
	dt := ctx.GetInputMessage("dt").(*base.Float64).Body.Data / msecPerHour
	waterLevel := ctx.GetInputMessage("water-level").(*base.Float64).Body.Data // maxWaterLevel (the well is full by default)

	// Do backfill up to the max level
	backfillHeight := (wellBackfillCapacity / wellCrossSection) * dt // [m]
	waterLevel = waterLevel + backfillHeight
	if waterLevel > maxWaterLevel {
		waterLevel = maxWaterLevel
	}

	// Determine the real consumption according to the needed one and to the actual level
	consumptionHeight := waterNeed / wellCrossSection // [m]
	if waterLevel-consumptionHeight < minWaterLevel {
		consumptionHeight = waterLevel - minWaterLevel
	}
	waterOutput := consumptionHeight * wellCrossSection // [m3]

	// Determine the new water level removing the water consumption from it
	waterLevel = waterLevel - consumptionHeight

	// Outputs
	ctx.SetOutputMessage("water-output", base.NewFloat64Message(waterOutput))
	ctx.SetOutputMessage("water-level", base.NewFloat64Message(waterLevel))
	return nil
}
