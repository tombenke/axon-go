package sensors

import (
	"encoding/json"
	"fmt"
	"github.com/tombenke/axon-go/common/msgs"
	"github.com/tombenke/axon-go/common/msgs/common"
	"time"
)

const (
	// TemperatureTypeName is the printable name of the `Temperature` message-type
	TemperatureTypeName = "sensors/Temperature"
)

func init() {
	msgs.RegisterMessageType(TemperatureTypeName, []msgs.Representation{msgs.JSONRepresentation}, func() msgs.Message {
		return NewTemperatureMessage(float64(0))
	})
}

// Temperature represents the structure of the messages emitted by the Temperature sensors
type Temperature struct {
	Header common.Header
	Body   common.Float64VarBody
}

// GetType returns with the printable name of the `Temperature` message-type
func (msg *Temperature) GetType() string {
	return TemperatureTypeName
}

// Encode returns with the `Bool` message content in a representation format selected by `representation`
func (msg *Temperature) Encode(representation msgs.Representation) (results []byte) {
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
func (msg *Temperature) Decode(representation msgs.Representation, content []byte) error {
	switch representation {
	case msgs.JSONRepresentation:
		return json.Unmarshal(content, msg)
	default:
		panic(fmt.Errorf("Decode error: unknown representational format '%s'", representation))
	}
}

// JSON returns with the `Temperature` message content in JSON representation format
func (msg *Temperature) JSON() []byte {
	jsonBytes, err := json.Marshal(*msg)
	if err != nil {
		panic(err)
	}
	return jsonBytes
}

// String returns with the `Temperature` message content in JSON format string
func (msg *Temperature) String() string {
	jsonBytes, err := json.Marshal(*msg)
	if err != nil {
		panic(err)
	}
	return string(jsonBytes)
}

// ParseJSON parses the JSON representation of a `Temperature` messages from the `jsonBytes` argument.
func (msg *Temperature) ParseJSON(jsonBytes []byte) error {
	return json.Unmarshal(jsonBytes, msg)
}

// NewTemperatureMessage returns with a new `Temperature` message. The header will contain the current time in `Nanoseconds` precision.
func NewTemperatureMessage(data float64) msgs.Message {
	return NewTemperatureMessageAt(data, time.Now().UnixNano(), "ns")
}

// NewTemperatureMessageAt returns with a new `Temperature` message. The header will contain the `at` time in `withPrecision` precision.
func NewTemperatureMessageAt(data float64, at int64, withPrecision common.TimePrecision) msgs.Message {
	var msg Temperature
	msg.Header = common.NewHeaderAt(at, withPrecision)
	msg.Body.Data = data
	return &msg
}
