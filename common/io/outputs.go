package io

import (
	"fmt"
	"github.com/tombenke/axon-go/common/config"
	"github.com/tombenke/axon-go/common/msgs"
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

// GetOutputMessage returns the last message set to the output port for sending selected by the `name` parameter
func (outputs Outputs) GetOutputMessage(name string) msgs.Message {

	if output, ok := outputs[name]; ok {
		return output.Message
	}
	errorMessage := fmt.Sprintf("There is no output port named to '%s'", name)
	panic(errorMessage)
}

// SetOutputMessage sets the message to emit via the output port selected by the `name` parameter
func (outputs *Outputs) SetOutputMessage(name string, outMsg msgs.Message) {
	if _, ok := (*outputs)[name]; !ok {
		errorMessage := fmt.Sprintf("'%s' port does not exist, so can not set message to it.", name)
		panic(errorMessage)
	}

	outMsgType := outMsg.GetType()
	portMsgType := string((*outputs)[name].Type)
	if outMsgType != portMsgType {
		errorMessage := fmt.Sprintf("'%s' message-type mismatch to port's '%s' message-type.", outMsgType, portMsgType)
		panic(errorMessage)
	}

	(*outputs)[name] = Output{IO: IO{Name: name, Type: outMsgType, Message: outMsg}}
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
