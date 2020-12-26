package base

import (
	"encoding/json"
	"fmt"
	"github.com/tombenke/axon-go/common/msgs"
)

const (
	// AnyTypeName is the printable name of the `Any` message-type
	AnyTypeName = "base/Any"
)

// Any represents the structure of a generic message that may contain anything
type Any map[string]interface{}

// GetType returns with the printable name of the `Any` message-type
func (msg *Any) GetType() string {
	return AnyTypeName
}

// Encode returns with the `Any` message content in a representation format selected by `representation`
func (msg *Any) Encode(representation msgs.Representation) (results []byte) {
	switch representation {
	case msgs.JSONRepresentation:
		var err error
		results, err = json.Marshal(*msg)
		if err != nil {
			panic(err)
		}
	default:
		panic(fmt.Errorf("Encode error: unknown representational format '%s'", representation))
	}
	return results
}

// Decode parses the `content` using the selected `representation` format
func (msg *Any) Decode(representation msgs.Representation, content []byte) error {
	switch representation {
	case msgs.JSONRepresentation:
		return json.Unmarshal(content, msg)
	default:
		panic(fmt.Errorf("Decode error: unknown representational format '%s'", representation))
	}
}

// JSON returns with the `Any` message content in JSON representation format
func (msg *Any) JSON() []byte {
	jsonBytes, err := json.Marshal(*msg)
	if err != nil {
		panic(err)
	}
	return jsonBytes
}

// String returns with the `Any` message content in JSON format string
func (msg *Any) String() string {
	jsonBytes, err := json.Marshal(*msg)
	if err != nil {
		panic(err)
	}
	return string(jsonBytes)
}

// ParseJSON parses the JSON representation of a `Any` messages from the `jsonBytes` argument.
func (msg *Any) ParseJSON(jsonBytes []byte) error {
	return json.Unmarshal(jsonBytes, msg)
}

// NewAnyMessage returns with a new `Any` message. The header will contain the current time in `Nanoseconds` precision.
func NewAnyMessage(data map[string]interface{}) msgs.Message {
	var msg Any
	msg = data
	return &msg
}
