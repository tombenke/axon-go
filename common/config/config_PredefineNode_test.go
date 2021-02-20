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
	extDescriptionValue := "This is an extensional property..."

	expectedAppConfigPredefined := AppConfigPredefined{
		Node:           GetDefaultNode(),
		ExtDescription: extDescriptionValue,
	}
	expectedAppConfigPredefined.Node.Name = nodeName
	expectedAppConfigPredefined.Node.Type = nodeType

	appConfigPredefined := AppConfigPredefined{
		Node:           NewNode(nodeName, nodeType, true, true, true, true),
		ExtDescription: extDescriptionValue,
	}
	fmt.Printf("appConfigPredefined:\n%v\n", appConfigPredefined)

	assert.Equal(t, expectedAppConfigPredefined, appConfigPredefined)
}
