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
	// TODO: Implement the business logic

	return nil
}
