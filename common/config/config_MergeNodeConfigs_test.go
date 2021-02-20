package config

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var hcInputs = Inputs{
	In{IO: IO{
		Name:           "reference-water-level",
		Type:           "base/Float64",
		Representation: "application/json",
		Channel:        "",
	}, Default: `{"Body": {"Data": 0.75}}`},
	In{IO: IO{
		Name:           "water-level",
		Type:           "base/Float64",
		Representation: "application/json",
		Channel:        "well-water-buffer-tank-level",
	}, Default: ""},
}

var hcOutputs = Outputs{
	Out{IO: IO{
		Name:           "water-level-state",
		Type:           "base/Bool",
		Representation: "application/json",
		Channel:        "buffer-water-tank-upper-level-state",
	}},
}

var cliModInputs = Inputs{
	In{IO: IO{
		Name:           "reference-water-level",
		Type:           "base/Float64",
		Representation: "text/xml",
		Channel:        "reference-water-level-ch",
	}, Default: ""},
	In{IO: IO{
		Name:           "water-level",
		Type:           "base/Float64",
		Representation: "application/json",
		Channel:        "well-water-buffer-tank-level",
	}, Default: ""},
}

var cliModOutputs = Outputs{
	Out{IO: IO{
		Name:           "water-level-state",
		Type:           "base/Bool",
		Representation: "application/protobuf",
		Channel:        "/dev/null",
	}},
}

var cliExtOutputs = Outputs{
	Out{IO: IO{
		Name:           "water-level-state",
		Type:           "base/Bool",
		Representation: "application/json",
		Channel:        "buffer-water-tank-upper-level-state",
	}},
	Out{IO: IO{
		Name:           "aditional-output",
		Type:           "base/Bool",
		Representation: "application/json",
		Channel:        "another-channel",
	}},
}

func makeNode(nodeName string, nodeType string, extend bool, modify bool, presence bool, sync bool, inputs Inputs, outputs Outputs) Node {
	node := NewNode(nodeName, nodeType, extend, modify, presence, sync)
	node.Ports.Inputs = inputs
	node.Ports.Outputs = outputs
	return node
}

func TestMergeNodeConfigs_noExt_noMod_noAdd(t *testing.T) {
	hardCoded := makeNode("test-node", "test-node-type", false, false, true, true, hcInputs, hcOutputs)
	cli := makeNode("test-node", "test-node-type", false, false, true, true, Inputs{}, Outputs{})
	resulting, err := MergeNodeConfigs(hardCoded, cli)
	assert.Nil(t, err)
	assert.Equal(t, hardCoded, resulting)
}

func TestMergeNodeConfigs_noExt_noMod_Add(t *testing.T) {
	hardCoded := makeNode("test-node", "test-node-type", false, false, true, true, Inputs{}, Outputs{})
	cli := makeNode("test-node", "test-node-type", false, false, true, true, hcInputs, hcOutputs)
	resulting, err := MergeNodeConfigs(hardCoded, cli)
	assert.NotNil(t, err)
	assert.Equal(t, "port extension is disabled", err.Error())
	assert.Equal(t, hardCoded, resulting)
}

func TestMergeNodeConfigs_Ext_Mod_noAdd(t *testing.T) {
	hardCoded := makeNode("test-node", "test-node-type", true, true, true, true, hcInputs, hcOutputs)
	cli := makeNode("test-node", "test-node-type", true, true, true, true, Inputs{}, Outputs{})
	resulting, err := MergeNodeConfigs(hardCoded, cli)
	assert.Nil(t, err)
	assert.Equal(t, hardCoded, resulting)
}

func TestMergeNodeConfigs_Ext_Mod_Add(t *testing.T) {
	hardCoded := makeNode("test-node", "test-node-type", true, true, true, true, Inputs{}, Outputs{})
	cli := makeNode("test-node", "test-node-type", true, true, true, true, hcInputs, hcOutputs)
	resulting, err := MergeNodeConfigs(hardCoded, cli)
	assert.Nil(t, err)
	assert.Equal(t, cli, resulting)
}

func TestMergeNodeConfigs_Ext_Mod_Add2(t *testing.T) {
	hardCoded := makeNode("test-node", "test-node-type", true, true, true, true, Inputs{}, Outputs{})
	cli := makeNode("test-node", "test-node-type", true, true, true, true, Inputs{}, cliExtOutputs)
	resulting, err := MergeNodeConfigs(hardCoded, cli)
	assert.Nil(t, err)
	assert.Equal(t, cli, resulting)
}

func TestMergeNodeConfigs_noExt_noMod_Mod(t *testing.T) {
	hardCoded := makeNode("test-node", "test-node-type", false, false, true, true, hcInputs, hcOutputs)
	cli := makeNode("test-node", "test-node-type", false, false, true, true, cliModInputs, cliModOutputs)

	resulting, err := MergeNodeConfigs(hardCoded, cli)
	assert.NotNil(t, err)
	assert.Equal(t, "port modification is disabled", err.Error())
	assert.Equal(t, hardCoded, resulting)
}

func TestMergeNodeConfigs_noExt_Mod_Mod(t *testing.T) {
	hardCoded := makeNode("test-node", "test-node-type", false, true, true, true, hcInputs, hcOutputs)
	cli := makeNode("test-node", "test-node-type", false, true, true, true, cliModInputs, hcOutputs)

	resulting, err := MergeNodeConfigs(hardCoded, cli)
	assert.Nil(t, err)
	assert.Equal(t, cli, resulting)
}

func TestMergeNodeConfigs_noExt_Mod_Mod2(t *testing.T) {
	hardCoded := makeNode("test-node", "test-node-type", false, true, true, true, hcInputs, hcOutputs)
	cli := makeNode("test-node", "test-node-type", false, true, true, true, hcInputs, cliModOutputs)

	resulting, err := MergeNodeConfigs(hardCoded, cli)
	assert.Nil(t, err)
	assert.Equal(t, cli, resulting)
	fmt.Println("resulting", resulting)
}
