package processor

import (
	"github.com/sirupsen/logrus"
	"github.com/tombenke/axon-go/common/io"
	"github.com/tombenke/axon-go/common/msgs"
)

// Context is the structure of the Processor context.
// It holds the actual messages arrived through the `Inputs` ports
// as well as the messages will be emitted through the `Outputs` ports.
type Context struct {
	Inputs  io.Inputs
	Outputs io.Outputs
	Logger  *logrus.Logger
}

// GetInputMessage returns the latest input message arrived to the input port selected by its `name`.
func (ctx Context) GetInputMessage(name string) msgs.Message {
	return ctx.Inputs.GetInputMessage(name)
}

// SetOutputMessage sets the `outMsg` message to be emitted via the output port selected by its `name`.
func (ctx Context) SetOutputMessage(name string, outMsg msgs.Message) {
	ctx.Outputs.SetOutputMessage(name, outMsg)
}

// NewContext creates a new processor context object and returns with it
func NewContext(logger *logrus.Logger, inputs io.Inputs, outputs io.Outputs) Context {
	context := Context{Inputs: inputs, Outputs: outputs, Logger: logger}

	return context
}
