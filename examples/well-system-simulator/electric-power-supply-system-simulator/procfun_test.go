package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/tombenke/axon-go/common/msgs/base"
	at "github.com/tombenke/axon-go/common/testing"
	"testing"
)

var testCases at.TestCases = at.TestCases{
	at.TestCase{
		at.TestCaseMsgs{
			"max-power":  base.NewFloat64Message(2000.0),
			"power-need": base.NewFloat64Message(0.0),
		},
		at.TestCaseMsgs{
			"power-output": base.NewFloat64Message(0.0),
		},
	},
	at.TestCase{
		at.TestCaseMsgs{
			"max-power":  base.NewFloat64Message(2000.0),
			"power-need": base.NewFloat64Message(1500.0),
		},
		at.TestCaseMsgs{
			"power-output": base.NewFloat64Message(1500.0),
		},
	},
	at.TestCase{
		at.TestCaseMsgs{
			"max-power":  base.NewFloat64Message(2000.0),
			"power-need": base.NewFloat64Message(2000.0),
		},
		at.TestCaseMsgs{
			"power-output": base.NewFloat64Message(2000.0),
		},
	},
	at.TestCase{
		at.TestCaseMsgs{
			"max-power":  base.NewFloat64Message(2000.0),
			"power-need": base.NewFloat64Message(4599.0),
		},
		at.TestCaseMsgs{
			"power-output": base.NewFloat64Message(2000.0),
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
