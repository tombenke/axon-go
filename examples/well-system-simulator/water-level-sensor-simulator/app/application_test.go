package app

import (
	"testing"
)

func TestApplication(t *testing.T) {
	// Create the axon node application with the test parameters
	testArgs := []string{}

	// Create a new application instance
	a := NewApplication(testArgs)

	// Start the application
	a.Start()
}
