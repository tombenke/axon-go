package config

import (
	//"fmt"
	"github.com/tombenke/axon-go/common/messenger"
)

// Node is the main aggregate that holds the default config struct that every axon actor node inherits
type Node struct {
	// Messenger holds the configuration parameters of the messaging middleware
	Messenger messenger.Config `yaml:"messenger"`

	// Name is the name of the node. It should be unique in a specific network
	Name string `yaml:"name"`

	// Type is the symbolic name of the node type, that refers to how the node is working.
	Type string `yaml:-`

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

	// Ports holds the I/O port definitions
	Ports Ports `yaml:"ports"`

	Orchestration Orchestration `yaml:"orchestration"`
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
	// Extends is a flag, that determines if the set of ports can be extended.
	// If it is `true`, then it is possible to add new I/O ports to the node.
	// if `false` there is only the predefined ports can be used.
	Extend bool `yaml:"extend"`

	// Modify is a flag that determines if the values of the configuration properties of
	// the ports can be changed or not. If `true` the properties can be changed,
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
	StatusRequest string `yaml:"statusRequest"`

	// StatusReport is the name of the channel that the orchestrator uses
	// to receive status response messages from the nodes of the network.
	// The Nodes that uses the presence protocol must publish their status response messages to this
	// this channel after they received a status request from the orchestrator.
	StatusReport string `yaml:"statusReport"`

	// SendResults is the name of the channel that the orchestrator uses
	// to notify the nodes of the network to send their processing results.
	// The Nodes that work in synchronous mode must subscribe to this channel,
	// and they have to publish their results after receiving this message.
	SendResults string `yaml:"sendResults"`

	// SendingCompleted is the name of the channel that the orchestrator subscribes to
	// in order to get notified by those Nodes that completed the sending of their processing results.
	// The Nodes that work in synchronous mode must publish to this channel
	// the sending-completed message, which includes the ID of the Node.
	SendingCompleted string `yaml:"sendingCompleted"`

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

// GetDefaultNode returns with a new Node structure with default values
func GetDefaultNode() Node {
	return Node{
		Messenger: messenger.Config{
			Urls:      defaultMessagingURL,
			UserCreds: defaultMessagingUserCreds,
		},
		Name:           "anonymous",
		Type:           "untyped",
		ConfigFileName: "config.yml",
		LogLevel:       defaultLogLevel,
		LogFormat:      defaultLogFormat,
		Ports: Ports{
			Configure: Configure{
				Extend: true,
				Modify: true,
			},
		},
		Orchestration: Orchestration{
			Presence:        true,
			Synchronization: true,
			Channels: Channels{
				StatusRequest:       "status-request",
				StatusReport:        "status-report",
				SendResults:         "send-results",
				SendingCompleted:    "sending-completed",
				ReceiveAndProcess:   "receive-and-process",
				ProcessingCompleted: "processing-completed",
			},
		},
	}
}

// NewNode returns with a new Node configuration object
func NewNode(nodeName string, nodeType string, extend bool, modify bool) Node {
	return Node{
		Name:  nodeName,
		Type:  nodeType,
		Ports: Ports{},
	}
}

// SetPortsConfigurability sets if the I/O ports can be extended and/or modified
func (n *Node) SetPortsConfigurability(extend bool, modify bool) {
	n.Ports.Configure.Extend = extend
	n.Ports.Configure.Modify = modify
}

// AddInputPort Add a new input port to the Node
func (n *Node) AddInputPort(portName string, portType string, representation string, channel string, defaultMsg string) {
	input := In{IO: IO{Name: portName, Channel: channel, Type: portType, Representation: representation}, Default: defaultMsg}
	n.Ports.Inputs = append(n.Ports.Inputs, input)
}

// AddOutputPort Add a new output port to the Node
func (n *Node) AddOutputPort(portName string, portType string, representation string, channel string) {
	output := Out{IO: IO{Name: portName, Channel: channel, Type: portType, Representation: representation}}
	n.Ports.Outputs = append(n.Ports.Outputs, output)
}

// MergeNodeConfigs returns with the resulting config parameters set of the Node
// after merging the coming from the three sources
func MergeNodeConfigs(hardCoded Node, cli Node, configFile Node) Node {
	resulting := hardCoded

	//TODO: Implement
	resulting.Name = cli.Name
	resulting.LogLevel = cli.LogLevel
	resulting.LogFormat = cli.LogFormat
	resulting.Messenger = cli.Messenger
	//resulting.Orchestration = cli.Orchestration

	//fmt.Println("MergeConfig:", hardCoded, cli, configFile, "=>", resulting)
	return resulting
}
