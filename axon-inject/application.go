package main

import (
	"github.com/tombenke/axon-go-common/actor/node"
	"github.com/tombenke/axon-go-common/gsd"
	"github.com/tombenke/axon-go-common/log"
	"os"
	"sync"
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
		config: GetConfig(actorName, args),

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

// Start initializes and starts new actor-node application according to its configuration
func (a Application) Start() {

	// Register the graceful shutdown
	sigsCh := gsd.Register(a.wg, func(s os.Signal) {
		a.Shutdown()
	})

	// Start the node
	nodeStartedCh := a.Node.Start()
	<-nodeStartedCh

	injectDoneCh := make(chan interface{})
	waitForShutdownCh := make(chan interface{})

	// Let the application running
	a.wg.Add(1)
	go func() {
		// Wait until the actor will be shut down
		close(waitForShutdownCh)
		<-a.done
		log.Logger.Infof("%s is shutting down", actorName)
		close(injectDoneCh)
		a.Node.Shutdown()
		a.Node.Wait()
		a.wg.Done()
	}()

	// Start the injector only when the application is ready to handle shut down signal
	<-waitForShutdownCh
	startInjector(a.Node, a.config, a.wg, injectDoneCh, sigsCh)
}

// Shutdown stops the application process
func (a Application) Shutdown() {
	close(a.done)
}

// Wait waits until the internal components of the Application terminates
func (a Application) Wait() {
	a.wg.Wait()
}
