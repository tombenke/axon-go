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
		Name:           "relay-state",
		Type:           "base/Bool",
		Representation: "application/json",
		Channel:        "",
	}, Default: ""},
	config.In{IO: config.IO{
		Name:           "power-input",
		Type:           "base/Float64",
		Representation: "application/json",
		Channel:        "well-pump-relay.electric-power-input",
	}, Default: ""},
	config.In{IO: config.IO{
		Name:           "power-need",
		Type:           "base/Float64",
		Representation: "application/json",
		Channel:        "well-pump-power-need",
	}, Default: ""},
}

var outputsCfg = config.Outputs{
	config.Out{IO: config.IO{
		Name:           "power-output",
		Type:           "base/Float64",
		Representation: "application/json",
		Channel:        "well-pump-power",
	}},
	config.Out{IO: config.IO{
		Name:           "power-need",
		Type:           "base/Float64",
		Representation: "application/json",
		Channel:        "well-pump-relay.electric-power-need",
	}},
}

var testCases at.TestCases = at.TestCases{
	// Relay is OFF
	at.TestCase{
		at.TestCaseMsgs{
			"relay-state": base.NewBoolMessage(false),
			"power-input": base.NewFloat64Message(0.0),
			"power-need":  base.NewFloat64Message(0.0),
		},
		at.TestCaseMsgs{
			"power-output": base.NewFloat64Message(0.0),
			"power-need":   base.NewFloat64Message(0.0),
		},
	},
	at.TestCase{
		at.TestCaseMsgs{
			"relay-state": base.NewBoolMessage(false),
			"power-input": base.NewFloat64Message(10.0),
			"power-need":  base.NewFloat64Message(10.0),
		},
		at.TestCaseMsgs{
			"power-output": base.NewFloat64Message(0.0),
			"power-need":   base.NewFloat64Message(0.0),
		},
	},

	// Relay is ON
	at.TestCase{
		at.TestCaseMsgs{
			"relay-state": base.NewBoolMessage(true),
			"power-input": base.NewFloat64Message(0.0),
			"power-need":  base.NewFloat64Message(0.0),
		},
		at.TestCaseMsgs{
			"power-output": base.NewFloat64Message(0.0),
			"power-need":   base.NewFloat64Message(0.0),
		},
	},
	at.TestCase{
		at.TestCaseMsgs{
			"relay-state": base.NewBoolMessage(true),
			"power-input": base.NewFloat64Message(10.0),
			"power-need":  base.NewFloat64Message(10.0),
		},
		at.TestCaseMsgs{
			"power-output": base.NewFloat64Message(10.0),
			"power-need":   base.NewFloat64Message(10.0),
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
