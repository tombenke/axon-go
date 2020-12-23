package base

import (
	"encoding/json"
	"fmt"
	"github.com/tombenke/axon-go/common/msgs"
	"github.com/tombenke/axon-go/common/msgs/common"
	"time"
)

const (
	// Float64TypeName is the printable name of the `Float64` message-type
	Float64TypeName = "base/Float64"
)

// Float64 represents the structure of the messages emitted or consumed by the float64 sensors and actuators.
type Float64 struct {
	Header common.Header
	Body   common.Float64Body
}

// GetType returns with the printable name of the `Float64` message-type
func (msg *Float64) GetType() string {
	return Float64TypeName
}

// Encode returns with the `Float64` message content in a representation format selected by `representation`
func (msg *Float64) Encode(representation msgs.Representation) (results []byte) {
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
func (msg *Float64) Decode(representation msgs.Representation, content []byte) error {
	switch representation {
	case msgs.JSONRepresentation:
		return json.Unmarshal(content, msg)
	default:
		panic(fmt.Errorf("Decode error: unknown representational format '%s'", representation))
	}
}

// JSON returns with the `Float64` message content in JSON representation format
func (msg *Float64) JSON() []byte {
	jsonBytes, err := json.Marshal(*msg)
	if err != nil {
		panic(err)
	}
	return jsonBytes
}

// String returns with the `Float64` message content in JSON format string
func (msg *Float64) String() string {
	jsonBytes, err := json.Marshal(*msg)
	if err != nil {
		panic(err)
	}
	return string(jsonBytes)
}

// ParseJSON parses the JSON representation of a `Float64` messages from the `jsonBytes` argument.
func (msg *Float64) ParseJSON(jsonBytes []byte) error {
	return json.Unmarshal(jsonBytes, msg)
}

// NewFloat64Message returns with a new `Float64` message. The header will contain the current time in `Nanoseconds` precision.
func NewFloat64Message(data float64) msgs.Message {
	return NewFloat64MessageAt(data, time.Now().UnixNano(), "ns")
}

// NewFloat64MessageAt returns with a new `Float64` message. The header will contain the `at` time in `withPrecision` precision.
func NewFloat64MessageAt(data float64, at int64, withPrecision common.TimePrecision) msgs.Message {
	var msg Float64
	msg.Header = common.NewHeaderAt(at, withPrecision)
	msg.Body.Data = data
	return &msg
}
