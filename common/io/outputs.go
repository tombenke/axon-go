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
func NewOutputs(outputsCfg config.Outputs) Outputs {
	outputs := make(Outputs)
	for _, o := range outputsCfg {
		Name := o.IO.Name
		Type := o.IO.Type
		Repr := msgs.Representation(o.IO.Representation)
		Chan := o.IO.Channel
		if !msgs.IsMessageTypeRegistered(Type) {
			errorString := fmt.Sprintf("The '%s' message type has not been registered!", Type)
			panic(errorString)
		}
		if !msgs.DoesMessageTypeImplementsRepresentation(Type, Repr) {
			errorString := fmt.Sprintf("'%s' message-type does not implement codec for '%s' representation format", Type, Repr)
			panic(errorString)
		}
		outputs[Name] = Output{IO{Name: Name, Type: Type, Representation: Repr, Channel: Chan}}
	}
	return outputs
}
