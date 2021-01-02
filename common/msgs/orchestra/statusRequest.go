package orchestra

import (
	"encoding/json"
	"fmt"
	"github.com/tombenke/axon-go/common/msgs"
	"github.com/tombenke/axon-go/common/msgs/common"
	"time"
)

const (
	// StatusRequestTypeName is the printable name of the `StatusRequest` message-type
	StatusRequestTypeName = "orchestra/StatusRequest"
)

func init() {
	msgs.RegisterMessageType(StatusRequestTypeName, []msgs.Representation{msgs.JSONRepresentation}, func() msgs.Message {
		return NewStatusRequestMessage()
	})
}

// StatusRequest represents the structure of the `status-request` message that is usually sent
// by the orchestrator to get a `status-report` response from every actors.
type StatusRequest struct {
	Header common.Header
	Body   common.EmptyBody
}

// GetType returns with the printable name of the `StatusRequest` message-type
func (msg *StatusRequest) GetType() string {
	return StatusRequestTypeName
}

// Encode returns with the `StatusRequest` message content in a representation format selected by `representation`
func (msg *StatusRequest) Encode(representation msgs.Representation) (results []byte) {
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
func (msg *StatusRequest) Decode(representation msgs.Representation, content []byte) error {
	switch representation {
	case msgs.JSONRepresentation:
		return json.Unmarshal(content, msg)
	default:
		panic(fmt.Errorf("Decode error: unknown representational format '%s'", representation))
	}
}

// JSON returns with the `StatusRequest` message content in JSON representation format
func (msg *StatusRequest) JSON() []byte {
	jsonBytes, err := json.Marshal(*msg)
	if err != nil {
		panic(err)
	}
	return jsonBytes
}

// String returns with the `StatusRequest` message content in JSON format string
func (msg *StatusRequest) String() string {
	jsonBytes, err := json.Marshal(*msg)
	if err != nil {
		panic(err)
	}
	return string(jsonBytes)
}

// ParseJSON parses the JSON representation of a `StatusRequest` messages from the `jsonBytes` argument.
func (msg *StatusRequest) ParseJSON(jsonBytes []byte) error {
	return json.Unmarshal(jsonBytes, msg)
}

// NewStatusRequestMessage returns with a new `StatusRequest` message. The header will contain the current time in `Nanoseconds` precision.
func NewStatusRequestMessage() msgs.Message {
	return NewStatusRequestMessageAt(time.Now().UnixNano(), "ns")
}

// NewStatusRequestMessageAt returns with a new `StatusRequest` message. The header will contain the `at` time in `withPrecision` precision.
func NewStatusRequestMessageAt(at int64, withPrecision common.TimePrecision) msgs.Message {
	var msg StatusRequest
	msg.Header = common.NewHeaderAt(at, withPrecision)
	return &msg
}
