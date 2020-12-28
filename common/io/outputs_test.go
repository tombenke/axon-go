package io

import (
	"github.com/stretchr/testify/assert"
	"github.com/tombenke/axon-go/common/config"
	"github.com/tombenke/axon-go/common/msgs/base"
	"testing"
)

func TestGetOutputMessage(t *testing.T) {
	bmsg := base.NewBoolMessage(true)
	out := Outputs{"State": Output{IO: IO{Name: "State", Type: base.BoolTypeName, Message: bmsg}}}
	rmsg := (out).GetOutputMessage("State")
	assert.Equal(t, rmsg, bmsg)
}

func TestGetOutputMessageWrongPort(t *testing.T) {
	bmsg := base.NewBoolMessage(true)
	out := Outputs{"State": Output{IO: IO{Name: "State", Type: base.BoolTypeName, Message: bmsg}}}
	assert.Panics(t, func() { out.GetOutputMessage("WrongPort") })
}

func TestSetOutputMessage(t *testing.T) {
	bmsg := base.NewBoolMessage(true)
	out := Outputs{"State": Output{IO{Name: "State", Type: base.BoolTypeName, Message: bmsg}}}
	(out).SetOutputMessage("State", bmsg)
	assert.Equal(t, out["State"].IO.Message.String(), bmsg.String())
}

func TestSetOutputMessageWrongPort(t *testing.T) {
	bmsg := base.NewBoolMessage(true)
	out := Outputs{"State": Output{IO{Name: "State", Type: base.BoolTypeName, Message: bmsg}}}
	assert.Panics(t, func() { out.SetOutputMessage("WrongPortName", bmsg) })
}

func TestSetOutputMessageWrongMessageType(t *testing.T) {
	bmsg := base.NewBoolMessage(true)
	out := Outputs{"State": Output{IO{Name: "State", Type: base.BoolTypeName, Message: bmsg}}}
	smsg := base.NewStringMessage("Wrong message")
	assert.Panics(t, func() { out.SetOutputMessage("State", smsg) })
}

func TestNewOutputs(t *testing.T) {
	outputsCfg := config.Outputs{
		config.Out{IO: config.IO{Name: "sensor-value", Type: "base/Bool", Representation: "application/json", Channel: "value-of-sensor-1"}},
		config.Out{IO: config.IO{Name: "node-state", Type: "base/String", Representation: "application/json", Channel: "state-of-the-node"}},
	}
	outputs := NewOutputs(outputsCfg)
	assert.Equal(t, len(outputs), 2)
}

func TestNewOutputsWithUnregisteredMessageType(t *testing.T) {
	outputsCfg := config.Outputs{
		config.Out{IO: config.IO{Name: "sensor-value", Type: "base/WrongType", Representation: "application/json", Channel: "value-of-sensor-1"}},
	}
	assert.Panics(t, func() { NewOutputs(outputsCfg) })
}

func TestNewOutputsWithMissingRepresentation(t *testing.T) {
	outputsCfg := config.Outputs{
		config.Out{IO: config.IO{Name: "sensor-value", Type: "base/Bool", Representation: "wrong/representation", Channel: "value-of-sensor-1"}},
	}
	assert.Panics(t, func() { NewOutputs(outputsCfg) })
}
