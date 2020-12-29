package main

import (
	"fmt"
	"github.com/tombenke/axon-go/common/actor/processor"
	"github.com/tombenke/axon-go/common/msgs/base"
)

func ProcessorFun(ctx processor.Context) error {

	fmt.Printf("ctx: %v\n", ctx)
	pumpStates := map[string]bool{
		"REFILL-THE-WELL":   false,
		"STANDBY":           false,
		"FILL-THE-TANK":     true,
		"WELL-SENSOR-ERROR": false,
	}

	// Get the inputs
	currentControllerState := ctx.GetInputMessage("well-pump-controller-state").(*base.String).Body.Data
	newControllerState := currentControllerState // Keep the current state in case there is no state change

	tankMax := ctx.GetInputMessage("buffer-tank-upper-level-state").(*base.Bool).Body.Data
	var tankState string
	if tankMax {
		tankState = "FULL"
	} else {
		tankState = "REFILL"
	}
	wellMax := ctx.GetInputMessage("well-water-upper-level-state").(*base.Bool).Body.Data
	wellMin := ctx.GetInputMessage("well-water-lower-level-state").(*base.Bool).Body.Data

	var wellState string
	if wellMax {
		if wellMin {
			wellState = "FULL"
		} else {
			wellState = "ERROR"
		}
	} else {
		if wellMin {
			wellState = "REFILL"
		} else {
			wellState = "EMPTY"
		}
	}

	// Determine the new state
	switch currentControllerState {
	case "REFILL-THE-WELL":
		if tankState == "FULL" && wellState == "FULL" {
			// Both tank and well got full
			newControllerState = "STANDBY"
		} else if tankState == "REFILL" && wellState == "FULL" {
			// Well got full, tank needs refill
			newControllerState = "FILL-THE-TANK"
		} else if wellState == "ERROR" {
			// Well sensor error
			newControllerState = "WELL-SENSOR-ERROR"
		}
		break

	case "STANDBY":
		if tankState == "REFILL" && wellState == "FULL" {
			// Start to fill the tank
			newControllerState = "FILL-THE-TANK"
		} else if wellState == "ERROR" {
			// Well sensor error
			newControllerState = "WELL-SENSOR-ERROR"
		}
		break

	case "FILL-THE-TANK":
		if tankState == "FULL" && wellState == "FULL" {
			// Both tank and well got full
			newControllerState = "STANDBY"
		} else if tankState == "FULL" && (wellState == "EMPTY" || wellState == "REFILL") {
			// Tank got full, well needs refill
			newControllerState = "REFILL-THE-WELL"
		} else if wellState == "ERROR" {
			// Well sensor error
			newControllerState = "WELL-SENSOR-ERROR"
		}
		break

	case "WELL-SENSOR-ERROR":
		if (tankState == "FULL" || tankState == "REFILL") && wellState == "FULL" {
			// go standby
			newControllerState = "STANDBY"
		} else if (tankState == "FULL" || tankState == "REFILL") && (wellState == "EMPTY" || wellState == "REFILL") {
			// Fill the well
			newControllerState = "REFILL-THE-WELL"
		}
		break
	}

	// Set the outputs
	ctx.SetOutputMessage("well-pump-controller-state", base.NewStringMessage(newControllerState))
	ctx.SetOutputMessage("well-pump-relay-state", base.NewBoolMessage(pumpStates[newControllerState]))
	return nil
}
