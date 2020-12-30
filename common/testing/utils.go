package testing

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

// Logger is a global logger for testing
var Logger logrus.Logger

// TestCase holds the input and output messages of one test case for a processor function.
// The inputs is put onto the context that the processor function is using
// during the execution of one specific test-case.
// The outputs hold the expected messages the processor function has to place
// during the test to the outputs port via the context.
type TestCase struct {
	Inputs  TestCaseMsgs
	Outputs TestCaseMsgs
}

// TestCases holds an array of TestCases for a specific processor function.
type TestCases []TestCase

// TestCaseMsgs is a type that TestCase descriptor structures use for
// both the inputs and outputs messages.
// This is a map of messages where the portname is the key.
type TestCaseMsgs map[string]msgs.Message

// CompareOutputsData is a helper function that compares the `Data` properties
// of output messages between the context and the `TestCase` descriptor.
// The function is called after the execution of a specific `TestCase`.
// The output messages are taken by their port name from the context,
// and from the `TestCase` descriptor.
//
// NOTE: The output messages to compare can not have any arbitrary structure.
// This function is made for those output messages only, that has a `Body.Data` property,
// which can be derived into a single base type provided for `reflect.Value` fields.
func CompareOutputsData(t *testing.T, ctx processor.Context, tc TestCase) {
	for portName := range tc.Outputs {
		expected := tc.Outputs[portName]
		actual := ctx.Outputs.GetMessage(portName)
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

// SetInputs sets the input port values according to the content of the `inputMsgs` test case descriptor.
func SetInputs(inputs *io.Inputs, inputMsgs TestCaseMsgs) {
	for portName := range inputMsgs {
		(*inputs).SetMessage(portName, inputMsgs[portName])
	}
}

// SetupPorts Initializes the input and output ports according to their configurations.
func SetupPorts(inputsCfg config.Inputs, outputsCfg config.Outputs) (io.Inputs, io.Outputs) {
	// Setup the input ports
	inputs := io.NewInputs(inputsCfg)

	// Setup the output ports
	outputs := io.NewOutputs(outputsCfg)

	return inputs, outputs
}

// SetupContext creates a new `processor.Context` for the processor function of the actor.
// This context will contain the inputs, and outputs as they are defined by the port configuration,
// and by the `tc` test case. It also provides a Logger.
func SetupContext(tc TestCase, inputsCfg config.Inputs, outputsCfg config.Outputs) processor.Context {
	inputs, outputs := SetupPorts(inputsCfg, outputsCfg)
	SetInputs(&inputs, tc.Inputs)
	context := processor.NewContext(&Logger, inputs, outputs)
	return context
}
