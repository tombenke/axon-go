package io

import (
	"fmt"
	"github.com/tombenke/axon-go/common/config"
	"github.com/tombenke/axon-go/common/msgs"
)

// Input holds the data of an input port of the actor
type Input struct {
	IO
	DefaultMessage msgs.Message
}

// Inputs holds a map of the the input ports of the actor. The key is the name of the port.
type Inputs map[string]Input

// GetMessage returns the last message received via the input port selected by the `name` parameter
func (inputs Inputs) GetMessage(name string) msgs.Message {

	if input, ok := inputs[name]; ok {
		return input.Message
	}
	errorMessage := fmt.Sprintf("There is no input port named to '%s'", name)
	panic(errorMessage)
}

// SetMessage sets the message that received via the input channel to the port selected by the `name` parameter
func (inputs *Inputs) SetMessage(name string, inMsg msgs.Message) {
	if _, ok := (*inputs)[name]; !ok {
		errorMessage := fmt.Sprintf("'%s' port does not exist, so can not set message to it.", name)
		panic(errorMessage)
	}

	inMsgType := inMsg.GetType()
	portMsgType := string((*inputs)[name].Type)
	if inMsgType != portMsgType {
		errorMessage := fmt.Sprintf("'%s' message-type mismatch to port's '%s' message-type.", inMsgType, portMsgType)
		panic(errorMessage)
	}

	(*inputs)[name] = Input{
		IO: IO{
			Name:           name,
			Type:           inMsgType,
			Representation: (*inputs)[name].Representation,
			Channel:        (*inputs)[name].Channel,
			Message:        inMsg,
		},
		DefaultMessage: (*inputs)[name].DefaultMessage,
	}
}

// NewInputs creates a new Inputs map based on the config parameters
func NewInputs(inputsCfg config.Inputs) Inputs {
	inputs := make(Inputs)
	for _, in := range inputsCfg {
		Name := in.IO.Name
		Type := in.IO.Type
		Repr := msgs.Representation(in.IO.Representation)
		Chan := in.IO.Channel

		// Validates if the message-type is registered
		if !msgs.IsMessageTypeRegistered(Type) {
			errorString := fmt.Sprintf("The '%s' message type has not been registered!", Type)
			panic(errorString)
		}

		// Validates if the representation format is supported
		if !msgs.DoesMessageTypeImplementsRepresentation(Type, Repr) {
			errorString := fmt.Sprintf("'%s' message-type does not implement codec for '%s' representation format", Type, Repr)
			panic(errorString)
		}

		// Determines the default value
		// In case the default value is empty, then use the original one defined by the message-type itself
		Default := msgs.GetDefaultMessageByType(Type)
		if in.Default != "" {
			// The default config value is not empty, so it should be a valid message in JSON format
			err := Default.Decode(msgs.JSONRepresentation, []byte(in.Default))
			if err != nil {
				panic(err)
			}
		}

		inputs[Name] = Input{IO: IO{Name: Name, Type: Type, Representation: Repr, Channel: Chan}, DefaultMessage: Default}
	}
	return inputs
}
