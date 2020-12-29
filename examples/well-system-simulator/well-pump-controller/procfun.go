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
	   const pumpStates = {
	       "REFILL-THE-WELL": false,
	       "STANDBY": false,
	       "FILL-THE-TANK": true,
	       "WELL-SENSOR-ERROR": false
	   }

	   const defaultControllerState = "REFILL-THE-WELL" // The initial state

	   // Get the inputs
	   const currentControllerState = ctx.GetInputData("well-pump-controller-state") || defaultControllerState
	   let newControllerState = currentControllerState // Keep the current state in case there is no state change

	   const tankMax = ctx.GetInputData("buffer-tank-upper-level-state")
	   const tankState = tankMax ? "FULL" : "REFILL"

	   const wellMax = ctx.GetInputData("well-water-upper-level-state")
	   const wellMin = ctx.GetInputData("well-water-lower-level-state")
	   const wellState = wellMax
	       ? wellMin ? "FULL" : "ERROR"
	       : wellMin ? "REFILL" : "EMPTY"

	   // Determine the new state
	   switch (currentControllerState) {
	       case "REFILL-THE-WELL":
	           if (tankState === "FULL" && wellState === "FULL") {
	               // Both tank and well got full
	               newControllerState = "STANDBY"
	           } else if (tankState === "REFILL" && wellState === "FULL") {
	               // Well got full, tank needs refill
	               newControllerState = "FILL-THE-TANK"
	           } else if (wellState === "ERROR") {
	               // Well sensor error
	               newControllerState = "WELL-SENSOR-ERROR"
	           }
	           break

	       case "STANDBY":
	           if (tankState === "REFILL" && wellState === "FULL") {
	               // Start to fill the tank
	               newControllerState = "FILL-THE-TANK"
	           } else if (wellState === "ERROR") {
	               // Well sensor error
	               newControllerState = "WELL-SENSOR-ERROR"
	           }
	           break

	       case "FILL-THE-TANK":
	           if (tankState === "FULL" && wellState === "FULL") {
	               // Both tank and well got full
	               newControllerState = "STANDBY"
	           } else if (tankState === "FULL" && (wellState === "EMPTY" || wellState === "REFILL")) {
	               // Tank got full, well needs refill
	               newControllerState = "REFILL-THE-WELL"
	           } else if (wellState === "ERROR") {
	               // Well sensor error
	               newControllerState = "WELL-SENSOR-ERROR"
	           }
	           break

	       case "WELL-SENSOR-ERROR":
	           if ((tankState === "FULL" || tankState === "REFILL") && wellState === "FULL") {
	               // go standby
	               newControllerState = "STANDBY"
	           } else if ((tankState === "FULL" || tankstate === "REFILL") && (wellState === "EMPTY" || wellState === "REFILL")) {
	               // Fill the well
	               newControllerState = "REFILL-THE-WELL"
	           }
	           break
	   }

	   // Set the outputs
	   ctx.SetOutputData("well-pump-controller-state") = newControllerState
	   ctx.SetOutputData("well-pump-relay-state") = pumpStates[newControllerState]
	*/
	return nil
}
