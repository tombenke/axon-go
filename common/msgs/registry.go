package msgs

import (
	"fmt"
)

// registry keeps track of the registered message types
// as well as of the representation formats can be used to a specific message type.
// It also holds funtions to each message type that return with the default value of the corresponding type.
// The message-types need to be registered before use this registry.
// It can happen typically in the `init()` function of the implementation of the message-types.
var registry map[string]MessageTypeDescriptor = make(map[string]MessageTypeDescriptor)

// MessageTypeDescriptor describes one message-type registered into the `registry`.
type MessageTypeDescriptor struct {
	// The name of the message-type
	Type string
	// The map of representation formats the message-type supports
	Representations map[Representation]bool
	// Function to get the default message value of the message-type
	GetDefaultMessageFun func() Message
}

// RegisterMessageType registers a specific message-type into the central registry
func RegisterMessageType(Type string, Representations []Representation, GetDefaultMessageFun func() Message) {
	if _, isPresent := registry[Type]; isPresent {

		errorString := fmt.Sprintf("The '%s' message type has already been registered yet!", Type)
		panic(errorString)
	}
	rmap := make(map[Representation]bool)
	for _, r := range Representations {
		rmap[r] = true
	}
	registry[Type] = MessageTypeDescriptor{Type, rmap, GetDefaultMessageFun}
}

// GetDefaultMessageByType returns with the default message value of the `Type` message-type
func GetDefaultMessageByType(Type string) Message {
	if _, isPresent := registry[Type]; isPresent {
		return registry[Type].GetDefaultMessageFun()
	}

	errorString := fmt.Sprintf("The '%s' message type has not been registered!", Type)
	panic(errorString)
}

// IsMessageTypeRegistered returns true if `Type` message-type is registered, othewise returns with false.
func IsMessageTypeRegistered(Type string) bool {
	if _, isPresent := registry[Type]; isPresent {
		return true
	}
	return false
}

// DoesMessageTypeImplementsRepresentation returns true if `Type` message-type is registered,
// ant it has implementation for Encoding and Decoding the `Representation` format, othewise returns with false.
func DoesMessageTypeImplementsRepresentation(Type string, Representation Representation) bool {
	if t, isTypePresent := registry[Type]; isTypePresent {
		if _, isRepPresent := t.Representations[Representation]; isRepPresent {
			return true
		}
	}
	return false
}
