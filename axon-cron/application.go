package main

import (
	"github.com/tombenke/axon-go/common/actor/node"
	"github.com/tombenke/axon-go/common/config"
	"github.com/tombenke/axon-go/common/gsd"
	"github.com/tombenke/axon-go/common/log"
	"os"
	"sync"
)

const (
	actorName = "axon-cron"
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
	app.Node = node.NewNode(app.config.Node, getProcessorFun(app.config))

	return app
}

// makeHardCodedConfig returns with the built-in configuration of the application
func makeHardCodedConfig() Config {
	// Create the new, empty node with its name and configurability parameters
	node := config.NewNode(actorName, actorName, false, true)

	// Add I/O ports
	node.AddInputPort("trigger", "base/Bool", "application/json", "axon.cron.trigger", "")
	node.AddOutputPort("cron", "base/Any", "application/json", "axon.cron")

	return Config{
		Node:    node,
		CronDef: "@every 10s",
	}
}

// Start initializes and starts new actor-node application according to its configuration
func (a Application) Start(appWg *sync.WaitGroup) {

	appWg.Add(1)
	// Start the additional components, if there is any
	// TODO

	// Start the node
	nodeWg := sync.WaitGroup{}
	go a.Node.Start(&nodeWg)

	go func() {
		// Wait until the actor will be shut down
		<-a.done
		log.Logger.Infof("%s is shutting down", actorName)
		a.Node.Shutdown()
		nodeWg.Wait()
		appWg.Done()
	}()
}

// Shutdown stops the application process
func (a Application) Shutdown() {
	a.done <- true
}
