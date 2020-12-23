package base

import (
	"encoding/json"
	"fmt"
	"github.com/tombenke/axon-go/common/msgs"
	"github.com/tombenke/axon-go/common/msgs/common"
	"time"
)

const (
	// Int64TypeName is the printable name of the `Int64` message-type
	Int64TypeName = "base/Int64"
)

// Int64 represents the structure of the messages emitted or consumed by the int64 sensors and actuators.
type Int64 struct {
	Header common.Header
	Body   common.Int64Body
}

// GetType returns with the printable name of the `Int64` message-type
func (msg *Int64) GetType() string {
	return Int64TypeName
}

// Encode returns with the `Int64` message content in a representation format selected by `representation`
func (msg *Int64) Encode(representation msgs.Representation) (results []byte) {
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
func (msg *Int64) Decode(representation msgs.Representation, content []byte) error {
	switch representation {
	case msgs.JSONRepresentation:
		return json.Unmarshal(content, msg)
	default:
		panic(fmt.Errorf("Decode error: unknown representational format '%s'", representation))
	}
}

// JSON returns with the `Int64` message content in JSON representation format
func (msg *Int64) JSON() []byte {
	jsonBytes, err := json.Marshal(*msg)
	if err != nil {
		panic(err)
	}
	return jsonBytes
}

// String returns with the `Int64` message content in JSON format string
func (msg *Int64) String() string {
	jsonBytes, err := json.Marshal(*msg)
	if err != nil {
		panic(err)
	}
	return string(jsonBytes)
}

// ParseJSON parses the JSON representation of a `Int64` messages from the `jsonBytes` argument.
func (msg *Int64) ParseJSON(jsonBytes []byte) error {
	return json.Unmarshal(jsonBytes, msg)
}

// NewInt64Message returns with a new `Int64` message. The header will contain the current time in `Nanoseconds` precision.
func NewInt64Message(data int64) msgs.Message {
	return NewInt64MessageAt(data, time.Now().UnixNano(), "ns")
}

// NewInt64MessageAt returns with a new `Int64` message. The header will contain the `at` time in `withPrecision` precision.
func NewInt64MessageAt(data int64, at int64, withPrecision common.TimePrecision) msgs.Message {
	var msg Int64
	msg.Header = common.NewHeaderAt(at, withPrecision)
	msg.Body.Data = data
	return &msg
}
