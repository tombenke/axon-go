package main

import (
	"github.com/tombenke/axon-go/common/gsd"
	"os"
	"sync"
	"syscall"
	"testing"
)

func TestApplication(t *testing.T) {
	// Create the axon node application with the test parameters
	testArgs := []string{}

	// Create a new application instance
	a := NewApplication(testArgs)

	// Start the axon node application
	wg := sync.WaitGroup{}
	/* TODO: Implement
	a.Start(&wg)
	*/

	gsd.Register(&wg, func(s os.Signal) {
		a.Shutdown()
	})
	syscall.Kill(syscall.Getpid(), syscall.SIGTERM)
	wg.Wait()
}
