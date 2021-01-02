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
			"dt":                     base.NewFloat64Message(1000),
			"well-backfill-capacity": base.NewFloat64Message(0.6),
			"well-cross-section":     base.NewFloat64Message(0.0201056),
			"max-water-level":        base.NewFloat64Message(-30),
			"min-water-level":        base.NewFloat64Message(-36),
			"water-need":             base.NewFloat64Message(0),
			"water-level":            base.NewFloat64Message(-30),
		},
		Outputs: at.TestCaseMsgs{
			"water-level":  base.NewFloat64Message(-30),
			"water-output": base.NewFloat64Message(0),
		},
	},
}

func TestProcessorFun(t *testing.T) {
	for _, tc := range testCases {
		context := processor.SetupContext(tc, inputsCfg, outputsCfg)
		err := ProcessorFun(context)
		assert.Nil(t, err)
		processor.CompareOutputsData(t, context, tc)
	}
}
