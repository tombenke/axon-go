package app

import (
	"github.com/tombenke/axon-go/common/actor/node"
	"github.com/tombenke/axon-go/common/config"
	"github.com/tombenke/axon-go/common/gsd"
	"github.com/tombenke/axon-go/common/log"
	"os"
	"sync"
)

const (
	actorName = "water-level-sensor-simulator"
)

// Application struct represents the actor-node application
type Application struct {
	Node   node.Node
	config Config
	done   chan bool
}

// Run creates a new application instance, and starts it
func Run(args []string) {
	// Create a new application instance using the CLI and config parameters
	a := NewApplication(args)

	// Start the axon node application
	wg := sync.WaitGroup{}
	a.Start(&wg)

	gsd.Register(&wg, func(s os.Signal) {
		a.Shutdown()
	})

	wg.Wait()
}

// NewApplication creates a new actor-node application object
func NewApplication(args []string) Application {
	app := Application{}

	// Create the channel to notify if the application must shut down
	app.done = make(chan bool)

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
	app.Node = node.NewNode(app.config.Node, ProcessorFun)

	return app
}

// makeHardCodedConfig returns with the built-in configuration of the application
func makeHardCodedConfig() Config {
	// Create the new, empty node with its name and configurability parameters
	// nodeName, nodeType, extend, modify, presence, sync
	node := config.NewNode(actorName, actorName, false, true, true, true)

	// Add I/O ports
	node.AddInputPort("reference-water-level", "base/Float64", "application/json", "", `{ "Body": { "Data": 0.75 } }`)
	node.AddInputPort("water-level", "base/Float64", "application/json", "well-water-level", `{ "Body": { "Data": 0.0 } }`)
	node.AddOutputPort("water-level-state", "base/Bool", "application/json", "well-water-upper-level-state")

	return Config{
		Node: node,
	}
}

// Start initializes and starts new actor-node application according to its configuration
func (a Application) Start(appWg *sync.WaitGroup) {

	appWg.Add(1)
	// Start the additional components, if there is any
	// TODO

	// Start the node
	go a.Node.Start()

	go func() {
		// Wait until the actor will be shut down
		<-a.done
		log.Logger.Infof("%s is shutting down", actorName)
		a.Node.Shutdown()
		a.Node.Wait()
		appWg.Done()
	}()
}

// Shutdown stops the application process
func (a Application) Shutdown() {
	a.done <- true
}
