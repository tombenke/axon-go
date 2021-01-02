package testing

import (
	"github.com/sirupsen/logrus"
	"github.com/tombenke/axon-go/common/msgs"
)

// Logger is a global logger for testing
var Logger logrus.Logger

// TestCase holds the input and output messages of one test case for a processor function.
// The inputs is put onto the context that the processor function is using
// during the execution of one specific test-case.
// The outputs hold the expected messages the processor function has to place
// during the test to the outputs port via the context.
type TestCase struct {
	Inputs  TestCaseMsgs
	Outputs TestCaseMsgs
}

// TestCases holds an array of TestCases for a specific processor function.
type TestCases []TestCase

// TestCaseMsgs is a type that TestCase descriptor structures use for
// both the inputs and outputs messages.
// This is a map of messages where the portname is the key.
type TestCaseMsgs map[string]msgs.Message
