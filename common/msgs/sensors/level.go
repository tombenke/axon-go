package sensors

import (
	"encoding/json"
	"github.com/tombenke/axon-go/common/msgs"
	"github.com/tombenke/axon-go/common/msgs/common"
)

const (
	LevelTypeName = "sensors/Level"
)

// Level message structure represent a physical level value, such as water level.
type Level struct {
	Header common.Header
	Body   common.Float64Body
}

func (this *Level) GetType() string {
	return LevelTypeName
}

func (this *Level) JSON() []byte {
	jsonBytes, err := json.Marshal(*this)
	if err != nil {
		panic(err)
	}
	return jsonBytes
}

func (this *Level) String() string {
	jsonBytes, err := json.Marshal(*this)
	if err != nil {
		panic(err)
	}
	return string(jsonBytes)
}

func (this *Level) ParseJSON(jsonBytes []byte) error {
	return json.Unmarshal(jsonBytes, this)
}

func NewLevelMessage(data float64) msgs.Message {
	var this Level
	this.Body.Data = data
	return &this
}
