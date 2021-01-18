package config

import (
	"github.com/tombenke/axon-go/common/messenger"
)

// Node is the main aggregate that holds the default config struct that every axon actor node inherits
type Node struct {
	// Messenger holds the configuration parameters of the messaging middleware
	Messenger messenger.Config `yaml:"messenger"`

	// Name is the name of the node. It should be unique in a specific network
	Name string `yaml:"name"`

	// Type is the symbolic name of the node type, that refers to how the node is working.
	Type string `yaml:"type"`

	// ConfigFileName is the name of the config file to load
	// the configuration parameters of the application.
	// Its default value is `config.yml`.
	// It is optional to use config files. When the node starts, it tries to find the config file
	// identified by this property, and loads it it it has been found.
	ConfigFileName string `yaml:"configFileName"`

	// LogLevel is the log level of the application
	LogLevel string `yaml:"logLevel"`

	// LogFormat the log format of the application
	LogFormat string `yaml:"logFormat"`

	// Configure holds the properties that determine how the node configuration properties
	// might be changed, during the configuration process
	Configure Configure `yaml:"configure"`

	// Ports holds the I/O port definitions
	Ports Ports `yaml:"ports"`
}

// Ports structure is an aggregate that holds the I/O port definitions
type Ports struct {
	// Inputs is a list of input-type port descriptors

	Inputs Inputs `yaml:"inputs"`
	// Outputs is a list of output-type port descriptors

	Outputs Outputs `yaml:"outputs"`
	// Configure holds the properties that determine how the I/O port configuration properties
	// might be changed, during the configuration process
	Configure Configure `yaml:"configure"`
}

// Configure is a generic structure that holds the flags that control either the node level,
// and/or the port-level configurability
type Configure struct {
	// Extends is a flag, that determines if the node and/or port can be extended.
	// In case of Node `true` means, it is allowed to add extensional properties
	// to the configuration of the application.
	// If `false`, there is only the predefined application config can be used.
	// In case of ports `true` means, it is possible to add new I/O ports to the node.
	// if `false` there is only the predefined ports can be used.

	Extend bool `yaml:"extend"`
	// Modify is a flag that determines if the values of the configuration properties of
	// the node and/or port can be changed or not. If `true` the properties can be changed,
	// othewise only the predefined values can be used.
	Modify bool `yaml:"modify"`
}

// Orchestration structure holds those configuration parameters of a Node that determine
// how the Node behaves in the network from organizational point of view, e.g.
// if it uses synchronization and presence or not.
// This structure also holds the names of the channels the presence and synchronization processes use.
type Orchestration struct {
	// Presence is a flag. If it is `true` the Node uses presence protocol, otherwise not.
	Presence bool `yaml:"presence"`

	// Synchronize is a flag. If it is `true` the Node is working in syncronized mode,
	// otherwise it uses no synchronozation protocol.
	Synchronization bool `yaml:"synchronization"`

	// Channel holds the names of the channels used by the presence and the synchronization protocols
	Channels Channels `yaml:"channels"`
}

// Channels is s structure, that holds the names of the channels used by the presence
// and the synchronization protocols
type Channels struct {
	// StatusRequest is the name of the channel that the orchestrator uses
	// to send status request message to the nodes of the network.
	// The Nodes that uses the presence protocol must subscribe to this channel.
	StatusRequest string `yaml:"StatusRequest"`

	// StatusReport is the name of the channel that the orchestrator uses
	// to receive status response messages from the nodes of the network.
	// The Nodes that uses the presence protocol must publish their status response messages to this
	// this channel after they received a status request from the orchestrator.
	StatusReport string `yaml:"StatusReport"`

	// SendResults is the name of the channel that the orchestrator uses
	// to notify the nodes of the network to send their processing results.
	// The Nodes that work in synchronous mode must subscribe to this channel,
	// and they have to publish their results after receiving this message.
	SendResults string `yaml:"SendResults"`

	// SendingCompleted is the name of the channel that the orchestrator subscribes to
	// in order to get notified by those Nodes that completed the sending of their processing results.
	// The Nodes that work in synchronous mode must publish to this channel
	// the sending-completed message, which includes the ID of the Node.
	SendingCompleted string `yaml:"SendingCompleted"`

	// ReceiveAndProcess is the name of the channel that the orchestrator uses
	// to notify the nodes of the network to collect the messages they have received via the intput ports,
	// then send the this collection to the processing function for to with with.
	// The Nodes that work in synchronous mode must subscribe to this channel.
	ReceiveAndProcess string `yaml:"receiveAndProcess"`

	// ProcessingCompleted is the name of the channel that the orchestrator subscribes to
	// in order to get notified by those Nodes that completed the processing of the incoming messages.
	// The Nodes that work in synchronous mode must publish to this channel
	// the processing-completed message which includes the ID of the Node.
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
