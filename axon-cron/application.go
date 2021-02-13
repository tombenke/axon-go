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
	wg     *sync.WaitGroup
}

// Run creates a new application instance, and starts it
func Run(args []string) {
	// Create a new application instance using the CLI and config parameters
	a := NewApplication(args)

	// Start the axon node application
	a.Start()

	// Wait until someone stops the application via a signal
	a.Wait()
}

// NewApplication creates a new actor-node application object
func NewApplication(args []string) Application {
	app := Application{
		// Merge hard-coded configuration with CLI and file config if there is any
		config: GetConfig(actorName, getBuiltInConfig(), args),

		// Create the channel to notify if the application must shut down
		done: make(chan bool),
		wg:   &sync.WaitGroup{},
	}

	// Print config to the console in YAML format
	if app.config.PrintConfig {
		printResultingConfig(app.config)
	}

	// Create the Node
	app.Node = node.NewNode(app.config.Node, getProcessorFun(app.config))

	return app
}

// getBuiltInConfig returns with the built-in configuration of the application
func getBuiltInConfig() Config {
	// Create the new, empty node with its name and configurability parameters
	node := config.NewNode(actorName, actorName, false, true)

	// Add I/O ports. The actor has no inputs.
	node.AddOutputPort("cron", "base/Any", "application/json", "axon.cron")

	return Config{
		Node:      node,
		CronDef:   "@every 10s",
		Precision: "ms",
	}
}

// Start initializes and starts new actor-node application according to its configuration
func (a Application) Start() {

	// Register the graceful shutdown
	gsd.Register(a.wg, func(s os.Signal) {
		a.Shutdown()
	})

	// Start the node
	go a.Node.Start()

	// Setup the Cron Job
	cronDoneCh := make(chan bool)
	go startCron(a.Node, a.config.CronDef, a.wg, cronDoneCh)

	// Let the application running
	a.wg.Add(1)
	go func() {
		// Wait until the actor will be shut down
		<-a.done
		log.Logger.Infof("%s is shutting down", actorName)
		close(cronDoneCh)
		a.Node.Shutdown()
		a.Node.Wait()
		a.wg.Done()
	}()
}

// Shutdown stops the application process
func (a Application) Shutdown() {
	close(a.done)
}

// Wait waits until the internal components of the Application terminates
func (a Application) Wait() {
	a.wg.Wait()
}
