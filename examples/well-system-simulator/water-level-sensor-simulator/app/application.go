package app

import (
	"github.com/tombenke/axon-go/common/actor/node"
	"github.com/tombenke/axon-go/common/log"
)

const (
	actorName = "water-level-sensor-simulator"
)

// Application struct represents the actor-node application
type Application struct {
	Node node.Node
}

// NewApplication creates a new actor-node application object
func NewApplication() Application {
	return Application{}
}

// Start initializes and starts new actor-node application according to its configuration
func (Application) Start(args []string) {
	// Loads node configuration
	config := GetConfig(actorName, args)

	// Configure the global logger of the application according to the configuration
	log.SetLevelStr(config.Node.LogLevel)
	log.SetFormatterStr(config.Node.LogFormat)
	config.Node.Logger = log.Logger
	config.Node.ClientID = "axon-go"
	config.Node.ClientName = "axon-go"
	config.Node.ClusterID = "test-cluster"

	// Start the additional components, if there is any
	// TODO

	// create and dtart the node
	actorNode := node.NewNode(config.Node)
	actorNode.Start()
}
