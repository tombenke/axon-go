package main

import (
	"github.com/tombenke/axon-go-common/gsd"
	"github.com/tombenke/axon-go-common/log"
	"github.com/tombenke/axon-go-common/messenger"
	messengerImpl "github.com/tombenke/axon-go-common/messenger/nats"
	"github.com/tombenke/axon-go/axon-orchestrator/actor-status"
	"github.com/tombenke/axon-go/axon-orchestrator/heartbeat"
	"os"
	"sync"
)

// Application struct represents the actor-node application
type Application struct {
	// config is the application configuration
	config Config

	// heartbeatGenerator is the Heartbeat Generator of the application
	heartbeatGenerator heartbeat.Generator

	// actorStatusRegistry is the Registry of actor statuses
	actorStatusRegistry actorStatus.Registry

	// doneCh is the channel to notify if the application must shut down
	doneCh chan struct{}

	// wg is the waitgroup for the application's goroutines
	wg *sync.WaitGroup

	// messsenger is the central messenger module used jointly by the the components of the application
	messenger messenger.Messenger
}

// Run creates a new application instance, and starts it
func Run(args []string) {
	// Create a new application instance using the CLI and config parameters
	a := NewApplication(args)

	// Start the application
	a.Start()

	// Wait until someone stops the application via a signal
	a.Wait()
}

// NewApplication creates a new actor-node application object
func NewApplication(args []string) Application {

	// Merge hard-coded configuration with CLI and file config if there is any
	config := GetConfig(appName, args)

	// Configure the global logger of the application according to the configuration
	log.SetLevelStr(config.LogLevel)
	log.SetFormatterStr(config.LogFormat)

	// Setup Messenger and connect to messaging
	config.Messenger.Logger = log.Logger
	config.Messenger.ClientID = appName
	config.Messenger.ClientName = appName
	config.Messenger.ClusterID = "test-cluster"
	messenger := messengerImpl.NewMessenger(config.Messenger)

	// Create the Heartbeat Generator
	heartbeatGenerator, heartbeatCh := heartbeat.NewGenerator(config.Heartbeat, messenger)

	// Create the Application
	app := Application{
		messenger:           messenger,
		heartbeatGenerator:  heartbeatGenerator,
		actorStatusRegistry: actorStatus.NewRegistry(heartbeatCh, messenger),
		config:              config,
		doneCh:              make(chan struct{}),
		wg:                  &sync.WaitGroup{},
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

	// Start the internal processes of the application
	go a.heartbeatGenerator.Start(a.wg)
	go a.actorStatusRegistry.Start(a.wg)

	// Let the application running
	a.wg.Add(1)
	go func() {
		// Wait until the actor will be shut down
		<-a.doneCh
		a.wg.Done()
	}()

	log.Logger.Infof("%s is started", appName)
}

// Shutdown stops the application process
func (a Application) Shutdown() {
	log.Logger.Infof("%s is shutting down", appName)
	a.actorStatusRegistry.Shutdown()
	a.heartbeatGenerator.Shutdown()
	close(a.doneCh)
}

// Wait waits until the internal components of the Application terminates
func (a Application) Wait() {
	a.wg.Wait()
}
