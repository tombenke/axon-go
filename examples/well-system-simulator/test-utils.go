package main

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/tombenke/axon-go/common/actor/processor"
	"github.com/tombenke/axon-go/common/config"
	"github.com/tombenke/axon-go/common/io"
	"github.com/tombenke/axon-go/common/msgs"
	"reflect"
	"testing"
)

var Logger logrus.Logger

type TestCase struct {
	Inputs  TestCaseMsgs
	Outputs TestCaseMsgs
}

type TestCases []TestCase
type TestCaseMsgs map[string]msgs.Message

func CompareOutputsData(t *testing.T, ctx processor.Context, tc TestCase) {
	for portName, _ := range tc.Outputs {
		expected := tc.Outputs[portName]
		actual := ctx.Outputs.GetOutputMessage(portName)
		assert.Equal(t, reflect.TypeOf(actual), reflect.TypeOf(expected))

		expectedDataValue := reflect.ValueOf(expected).Elem().FieldByName("Body").FieldByName("Data")
		actualDataValue := reflect.ValueOf(actual).Elem().FieldByName("Body").FieldByName("Data")

		switch expectedDataValue.Kind() {
		case reflect.Bool:
			assert.Equal(t, actualDataValue.Bool(), expectedDataValue.Bool())
		case reflect.Int64:
			assert.Equal(t, actualDataValue.Int(), expectedDataValue.Int())
		case reflect.Float64:
			assert.Equal(t, actualDataValue.Float(), expectedDataValue.Float())
		case reflect.String:
			assert.Equal(t, actualDataValue.String(), expectedDataValue.String())
		}
	}
}

func SetInputs(inputs *io.Inputs, inputMsgs TestCaseMsgs) {
	for portName, _ := range inputMsgs {
		(*inputs).SetInputMessage(portName, inputMsgs[portName])
	}
}

func SetupPorts(inputsCfg config.Inputs, outputsCfg config.Outputs) (io.Inputs, io.Outputs) {
	// Setup the input ports
	inputs := io.NewInputs(inputsCfg)

	// Setup the output ports
	outputs := io.NewOutputs(outputsCfg)

	return inputs, outputs
}

func SetupContext(tc TestCase, inputsCfg config.Inputs, outputsCfg config.Outputs) processor.Context {
	inputs, outputs := SetupPorts(inputsCfg, outputsCfg)
	SetInputs(&inputs, tc.Inputs)
	context := processor.NewContext(&Logger, inputs, outputs)
	return context
}
