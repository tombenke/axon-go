package app

import (
	"github.com/tombenke/axon-go/common/actor/node"
	"github.com/tombenke/axon-go/common/config"
	"github.com/tombenke/axon-go/common/log"
)

const (
	actorName = "water-level-sensor-simulator"
)

// Application struct represents the actor-node application
type Application struct {
	Node   node.Node
	config Config
}

// NewApplication creates a new actor-node application object
func NewApplication(args []string) Application {
	app := Application{}

	// Merge hard-coded configuration with CLI and file config if there is any
	app.config = GetConfig(actorName, makeHardCodedConfig(), args)

	// Configure the global logger of the application according to the configuration
	log.SetLevelStr(app.config.Node.LogLevel)
	log.SetFormatterStr(app.config.Node.LogFormat)

	// Print config to the console in YAML format
	if app.config.PrintConfig {
		printResultingConfig(app.config)
	}

	// Create the Node
	app.Node = node.NewNode(app.config.Node)

	return app
}

// makeHardCodedConfig returns with the built-in configuration of the application
func makeHardCodedConfig() Config {
	// Create the new, empty node with its name and configurability parameters
	node := config.NewNode(actorName, actorName, true, true)
	// Set the configurability of the I/O ports
	node.SetPortsConfigurability(false, true)

	// Add I/O ports
	node.AddInputPort("reference-water-level", "base/Float64", "application/json", "", `{ "Body": { "Data": 0.75 } }`)
	node.AddInputPort("water-level", "base/Float64", "application/json", "well-water-level", `{ "Body": { "Data": 0.0 } }`)
	node.AddOutputPort("water-level-state", "base/Bool", "application/json", "well-water-upper-level-state")

	return Config{
		Node: node,
	}
}

// Start initializes and starts new actor-node application according to its configuration
func (a Application) Start() {

	// Start the additional components, if there is any
	// TODO

	// Start the node
	a.Node.Start()
}
