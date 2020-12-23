package processor

import (
	"github.com/tombenke/axon-go/common/io"
	"github.com/tombenke/axon-go/common/msgs"
)

// ProcCtxIface defines the interface for the Processor to access to the input and output ports of the actor
type ProcCtxIface interface {
	io.Handler
	// TODO: Add Logger, and Config
}

// ProcCtx is the structure of the Processor context. It holds the actual messages arrived through the `Inputs` ports
// as well as the messages will be emitted through the `Outputs` ports.
type ProcCtx struct {
	Inputs  io.Inputs
	Outputs io.Outputs
	// TODO: Add Logger, and Config
}

// GetInputMessage returns the latest input message arrived to the input port selected by its `name`.
func (ctx ProcCtx) GetInputMessage(name string) (msgs.Message, error) {
	return ctx.Inputs.GetInputMessage(name)
}

// SetOutputMessage sets the `outMsg` message to be emitted via the output port selected by its `name`.
func (ctx ProcCtx) SetOutputMessage(name string, outMsg msgs.Message) error {
	return ctx.Outputs.SetOutputMessage(name, outMsg)
}
