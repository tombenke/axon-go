package io

import (
	"errors"
	"github.com/tombenke/axon-go/common/msgs"
)

// Input holds the data of an input port of the actor
type Input struct {
	IO
	DefaultMessage msgs.Message
}

// Inputs holds a map of the the input ports of the actor. The key is the name of the port.
type Inputs map[string]Input

// InputsHandler declares the methods to the management of the input ports
type InputsHandler interface {
	GetInputMessage(string) (msgs.Message, error)
}

// GetInputMessage returns the last message received via the inputput port selected by the `name` parameter
func (inputs Inputs) GetInputMessage(name string) (msgs.Message, error) {

	if input, ok := inputs[name]; ok {
		return input.Message, nil
	}
	return nil, errors.New("There is no input port named to " + name)
}
