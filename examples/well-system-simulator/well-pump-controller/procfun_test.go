package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/tombenke/axon-go/common/actor/processor"
	"github.com/tombenke/axon-go/common/msgs/base"
	at "github.com/tombenke/axon-go/common/testing"
	"testing"
)

var testCases at.TestCases = at.TestCases{
	at.TestCase{
		Inputs: at.TestCaseMsgs{
			"dt":                            base.NewFloat64Message(1000),
			"well-water-upper-level-state":  base.NewBoolMessage(false),
			"well-water-lower-level-state":  base.NewBoolMessage(false),
			"buffer-tank-upper-level-state": base.NewBoolMessage(false),
			"well-pump-controller-state":    base.NewStringMessage("REFILL-THE-WELL"),
		},
		Outputs: at.TestCaseMsgs{
			"well-pump-relay-state":      base.NewBoolMessage(false),
			"well-pump-controller-state": base.NewStringMessage("REFILL-THE-WELL"),
		},
	},
	// TODO: Add test cases
}

func TestProcessorFun(t *testing.T) {
	for _, tc := range testCases {
		context := processor.SetupContext(tc, inputsCfg, outputsCfg)
		err := ProcessorFun(context)
		assert.Nil(t, err)
		processor.CompareOutputsData(t, context.Outputs, tc)
	}
}
