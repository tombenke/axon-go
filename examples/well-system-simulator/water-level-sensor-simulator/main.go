package main

import (
	"github.com/tombenke/axon-go/examples/well-system-simulator/water-level-sensor-simulator/app"
	"os"
)

func main() {
	// Create a new application instance
	a := app.NewApplication()

	// Start the axon node application using the CLI and config parameters
	a.Start(os.Args[1:])
}
