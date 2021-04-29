package main

import (
	"github.com/tombenke/axon-go-common/gsd"
	"github.com/tombenke/axon-go-common/log"
	"github.com/tombenke/axon-go-common/messenger"
	messengerImpl "github.com/tombenke/axon-go-common/messenger/nats"
	"github.com/tombenke/axon-go/axon-tui/epn"
	"github.com/tombenke/axon-go/axon-tui/ui"
	"os"
	"sync"
)

// Application struct represents the actor-node application
type Application struct {
	// config is the application configuration
	config Config

	// epnStatus is the EPN Status Observer
	epnStatus epn.Status

	// The UI part of the application
	ui ui.UI

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

	// Merge default configuration with CLI and file config if there is any
	config := GetConfig(appName, args)

	// Print config to the console in YAML format
	if config.PrintConfig {
		printResultingConfig(config)
		os.Exit(0)
	}

	// Configure the global logger of the application according to the configuration
	log.SetLevelStr(config.LogLevel)
	log.SetFormatterStr(config.LogFormat)

	// Setup Messenger and connect to messaging
	messenger := messengerImpl.NewMessenger(config.Messenger)

	// Create the EPN Status Observer
	epnStatus := epn.NewStatus(config.EPNStatusChannel, messenger)
	ui := ui.New(&epnStatus)

	// Create the Application
	app := Application{
		messenger: messenger,
		config:    config,
		epnStatus: epnStatus,
		ui:        ui,
		doneCh:    make(chan struct{}),
		wg:        &sync.WaitGroup{},
	}

	return app
}

// Start initializes and starts new actor-node application according to its configuration
func (a Application) Start() {

	// Register the graceful shutdown
	sigsCh := gsd.Register(a.wg, func(s os.Signal) {
		a.Shutdown()
	})

	// Start the internal processes of the application
	go a.epnStatus.Start(a.wg)
	go a.ui.Start(a.wg, sigsCh)

	// Let the application running
	a.wg.Add(1)
	go func() {
		// Wait until the application will be shut down
		<-a.doneCh
		a.wg.Done()
	}()

	log.Logger.Debugf("%s is started", appName)
}

// Shutdown stops the application process
func (a Application) Shutdown() {
	log.Logger.Debugf("%s is shutting down", appName)

	// Shuts down the internal processes of the application
	a.epnStatus.Shutdown()
	a.ui.Shutdown()

	close(a.doneCh)
}

// Wait waits until the internal components of the Application terminates
func (a Application) Wait() {
	a.wg.Wait()
}
