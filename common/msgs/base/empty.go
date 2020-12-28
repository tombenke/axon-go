package base

import (
	"encoding/json"
	"fmt"
	"github.com/tombenke/axon-go/common/msgs"
	"github.com/tombenke/axon-go/common/msgs/common"
	"time"
)

const (
	// EmptyTypeName is the printable name of the `Empty` message-type
	EmptyTypeName = "base/Empty"
)

func init() {
	msgs.RegisterMessageType(EmptyTypeName, []msgs.Representation{msgs.JSONRepresentation}, func() msgs.Message {
		return NewEmptyMessage()
	})
}

// Empty represents the structure of the empty message.
type Empty struct {
	Header common.Header
	Body   common.EmptyBody
}

// GetType returns with the printable name of the `Empty` message-type
func (msg *Empty) GetType() string {
	return EmptyTypeName
}

// Encode returns with the `Empty` message content in a representation format selected by `representation`
func (msg *Empty) Encode(representation msgs.Representation) (results []byte) {
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
func (msg *Empty) Decode(representation msgs.Representation, content []byte) error {
	switch representation {
	case msgs.JSONRepresentation:
		return json.Unmarshal(content, msg)
	default:
		panic(fmt.Errorf("Decode error: unknown representational format '%s'", representation))
	}
}

// JSON returns with the `Empty` message content in JSON representation format
func (msg *Empty) JSON() []byte {
	jsonBytes, err := json.Marshal(*msg)
	if err != nil {
		panic(err)
	}
	return jsonBytes
}

// String returns with the `Empty` message content in JSON format string
func (msg *Empty) String() string {
	jsonBytes, err := json.Marshal(*msg)
	if err != nil {
		panic(err)
	}
	return string(jsonBytes)
}

// ParseJSON parses the JSON representation of a `Empty` messages from the `jsonBytes` argument.
func (msg *Empty) ParseJSON(jsonBytes []byte) error {
	return json.Unmarshal(jsonBytes, msg)
}

// NewEmptyMessage returns with a new `Empty` message. The header will contain the current time in `Nanoseconds` precision.
func NewEmptyMessage() msgs.Message {
	return NewEmptyMessageAt(time.Now().UnixNano(), "ns")
}

// NewEmptyMessageAt returns with a new `Empty` message. The header will contain the `at` time in `withPrecision` precision.
func NewEmptyMessageAt(at int64, withPrecision common.TimePrecision) msgs.Message {
	var msg Empty
	msg.Header = common.NewHeaderAt(at, withPrecision)
	return &msg
}
