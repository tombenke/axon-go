package orchestra

import (
	"encoding/json"
	"fmt"
	"github.com/tombenke/axon-go/common/msgs"
	"github.com/tombenke/axon-go/common/msgs/common"
	"time"
)

const (
	// SendingCompletedTypeName is the printable name of the `SendingCompleted` message-type
	SendingCompletedTypeName = "orchestra/SendingCompleted"
)

func init() {
	msgs.RegisterMessageType(SendingCompletedTypeName, []msgs.Representation{msgs.JSONRepresentation}, func() msgs.Message {
		return NewSendingCompletedMessage("")
	})
}

// SendingCompleted represents the structure of the `sending-completed` status message
// message that the actor is usually sent when it completed the sending of output messages to their
// target channels.
// The `Body.Data` property holds the name of the actor that sends the response.
type SendingCompleted struct {
	Header common.Header
	Body   common.StringBody
}

// GetType returns with the printable name of the `SendingCompleted` message-type
func (msg *SendingCompleted) GetType() string {
	return SendingCompletedTypeName
}

// Encode returns with the `SendingCompleted` message content in a representation format selected by `representation`
func (msg *SendingCompleted) Encode(representation msgs.Representation) (results []byte) {
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
func (msg *SendingCompleted) Decode(representation msgs.Representation, content []byte) error {
	switch representation {
	case msgs.JSONRepresentation:
		return json.Unmarshal(content, msg)
	default:
		panic(fmt.Errorf("Decode error: unknown representational format '%s'", representation))
	}
}

// JSON returns with the `SendingCompleted` message content in JSON representation format
func (msg *SendingCompleted) JSON() []byte {
	jsonBytes, err := json.Marshal(*msg)
	if err != nil {
		panic(err)
	}
	return jsonBytes
}

// SendingCompleted returns with the `SendingCompleted` message content in JSON format string
func (msg *SendingCompleted) String() string {
	jsonBytes, err := json.Marshal(*msg)
	if err != nil {
		panic(err)
	}
	return string(jsonBytes)
}

// ParseJSON parses the JSON representation of a `SendingCompleted` messages from the `jsonBytes` argument.
func (msg *SendingCompleted) ParseJSON(jsonBytes []byte) error {
	return json.Unmarshal(jsonBytes, msg)
}

// NewSendingCompletedMessage returns with a new `SendingCompleted` message. The header will contain the current time in `Nanoseconds` precision.
func NewSendingCompletedMessage(data string) msgs.Message {
	return NewSendingCompletedMessageAt(data, time.Now().UnixNano(), "ns")
}

// NewSendingCompletedMessageAt returns with a new `SendingCompleted` message. The header will contain the `at` time in `withPrecision` precision.
func NewSendingCompletedMessageAt(data string, at int64, withPrecision common.TimePrecision) msgs.Message {
	var msg SendingCompleted
	msg.Header = common.NewHeaderAt(at, withPrecision)
	msg.Body.Data = data
	return &msg
}
