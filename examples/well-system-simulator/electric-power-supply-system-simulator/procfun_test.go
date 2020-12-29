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
		Name:           "max-power",
		Type:           "base/Float64",
		Representation: "application/json",
		Channel:        "",
	}, Default: `{"Body": {"Data": 2000.0}}`},
	config.In{IO: config.IO{
		Name:           "power-need",
		Type:           "base/Float64",
		Representation: "application/json",
		Channel:        "well-pump-relay.electric-power-need",
	}, Default: ""},
}

var outputsCfg = config.Outputs{
	config.Out{IO: config.IO{
		Name:           "power-output",
		Type:           "base/Float64",
		Representation: "application/json",
		Channel:        "well-pump-relay.electric-power-input",
	}},
}

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
