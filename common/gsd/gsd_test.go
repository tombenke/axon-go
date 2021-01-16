package gsd

import (
	"github.com/stretchr/testify/assert"
	"os"
	"sync"
	"syscall"
	"testing"
)

func TestRegister(t *testing.T) {
	gsdCbCalled := false

	wg := sync.WaitGroup{}

	// Register the callback handler
	Register(&wg, func(s os.Signal) {
		gsdCbCalled = true
	})

	// Sent TERM signal, then wait for termination
	syscall.Kill(syscall.Getpid(), syscall.SIGTERM)
	wg.Wait()

	// Checks if callback was called
	assert.True(t, gsdCbCalled)
}
