package orchestra

import (
	"encoding/json"
	"fmt"
	"github.com/tombenke/axon-go/common/msgs"
	"github.com/tombenke/axon-go/common/msgs/common"
	"time"
)

const (
	// StatusReportTypeName is the printable name of the `StatusReport` message-type
	StatusReportTypeName = "orchestra/StatusReport"
)

func init() {
	msgs.RegisterMessageType(StatusReportTypeName, []msgs.Representation{msgs.JSONRepresentation}, func() msgs.Message {
		return NewStatusReportMessage("")
	})
}

// StatusReport represents the structure of the `status-report` message that the actor is usually sent
// to the orchestrator as a response to the `status-request` message.
// The `Body.Data` property holds the name of the actor that sends the response.
type StatusReport struct {
	Header common.Header
	Body   common.StringBody
}

// GetType returns with the printable name of the `StatusReport` message-type
func (msg *StatusReport) GetType() string {
	return StatusReportTypeName
}

// Encode returns with the `StatusReport` message content in a representation format selected by `representation`
func (msg *StatusReport) Encode(representation msgs.Representation) (results []byte) {
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
func (msg *StatusReport) Decode(representation msgs.Representation, content []byte) error {
	switch representation {
	case msgs.JSONRepresentation:
		return json.Unmarshal(content, msg)
	default:
		panic(fmt.Errorf("Decode error: unknown representational format '%s'", representation))
	}
}

// JSON returns with the `StatusReport` message content in JSON representation format
func (msg *StatusReport) JSON() []byte {
	jsonBytes, err := json.Marshal(*msg)
	if err != nil {
		panic(err)
	}
	return jsonBytes
}

// StatusReport returns with the `StatusReport` message content in JSON format string
func (msg *StatusReport) String() string {
	jsonBytes, err := json.Marshal(*msg)
	if err != nil {
		panic(err)
	}
	return string(jsonBytes)
}

// ParseJSON parses the JSON representation of a `StatusReport` messages from the `jsonBytes` argument.
func (msg *StatusReport) ParseJSON(jsonBytes []byte) error {
	return json.Unmarshal(jsonBytes, msg)
}

// NewStatusReportMessage returns with a new `StatusReport` message. The header will contain the current time in `Nanoseconds` precision.
func NewStatusReportMessage(data string) msgs.Message {
	return NewStatusReportMessageAt(data, time.Now().UnixNano(), "ns")
}

// NewStatusReportMessageAt returns with a new `StatusReport` message. The header will contain the `at` time in `withPrecision` precision.
func NewStatusReportMessageAt(data string, at int64, withPrecision common.TimePrecision) msgs.Message {
	var msg StatusReport
	msg.Header = common.NewHeaderAt(at, withPrecision)
	msg.Body.Data = data
	return &msg
}
