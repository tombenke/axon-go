package io

import (
	"errors"
	"github.com/tombenke/axon-go/common/msgs"
)

type Input struct {
	IO
	DefaultMessage msgs.Message
}

type Inputs map[string]Input

type InputsHandler interface {
	GetInputMessage(string) (msgs.Message, error)
}

func (inputs Inputs) GetInputMessage(name string) (msgs.Message, error) {

	if input, ok := inputs[name]; ok {
		return input.Message, nil
	}
	return nil, errors.New("There is no input port named to " + name)
}
