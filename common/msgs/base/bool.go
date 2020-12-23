package base

import (
	"encoding/json"
	"github.com/tombenke/axon-go/common/msgs"
)

const (
	BoolTypeName = "base/Bool"
)

type Bool struct {
	Data bool
}

func (this *Bool) GetType() string {
	return BoolTypeName
}

func (this *Bool) JSON() []byte {
	jsonBytes, err := json.Marshal(*this)
	if err != nil {
		panic(err)
	}
	return jsonBytes
}

func (this *Bool) String() string {
	jsonBytes, err := json.Marshal(*this)
	if err != nil {
		panic(err)
	}
	return string(jsonBytes)
}

func (this *Bool) ParseJSON(jsonBytes []byte) error {
	return json.Unmarshal(jsonBytes, this)
}

func NewBoolMessage(data bool) msgs.Message {
	var this Bool
	this.Data = data
	return &this
}
