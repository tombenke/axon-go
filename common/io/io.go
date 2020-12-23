package io

import (
	"github.com/tombenke/axon-go/common/msgs"
)

// IO is a generic port descriptor structure,
// that contains the common properties of both the input and output ports
type IO struct {
	Name    string       // The name of the inpu/output port
	Type    string       // The message-type of the io port
	Message msgs.Message // The actual message the io port holds
}

// Handler is an interface for both the input and output type ports
type Handler interface {
	InputsHandler
	OutputsHandler
}
