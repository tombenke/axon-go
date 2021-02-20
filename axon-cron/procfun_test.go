package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/tombenke/axon-go/common/actor/processor"
	"github.com/tombenke/axon-go/common/msgs/base"
	at "github.com/tombenke/axon-go/common/testing"
	"testing"
)

var testCases at.TestCases = at.TestCases{
	at.TestCase{
		Outputs: at.TestCaseMsgs{
			"output": base.NewAnyMessage(map[string]interface{}{
				"time": int64(1612693423000521770),
				"meta": map[string]interface{}{
					"timePrecision": "ns",
				}}),
		},
	},
}

func TestProcessorFun(t *testing.T) {
	for _, tc := range testCases {
		config := builtInConfig()
		context := processor.SetupContext(tc, config.Node.Ports.Inputs, config.Node.Ports.Outputs)
		err := getProcessorFun(config)(context)
		assert.Nil(t, err)
		fmt.Printf("\nOutputs: %v\n", context.Outputs)
		//processor.CompareOutputsData(t, context.Outputs, tc)
	}
}
