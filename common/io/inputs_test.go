package io

import (
	"github.com/stretchr/testify/assert"
	"github.com/tombenke/axon-go/common/config"
	"github.com/tombenke/axon-go/common/msgs/base"
	"testing"
)

func TestInputsGetMessage(t *testing.T) {
	bmsg := base.NewBoolMessage(true)
	in := Inputs{"State": Input{IO: IO{Name: "State", Type: base.BoolTypeName, Message: bmsg}, DefaultMessage: bmsg}}
	rmsg := (in).GetMessage("State")
	assert.Equal(t, rmsg, bmsg)
}

func TestInputsGetMessageWrongPort(t *testing.T) {
	bmsg := base.NewBoolMessage(true)
	in := Inputs{"State": Input{IO: IO{Name: "State", Type: base.BoolTypeName, Message: bmsg}, DefaultMessage: bmsg}}
	assert.Panics(t, func() { (in).GetMessage("WrongPort") })
}

func TestInputsSetMessage(t *testing.T) {
	bmsg := base.NewBoolMessage(true)
	in := Inputs{"State": Input{IO: IO{Name: "State", Type: base.BoolTypeName, Message: bmsg}, DefaultMessage: bmsg}}
	(in).SetMessage("State", bmsg)
	assert.Equal(t, in["State"].IO.Message.String(), bmsg.String())
}

func TestInputsSetMessageWrongPort(t *testing.T) {
	bmsg := base.NewBoolMessage(true)
	in := Inputs{"State": Input{IO: IO{Name: "State", Type: base.BoolTypeName, Message: bmsg}, DefaultMessage: bmsg}}
	assert.Panics(t, func() { (in).SetMessage("WrongPortName", bmsg) })
}

func TestInputsSetMessageWrongMessageType(t *testing.T) {
	bmsg := base.NewBoolMessage(true)
	in := Inputs{"State": Input{IO: IO{Name: "State", Type: base.BoolTypeName, Message: bmsg}, DefaultMessage: bmsg}}
	smsg := base.NewStringMessage("Wrong message")
	assert.Panics(t, func() { (in).SetMessage("State", smsg) })
}

func TestNewInputsNoDefault(t *testing.T) {
	inputsCfg := config.Inputs{
		config.In{IO: config.IO{Name: "sensor-value", Type: "base/Bool", Representation: "application/json", Channel: "value-of-sensor-1"}, Default: ""},
		config.In{IO: config.IO{Name: "node-state", Type: "base/String", Representation: "application/json", Channel: "state-of-the-node"}, Default: ""},
	}
	inputs := NewInputs(inputsCfg)
	assert.Equal(t, len(inputs), 2)
}

func TestNewInputsJSONDefault(t *testing.T) {
	inputsCfg := config.Inputs{
		config.In{IO: config.IO{Name: "sensor-value", Type: "base/Bool", Representation: "application/json", Channel: "value-of-sensor-1"}, Default: `{"Body": {"Data": true}}`},
		config.In{IO: config.IO{Name: "node-state", Type: "base/String", Representation: "application/json", Channel: "state-of-the-node"}, Default: `{"Body": {"Data": "Some text..."}}`},
	}
	inputs := NewInputs(inputsCfg)
	assert.Equal(t, len(inputs), 2)
}

func TestNewInputsWrongJSONDefault(t *testing.T) {
	inputsCfg := config.Inputs{
		config.In{IO: config.IO{Name: "node-state", Type: "base/String", Representation: "application/json", Channel: "state-of-the-node"}, Default: `{"Body": some wrong formatted default value { \}`},
	}
	assert.Panics(t, func() { NewInputs(inputsCfg) })
}

func TestNewInputsWithUnregisteredMessageType(t *testing.T) {
	inputsCfg := config.Inputs{
		config.In{IO: config.IO{Name: "sensor-value", Type: "base/WrongType", Representation: "application/json", Channel: "value-of-sensor-1"}, Default: ""},
	}
	assert.Panics(t, func() { NewInputs(inputsCfg) })
}

func TestNewInputsWithMissingRepresentation(t *testing.T) {
	inputsCfg := config.Inputs{
		config.In{IO: config.IO{Name: "sensor-value", Type: "base/Bool", Representation: "wrong/representation", Channel: "value-of-sensor-1"}, Default: ""},
	}
	assert.Panics(t, func() { NewInputs(inputsCfg) })
}
