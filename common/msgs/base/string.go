package base

import (
	"encoding/json"
	"fmt"
	"github.com/tombenke/axon-go/common/msgs"
	"github.com/tombenke/axon-go/common/msgs/common"
	"time"
)

const (
	// StringTypeName is the printable name of the `String` message-type
	StringTypeName = "base/String"
)

func init() {
	msgs.RegisterMessageType(StringTypeName, []msgs.Representation{msgs.JSONRepresentation}, func() msgs.Message {
		return NewStringMessage("")
	})
}

// String represents the structure of the messages emitted or consumed by the string-type sensors and actuators.
type String struct {
	Header common.Header
	Body   common.StringBody
}

// GetType returns with the printable name of the `String` message-type
func (msg *String) GetType() string {
	return StringTypeName
}

// Encode returns with the `String` message content in a representation format selected by `representation`
func (msg *String) Encode(representation msgs.Representation) (results []byte) {
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
func (msg *String) Decode(representation msgs.Representation, content []byte) error {
	switch representation {
	case msgs.JSONRepresentation:
		return json.Unmarshal(content, msg)
	default:
		panic(fmt.Errorf("Decode error: unknown representational format '%s'", representation))
	}
}

// JSON returns with the `String` message content in JSON representation format
func (msg *String) JSON() []byte {
	jsonBytes, err := json.Marshal(*msg)
	if err != nil {
		panic(err)
	}
	return jsonBytes
}

// String returns with the `String` message content in JSON format string
func (msg *String) String() string {
	jsonBytes, err := json.Marshal(*msg)
	if err != nil {
		panic(err)
	}
	return string(jsonBytes)
}

// ParseJSON parses the JSON representation of a `String` messages from the `jsonBytes` argument.
func (msg *String) ParseJSON(jsonBytes []byte) error {
	return json.Unmarshal(jsonBytes, msg)
}

// NewStringMessage returns with a new `String` message. The header will contain the current time in `Nanoseconds` precision.
func NewStringMessage(data string) msgs.Message {
	return NewStringMessageAt(data, time.Now().UnixNano(), "ns")
}

// NewStringMessageAt returns with a new `String` message. The header will contain the `at` time in `withPrecision` precision.
func NewStringMessageAt(data string, at int64, withPrecision common.TimePrecision) msgs.Message {
	var msg String
	msg.Header = common.NewHeaderAt(at, withPrecision)
	msg.Body.Data = data
	return &msg
}
