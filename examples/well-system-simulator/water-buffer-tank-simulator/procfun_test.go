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
			"dt": base.NewFloat64Message(1000),
		},
		Outputs: at.TestCaseMsgs{},
	},
	// TODO: Add test cases
}

func TestProcessorFun(t *testing.T) {
	for _, tc := range testCases {
		context := at.SetupContext(tc, inputsCfg, outputsCfg)
		err := ProcessorFun(context)
		assert.Nil(t, err)
		at.CompareOutputsData(t, context, tc)
	}
}
