package io

import (
	"errors"
	"fmt"
	"github.com/tombenke/axon-go/common/config"
	"github.com/tombenke/axon-go/common/msgs"
	"reflect"
)

// Output holds the data of an output port of the actor
type Output struct {
	IO
}

// Outputs holds a map of the the output ports of the actor. The key is the name of the port.
type Outputs map[string]Output

// OutputsHandler declares the methods to the management of the output ports
type OutputsHandler interface {
	SetOutputMessage(string, msgs.Message) error
}

// SetOutputMessage sets the message to emit via the output port selected by the `name` parameter
func (outputs *Outputs) SetOutputMessage(name string, outMsg msgs.Message) error {
	if _, ok := (*outputs)[name]; ok {
		fmt.Printf("outputs: %s", reflect.TypeOf((*outputs)[name]))
		(*outputs)[name] = Output{IO{Name: name, Type: outMsg.GetType(), Message: outMsg}}
		return nil
	}
	return errors.New("There is no output port named to " + name)
}

// NewOutputs creates a new Outputs map based on the config parameters
func NewOutputs(outputs config.Outputs) (Outputs, error) {
	return Outputs{}, nil
}
