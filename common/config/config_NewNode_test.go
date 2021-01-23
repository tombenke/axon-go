package config

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"testing"
)

type AppConfigFromNew struct {
	Node           Node   `yaml:"node"`
	ExtDescription string `yaml:"extDescription"`
}

func (c AppConfigFromNew) YAML() ([]byte, error) {
	return yaml.Marshal(&c)
}

func TestNewNode(t *testing.T) {

	nodeName := "app-node"
	nodeType := "app-node-type"
	extend := true
	modify := false
	extDescriptionValue := "This is an extensional property..."

	expectedAppConfigFromNew := AppConfigFromNew{
		Node: Node{
			Name: nodeName,
			Type: nodeType,
		},
		ExtDescription: extDescriptionValue,
	}
	appConfigFromNew := AppConfigFromNew{
		Node:           NewNode(nodeName, nodeType, extend, modify),
		ExtDescription: extDescriptionValue,
	}
	fmt.Printf("appConfigFromNew:\n%v\n", appConfigFromNew)

	assert.Equal(t, expectedAppConfigFromNew, appConfigFromNew)
}
