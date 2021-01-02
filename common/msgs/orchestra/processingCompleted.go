package orchestra

import (
	"encoding/json"
	"fmt"
	"github.com/tombenke/axon-go/common/msgs"
	"github.com/tombenke/axon-go/common/msgs/common"
	"time"
)

const (
	// ProcessingCompletedTypeName is the printable name of the `ProcessingCompleted` message-type
	ProcessingCompletedTypeName = "orchestra/ProcessingCompleted"
)

func init() {
	msgs.RegisterMessageType(ProcessingCompletedTypeName, []msgs.Representation{msgs.JSONRepresentation}, func() msgs.Message {
		return NewProcessingCompletedMessage("")
	})
}

// ProcessingCompleted represents the structure of the `processing-completed` status message
// message that the actor is usually sent when it is ready to send the results of the processing.
// The `Body.Data` property holds the name of the actor that sends the response.
type ProcessingCompleted struct {
	Header common.Header
	Body   common.StringBody
}

// GetType returns with the printable name of the `ProcessingCompleted` message-type
func (msg *ProcessingCompleted) GetType() string {
	return ProcessingCompletedTypeName
}

// Encode returns with the `ProcessingCompleted` message content in a representation format selected by `representation`
func (msg *ProcessingCompleted) Encode(representation msgs.Representation) (results []byte) {
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
func (msg *ProcessingCompleted) Decode(representation msgs.Representation, content []byte) error {
	switch representation {
	case msgs.JSONRepresentation:
		return json.Unmarshal(content, msg)
	default:
		panic(fmt.Errorf("Decode error: unknown representational format '%s'", representation))
	}
}

// JSON returns with the `ProcessingCompleted` message content in JSON representation format
func (msg *ProcessingCompleted) JSON() []byte {
	jsonBytes, err := json.Marshal(*msg)
	if err != nil {
		panic(err)
	}
	return jsonBytes
}

// ProcessingCompleted returns with the `ProcessingCompleted` message content in JSON format string
func (msg *ProcessingCompleted) String() string {
	jsonBytes, err := json.Marshal(*msg)
	if err != nil {
		panic(err)
	}
	return string(jsonBytes)
}

// ParseJSON parses the JSON representation of a `ProcessingCompleted` messages from the `jsonBytes` argument.
func (msg *ProcessingCompleted) ParseJSON(jsonBytes []byte) error {
	return json.Unmarshal(jsonBytes, msg)
}

// NewProcessingCompletedMessage returns with a new `ProcessingCompleted` message. The header will contain the current time in `Nanoseconds` precision.
func NewProcessingCompletedMessage(data string) msgs.Message {
	return NewProcessingCompletedMessageAt(data, time.Now().UnixNano(), "ns")
}

// NewProcessingCompletedMessageAt returns with a new `ProcessingCompleted` message. The header will contain the `at` time in `withPrecision` precision.
func NewProcessingCompletedMessageAt(data string, at int64, withPrecision common.TimePrecision) msgs.Message {
	var msg ProcessingCompleted
	msg.Header = common.NewHeaderAt(at, withPrecision)
	msg.Body.Data = data
	return &msg
}
