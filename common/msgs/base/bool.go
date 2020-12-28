package base

import (
	"encoding/json"
	"fmt"
	"github.com/tombenke/axon-go/common/msgs"
	"github.com/tombenke/axon-go/common/msgs/common"
	"time"
)

const (
	// BoolTypeName is the printable name of the `Bool` message-type
	BoolTypeName = "base/Bool"
)

func init() {
	msgs.RegisterMessageType(BoolTypeName, []msgs.Representation{msgs.JSONRepresentation}, func() msgs.Message {
		return NewBoolMessage(false)
	})
}

// Bool represents the structure of the messages emitted or consumed by the boolean-type sensors and actuators.
type Bool struct {
	Header common.Header
	Body   common.BoolBody
}

// GetType returns with the printable name of the `Bool` message-type
func (msg *Bool) GetType() string {
	return BoolTypeName
}

// Encode returns with the `Bool` message content in a representation format selected by `representation`
func (msg *Bool) Encode(representation msgs.Representation) (results []byte) {
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
func (msg *Bool) Decode(representation msgs.Representation, content []byte) error {
	switch representation {
	case msgs.JSONRepresentation:
		return json.Unmarshal(content, msg)
	default:
		panic(fmt.Errorf("Decode error: unknown representational format '%s'", representation))
	}
}

// JSON returns with the `Bool` message content in JSON representation format
func (msg *Bool) JSON() []byte {
	jsonBytes, err := json.Marshal(*msg)
	if err != nil {
		panic(err)
	}
	return jsonBytes
}

// String returns with the `Bool` message content in JSON format string
func (msg *Bool) String() string {
	jsonBytes, err := json.Marshal(*msg)
	if err != nil {
		panic(err)
	}
	return string(jsonBytes)
}

// ParseJSON parses the JSON representation of a `Bool` messages from the `jsonBytes` argument.
func (msg *Bool) ParseJSON(jsonBytes []byte) error {
	return json.Unmarshal(jsonBytes, msg)
}

// NewBoolMessage returns with a new `Bool` message. The header will contain the current time in `Nanoseconds` precision.
func NewBoolMessage(data bool) msgs.Message {
	return NewBoolMessageAt(data, time.Now().UnixNano(), "ns")
}

// NewBoolMessageAt returns with a new `Bool` message. The header will contain the `at` time in `withPrecision` precision.
func NewBoolMessageAt(data bool, at int64, withPrecision common.TimePrecision) msgs.Message {
	var msg Bool
	msg.Header = common.NewHeaderAt(at, withPrecision)
	msg.Body.Data = data
	return &msg
}
