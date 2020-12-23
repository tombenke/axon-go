package sensors

import (
	"encoding/json"
	"fmt"
	"github.com/tombenke/axon-go/common/msgs"
	"github.com/tombenke/axon-go/common/msgs/common"
	"time"
)

const (
	// HumidityTypeName is the printable name of the `Humidity` message-type
	HumidityTypeName = "sensors/Humidity"
)

// Humidity message structure represent a physical level value, such as water level.
type Humidity struct {
	Header common.Header
	Body   common.Float64Body
}

// GetType returns with the printable name of the `Humidity` message-type
func (msg *Humidity) GetType() string {
	return HumidityTypeName
}

// Encode returns with the `Bool` message content in a representation format selected by `representation`
func (msg *Humidity) Encode(representation msgs.Representation) (results []byte) {
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
func (msg *Humidity) Decode(representation msgs.Representation, content []byte) error {
	switch representation {
	case msgs.JSONRepresentation:
		return json.Unmarshal(content, msg)
	default:
		panic(fmt.Errorf("Decode error: unknown representational format '%s'", representation))
	}
}

// JSON returns with the `Humidity` message content in JSON representation format
func (msg *Humidity) JSON() []byte {
	jsonBytes, err := json.Marshal(*msg)
	if err != nil {
		panic(err)
	}
	return jsonBytes
}

// String returns with the `Humidity` message content in JSON format string
func (msg *Humidity) String() string {
	jsonBytes, err := json.Marshal(*msg)
	if err != nil {
		panic(err)
	}
	return string(jsonBytes)
}

// ParseJSON parses the JSON representation of a `Humidity` messages from the `jsonBytes` argument.
func (msg *Humidity) ParseJSON(jsonBytes []byte) error {
	return json.Unmarshal(jsonBytes, msg)
}

// NewHumidityMessage returns with a new `Humidity` message. The header will contain the current time in `Nanoseconds` precision.
func NewHumidityMessage(data float64) msgs.Message {
	return NewHumidityMessageAt(data, time.Now().UnixNano(), "ns")
}

// NewHumidityMessageAt returns with a new `Humidity` message. The header will contain the `at` time in `withPrecision` precision.
func NewHumidityMessageAt(data float64, at int64, withPrecision common.TimePrecision) msgs.Message {
	var msg Humidity
	msg.Header = common.NewHeaderAt(at, withPrecision)
	msg.Body.Data = data
	return &msg
}
