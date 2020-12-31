package io

import (
	"github.com/stretchr/testify/assert"
	"github.com/tombenke/axon-go/common/config"
	"github.com/tombenke/axon-go/common/msgs"
	"github.com/tombenke/axon-go/common/msgs/base"
	"testing"
)

func TestOutputsGetMessage(t *testing.T) {
	bmsg := base.NewBoolMessage(true)
	out := Outputs{"State": Output{IO: IO{Name: "State", Type: base.BoolTypeName, Message: bmsg}}}
	rmsg := (out).GetMessage("State")
	assert.Equal(t, rmsg, bmsg)
}

func TestOutputsGetMessageWrongPort(t *testing.T) {
	bmsg := base.NewBoolMessage(true)
	out := Outputs{"State": Output{IO: IO{Name: "State", Type: base.BoolTypeName, Message: bmsg}}}
	assert.Panics(t, func() { out.GetMessage("WrongPort") })
}

func TestOutputsSetMessage(t *testing.T) {
	bmsg := base.NewBoolMessage(true)
	out := Outputs{"State": Output{IO{Name: "State", Type: base.BoolTypeName, Message: bmsg}}}
	(out).SetMessage("State", bmsg)
	assert.Equal(t, out["State"].IO.Message.String(), bmsg.String())
}

func TestOutputsSetMessageWrongPort(t *testing.T) {
	bmsg := base.NewBoolMessage(true)
	out := Outputs{"State": Output{IO{Name: "State", Type: base.BoolTypeName, Message: bmsg}}}
	assert.Panics(t, func() { out.SetMessage("WrongPortName", bmsg) })
}

func TestOutputsSetMessageWrongMessageType(t *testing.T) {
	bmsg := base.NewBoolMessage(true)
	out := Outputs{"State": Output{IO{Name: "State", Type: base.BoolTypeName, Message: bmsg}}}
	smsg := base.NewStringMessage("Wrong message")
	assert.Panics(t, func() { out.SetMessage("State", smsg) })
}

func TestNewOutputs(t *testing.T) {
	outputsCfg := config.Outputs{
		config.Out{IO: config.IO{Name: "sensor-value", Type: "base/Bool", Representation: "application/json", Channel: "value-of-sensor-1"}},
		config.Out{IO: config.IO{Name: "node-state", Type: "base/String", Representation: "application/json", Channel: "state-of-the-node"}},
	}
	outputs := NewOutputs(outputsCfg)
	assert.Equal(t, len(outputs), 2)
	for _, oCfg := range outputsCfg {
		assert.Equal(t, oCfg.Name, outputs[oCfg.Name].Name)
		assert.Equal(t, oCfg.Type, outputs[oCfg.Name].Type)
		assert.Equal(t, msgs.Representation(oCfg.Representation), outputs[oCfg.Name].Representation)
		assert.Equal(t, oCfg.Channel, outputs[oCfg.Name].Channel)
	}
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
