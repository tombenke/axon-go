package main

import (
	//"github.com/stretchr/testify/assert"
	//"syscall"
	"testing"
)

func TestApplicationStartStop(t *testing.T) {
	// Create the axon node application with the test parameters
	testArgs := []string{}

	// Create a new application instance
	a := NewApplication(testArgs)

	// Start the axon node application
	a.Start()

	// Stop the application
	//err := syscall.Kill(syscall.Getpid(), syscall.SIGTERM)
	//assert.Nil(t, err)

	// Wait until the app really stops
	a.Wait()
}
