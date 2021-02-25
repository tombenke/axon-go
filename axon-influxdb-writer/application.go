package main

import (
	"github.com/tombenke/axon-go/common/actor/node"
	"github.com/tombenke/axon-go/common/gsd"
	"github.com/tombenke/axon-go/common/log"
	"os"
	"sync"
)

// Application struct represents the actor-node application
type Application struct {
	config   Config
	wg       *sync.WaitGroup
	done     chan bool
	influxDb InfluxDb
	Node     node.Node
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
	config := GetConfig(actorName, args)
	influxDb := NewInfluxDbConnection(config.InfluxDb)

	app := Application{
		// Merge hard-coded configuration with CLI and file config if there is any
		config: config,

		// Create the channel to notify if the application must shut down
		done:     make(chan bool),
		wg:       &sync.WaitGroup{},
		influxDb: influxDb,
		Node:     node.NewNode(config.Node, getProcessorFun(config, influxDb)),
	}

	// Print config to the console in YAML format
	if app.config.PrintConfig {
		printResultingConfig(app.config)
	}

	return app
}

// Start initializes and starts new actor-node application according to its configuration
func (a Application) Start() {

	// Register the graceful shutdown
	gsd.Register(a.wg, func(s os.Signal) {
		a.Shutdown()
	})

	// Start the node
	go a.Node.Start()

	// Let the application running
	a.wg.Add(1)
	go func() {
		// Wait until the actor will be shut down
		<-a.done
		log.Logger.Infof("%s is shutting down", actorName)
		a.Node.Shutdown()
		a.Node.Wait()
		a.wg.Done()
	}()
}

// Shutdown stops the application process
func (a Application) Shutdown() {
	a.influxDb.Close()
	close(a.done)
}

// Wait waits until the internal components of the Application terminates
func (a Application) Wait() {
	a.wg.Wait()
}
