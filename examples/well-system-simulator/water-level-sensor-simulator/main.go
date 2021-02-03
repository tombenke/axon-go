package main

import (
	"github.com/tombenke/axon-go/examples/well-system-simulator/water-level-sensor-simulator/app"
	"os"
	"sync"

	"github.com/tombenke/axon-go/common/gsd"
)

func main() {
	// Create a new application instance using the CLI and config parameters
	a := app.NewApplication(os.Args[1:])

	// Start the axon node application
	wg := sync.WaitGroup{}
	a.Start(&wg)

	gsd.Register(&wg, func(s os.Signal) {
		a.Shutdown()
	})

	wg.Wait()
}
