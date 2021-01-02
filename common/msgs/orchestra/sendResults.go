package orchestra

import (
	"encoding/json"
	"fmt"
	"github.com/tombenke/axon-go/common/msgs"
	"github.com/tombenke/axon-go/common/msgs/common"
	"time"
)

const (
	// SendResultsTypeName is the printable name of the `SendResults` message-type
	SendResultsTypeName = "orchestra/SendResults"
)

func init() {
	msgs.RegisterMessageType(SendResultsTypeName, []msgs.Representation{msgs.JSONRepresentation}, func() msgs.Message {
		return NewSendResultsMessage()
	})
}

// SendResults represents the structure of the `send-results` message that is usually sent
// by the orchestrator to the actors in order to trigger the sending of the results of processing.
type SendResults struct {
	Header common.Header
	Body   common.EmptyBody
}

// GetType returns with the printable name of the `SendResults` message-type
func (msg *SendResults) GetType() string {
	return SendResultsTypeName
}

// Encode returns with the `SendResults` message content in a representation format selected by `representation`
func (msg *SendResults) Encode(representation msgs.Representation) (results []byte) {
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
func (msg *SendResults) Decode(representation msgs.Representation, content []byte) error {
	switch representation {
	case msgs.JSONRepresentation:
		return json.Unmarshal(content, msg)
	default:
		panic(fmt.Errorf("Decode error: unknown representational format '%s'", representation))
	}
}

// JSON returns with the `SendResults` message content in JSON representation format
func (msg *SendResults) JSON() []byte {
	jsonBytes, err := json.Marshal(*msg)
	if err != nil {
		panic(err)
	}
	return jsonBytes
}

// String returns with the `SendResults` message content in JSON format string
func (msg *SendResults) String() string {
	jsonBytes, err := json.Marshal(*msg)
	if err != nil {
		panic(err)
	}
	return string(jsonBytes)
}

// ParseJSON parses the JSON representation of a `SendResults` messages from the `jsonBytes` argument.
func (msg *SendResults) ParseJSON(jsonBytes []byte) error {
	return json.Unmarshal(jsonBytes, msg)
}

// NewSendResultsMessage returns with a new `SendResults` message. The header will contain the current time in `Nanoseconds` precision.
func NewSendResultsMessage() msgs.Message {
	return NewSendResultsMessageAt(time.Now().UnixNano(), "ns")
}

// NewSendResultsMessageAt returns with a new `SendResults` message. The header will contain the `at` time in `withPrecision` precision.
func NewSendResultsMessageAt(at int64, withPrecision common.TimePrecision) msgs.Message {
	var msg SendResults
	msg.Header = common.NewHeaderAt(at, withPrecision)
	return &msg
}
