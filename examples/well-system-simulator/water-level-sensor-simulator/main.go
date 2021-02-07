package main

import (
	"github.com/tombenke/axon-go/examples/well-system-simulator/water-level-sensor-simulator/app"
	"os"
)

func main() {
	app.Run(os.Args[1:])
}
