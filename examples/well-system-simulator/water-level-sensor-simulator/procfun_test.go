package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/tombenke/axon-go/common/msgs/base"
	at "github.com/tombenke/axon-go/common/testing"
	"testing"
)

var testCases at.TestCases = at.TestCases{
	at.TestCase{
		Inputs: at.TestCaseMsgs{
			"reference-water-level": base.NewFloat64Message(0.75),
			"water-level":           base.NewFloat64Message(0.0),
		},
		Outputs: at.TestCaseMsgs{
			"water-level-state": base.NewBoolMessage(false),
		},
	},
	at.TestCase{
		Inputs: at.TestCaseMsgs{
			"reference-water-level": base.NewFloat64Message(0.75),
			"water-level":           base.NewFloat64Message(0.75),
		},
		Outputs: at.TestCaseMsgs{
			"water-level-state": base.NewBoolMessage(true),
		},
	},
	at.TestCase{
		Inputs: at.TestCaseMsgs{
			"reference-water-level": base.NewFloat64Message(0.75),
			"water-level":           base.NewFloat64Message(0.8),
		},
		Outputs: at.TestCaseMsgs{
			"water-level-state": base.NewBoolMessage(true),
		},
	},
}

func TestProcessorFun(t *testing.T) {
	for _, tc := range testCases {
		context := at.SetupContext(tc, inputsCfg, outputsCfg)
		err := ProcessorFun(context)
		assert.Nil(t, err)
		at.CompareOutputsData(t, context, tc)
	}
}
