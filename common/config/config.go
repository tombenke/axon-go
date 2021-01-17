package config

import (
	"github.com/tombenke/axon-go/common/messenger"
)

// Node is the default config struct that every axon actor node inherits
type Node struct {
	Messenger      messenger.Config `yaml:"messenger"`
	Name           string           `yaml:"name"`
	Type           string           `yaml:"type"`
	ConfigFileName string           `yaml:"configFileName"`
	LogLevel       string           `yaml:"logLevel"`
	LogFormat      string           `yaml:"logFormat"`
	Configure      Configure        `yaml:"configure"`
	Ports          Ports            `yaml:"ports"`
}

type Ports struct {
	Inputs    Inputs    `yaml:"inputs"`
	Outputs   Outputs   `yaml:"outputs"`
	Configure Configure `yaml:"configure"`
}

type Configure struct {
	Extend bool `yaml:"extend"`
	Modify bool `yaml:"modify"`
}

type Messenger struct {
	Urls        string `yaml:"urls"`
	Credentials string `yaml:"credentials"`
}

type Orchestration struct {
	Presence        bool     `yaml:"presence"`
	Synchronization bool     `yaml:"synchronization"`
	Channels        Channels `yaml:"channels"`
}

type Channels struct {
	StatusRequest       string `yaml:"StatusRequest"`
	StatusReport        string `yaml:"StatusReport"`
	SendResults         string `yaml:"SendResults"`
	SendingCompleted    string `yaml:"SendingCompleted"`
	ReceiveAndProcess   string `yaml:"receiveAndProcess"`
	ProcessingCompleted string `yaml:"processingCompleted"`
}

// NewNode returns with a new Node configuration object
func NewNode(nodeName string, nodeType string, extend bool, modify bool) Node {
	return Node{
		Name: nodeName,
		Type: nodeType,
		Configure: Configure{
			Extend: extend,
			Modify: modify,
		},
	}
}

// SetPorts sets the predefined ports of the node, including their configurability settings
func (n *Node) SetPorts(inputs Inputs, outputs Outputs, extend bool, modify bool) {
	n.Ports.Inputs = inputs
	n.Ports.Outputs = outputs
	n.Ports.Configure.Extend = extend
	n.Ports.Configure.Modify = modify
}
