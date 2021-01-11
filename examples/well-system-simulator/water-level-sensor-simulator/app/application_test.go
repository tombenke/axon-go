package app

import (
	"testing"
)

func TestApplication(t *testing.T) {
	// Create a new application instance
	a := NewApplication()

	// Start the axon node application with the test parameters
	testArgs := []string{}
	a.Start(testArgs)
}
