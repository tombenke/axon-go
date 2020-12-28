package io

import (
	"github.com/stretchr/testify/assert"
	"github.com/tombenke/axon-go/common/config"
	"github.com/tombenke/axon-go/common/msgs/base"
	"testing"
)

func TestOutputsSetOutputMessage(t *testing.T) {
	bmsg := base.NewBoolMessage(true)
	o := Outputs{"State": Output{IO{Name: "State", Type: base.BoolTypeName, Message: bmsg}}}
	(o).SetOutputMessage("State", bmsg)
	//TODO: Write the test
}

func TestNewOutputs(t *testing.T) {
	outputsCfg := config.Outputs{
		config.Out{IO: config.IO{Name: "sensor-value", Type: "base/Bool", Representation: "application/json", Channel: "value-of-sensor-1"}},
		config.Out{IO: config.IO{Name: "node-state", Type: "base/String", Representation: "application/json", Channel: "state-of-the-node"}},
	}
	outputs := NewOutputs(outputsCfg)
	assert.Equal(t, len(outputs), 2)
}
