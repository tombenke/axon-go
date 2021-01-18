package config

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"testing"
)

type AppConfigPredefined struct {
	Node           Node   `yaml:"node"`
	ExtDescription string `yaml:"extDescription"`
}

func (c AppConfigPredefined) YAML() ([]byte, error) {
	return yaml.Marshal(&c)
}

func TestPredefineNode(t *testing.T) {

	nodeName := "app-node"
	nodeType := "app-node-type"
	extend := true
	modify := false
	extDescriptionValue := "This is an extensional property..."

	expectedAppConfigPredefined := AppConfigPredefined{
		Node: Node{
			Name: nodeName,
			Type: nodeType,
			Configure: Configure{
				Extend: extend,
				Modify: modify,
			},
		},
		ExtDescription: extDescriptionValue,
	}
	appConfigPredefined := AppConfigPredefined{
		Node:           NewNode(nodeName, nodeType, extend, modify),
		ExtDescription: extDescriptionValue,
	}
	fmt.Printf("appConfigPredefined:\n%v\n", appConfigPredefined)

	assert.Equal(t, expectedAppConfigPredefined, appConfigPredefined)
}
