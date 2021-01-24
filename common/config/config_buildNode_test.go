package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddInputPort(t *testing.T) {
	node := GetDefaultNode()
	assert.Equal(t, 0, len(node.Ports.Inputs))
	portName := "new-port"
	portType := "base/Float64"
	representation := "application/json"
	channel := "a-data-channel"
	defaultMsg := ""
	expectedInput := In{IO: IO{Name: portName, Type: portType, Representation: representation, Channel: channel}, Default: defaultMsg}
	node.AddInputPort(portName, portType, representation, channel, defaultMsg)
	assert.Equal(t, 1, len(node.Ports.Inputs))
	assert.Equal(t, expectedInput, node.Ports.Inputs[0])
}

func TestAddOutputPort(t *testing.T) {
	node := GetDefaultNode()
	assert.Equal(t, 0, len(node.Ports.Outputs))
	portName := "new-port"
	portType := "base/Float64"
	representation := "application/json"
	channel := "a-data-channel"
	expectedOutput := Out{IO: IO{Name: portName, Type: portType, Representation: representation, Channel: channel}}
	node.AddOutputPort(portName, portType, representation, channel)
	assert.Equal(t, 1, len(node.Ports.Outputs))
	assert.Equal(t, expectedOutput, node.Ports.Outputs[0])
}
