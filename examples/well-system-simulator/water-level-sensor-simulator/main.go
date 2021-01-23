package main

import (
	"github.com/tombenke/axon-go/examples/well-system-simulator/water-level-sensor-simulator/app"
	"os"
)

func main() {
	// Create a new application instance using the CLI and config parameters
	a := app.NewApplication(os.Args[1:])

	// Start the axon node application
	a.Start()
}
