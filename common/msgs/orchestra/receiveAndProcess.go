package orchestra

import (
	"encoding/json"
	"fmt"
	"github.com/tombenke/axon-go/common/msgs"
	"github.com/tombenke/axon-go/common/msgs/common"
	"time"
)

const (
	// ReceiveAndProcessTypeName is the printable name of the `ReceiveAndProcess` message-type
	ReceiveAndProcessTypeName = "orchestra/ReceiveAndProcess"
)

func init() {
	msgs.RegisterMessageType(ReceiveAndProcessTypeName, []msgs.Representation{msgs.JSONRepresentation}, func() msgs.Message {
		return NewReceiveAndProcessMessage(float64(0))
	})
}

// ReceiveAndProcess represents the structure of the messages emitted by the orchestrator,
// and consumed by every actor to get synchronized with the reading of inputs and start processing.
// The `Body.Data` property holds the `dt` value since the last processing operation.
type ReceiveAndProcess struct {
	Header common.Header
	Body   common.Float64Body
}

// GetType returns with the printable name of the `ReceiveAndProcess` message-type
func (msg *ReceiveAndProcess) GetType() string {
	return ReceiveAndProcessTypeName
}

// Encode returns with the `ReceiveAndProcess` message content in a representation format selected by `representation`
func (msg *ReceiveAndProcess) Encode(representation msgs.Representation) (results []byte) {
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
func (msg *ReceiveAndProcess) Decode(representation msgs.Representation, content []byte) error {
	switch representation {
	case msgs.JSONRepresentation:
		return json.Unmarshal(content, msg)
	default:
		panic(fmt.Errorf("Decode error: unknown representational format '%s'", representation))
	}
}

// JSON returns with the `ReceiveAndProcess` message content in JSON representation format
func (msg *ReceiveAndProcess) JSON() []byte {
	jsonBytes, err := json.Marshal(*msg)
	if err != nil {
		panic(err)
	}
	return jsonBytes
}

// String returns with the `ReceiveAndProcess` message content in JSON format string
func (msg *ReceiveAndProcess) String() string {
	jsonBytes, err := json.Marshal(*msg)
	if err != nil {
		panic(err)
	}
	return string(jsonBytes)
}

// ParseJSON parses the JSON representation of a `ReceiveAndProcess` messages from the `jsonBytes` argument.
func (msg *ReceiveAndProcess) ParseJSON(jsonBytes []byte) error {
	return json.Unmarshal(jsonBytes, msg)
}

// NewReceiveAndProcessMessage returns with a new `ReceiveAndProcess` message. The header will contain the current time in `Nanoseconds` precision.
func NewReceiveAndProcessMessage(data float64) msgs.Message {
	return NewReceiveAndProcessMessageAt(data, time.Now().UnixNano(), "ns")
}

// NewReceiveAndProcessMessageAt returns with a new `ReceiveAndProcess` message. The header will contain the `at` time in `withPrecision` precision.
func NewReceiveAndProcessMessageAt(data float64, at int64, withPrecision common.TimePrecision) msgs.Message {
	var msg ReceiveAndProcess
	msg.Header = common.NewHeaderAt(at, withPrecision)
	msg.Body.Data = data
	return &msg
}
