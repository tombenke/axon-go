package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/tombenke/axon-go/common/config"
	"github.com/tombenke/axon-go/common/msgs/base"
	at "github.com/tombenke/axon-go/common/testing"
	"testing"
)

var inputsCfg = config.Inputs{
	config.In{IO: config.IO{
		Name:           "reference-water-level",
		Type:           "base/Float64",
		Representation: "application/json",
		Channel:        "",
	}, Default: `{"Body": {"Data": 0.75}}`},
	config.In{IO: config.IO{
		Name:           "water-level",
		Type:           "base/Float64",
		Representation: "application/json",
		Channel:        "well-water-buffer-tank-level",
	}, Default: ""},
}

var outputsCfg = config.Outputs{
	config.Out{IO: config.IO{
		Name:           "water-level-state",
		Type:           "base/Bool",
		Representation: "application/json",
		Channel:        "buffer-water-tank-upper-level-state",
	}},
}

var testCases at.TestCases = at.TestCases{
	at.TestCase{
		at.TestCaseMsgs{
			"reference-water-level": base.NewFloat64Message(0.75),
			"water-level":           base.NewFloat64Message(0.0),
		},
		at.TestCaseMsgs{
			"water-level-state": base.NewBoolMessage(false),
		},
	},
	at.TestCase{
		at.TestCaseMsgs{
			"reference-water-level": base.NewFloat64Message(0.75),
			"water-level":           base.NewFloat64Message(0.75),
		},
		at.TestCaseMsgs{
			"water-level-state": base.NewBoolMessage(true),
		},
	},
	at.TestCase{
		at.TestCaseMsgs{
			"reference-water-level": base.NewFloat64Message(0.75),
			"water-level":           base.NewFloat64Message(0.8),
		},
		at.TestCaseMsgs{
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
