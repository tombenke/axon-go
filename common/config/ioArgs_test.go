package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// invalid input strings
var invalidIns []string = []string{
	"",              // no name string
	"|",             // empty name string
	"||",            // empty name string
	"|channel||",    // empty name string
	"||0.1",         // empty name string
	"|channel||0.1", // empty name string
	"name||||||",    // Wrong number of arguments
	"name|||",       // Wrong number of arguments
}

type validIn struct {
	Arg      string
	Expected In
}

var validIns []validIn = []validIn{
	validIn{"name", In{IO{"name", DefaultType, DefaultRepresentation, ""}, ""}},                                                 // name only
	validIn{"name||||0.1", In{IO{"name", DefaultType, DefaultRepresentation, ""}, "0.1"}},                                       // name and default value
	validIn{"name||||0.1", In{IO{"name", DefaultType, DefaultRepresentation, ""}, "0.1"}},                                       // name and default value
	validIn{"name|channel|||", In{IO{"name", DefaultType, DefaultRepresentation, "channel"}, ""}},                               // channel and name
	validIn{"name|channel|||false", In{IO{"name", DefaultType, DefaultRepresentation, "channel"}, "false"}},                     // channel and name
	validIn{"name|channel|base/Bool|application/json|true", In{IO{"name", "base/Bool", "application/json", "channel"}, "true"}}, // full
}

// Test input args
func TestParseInArgs(t *testing.T) {
	assert := assert.New(t)

	// Test valid cases
	for _, i := range validIns {
		assert.Equal(i.Expected, parseIn(i.Arg))
	}

	// Test invalid cases
	for _, i := range invalidIns {
		assert.Panics(
			func() {
				parseIn(i)
			},
			"It should panic!",
		)
	}
}

// invalid output strings
var invalidOuts []string = []string{
	"",          // no name string
	"|",         // empty name string
	"|channel",  // empty name string
	"name||",    // Wrong number of arguments
	"name|||||", // Wrong number of arguments
}

type validOut struct {
	Arg      string
	Expected Out
}

var validOuts []validOut = []validOut{
	validOut{"name", Out{IO{"name", DefaultType, DefaultRepresentation, ""}}},
	validOut{"name|", Out{IO{"name", DefaultType, DefaultRepresentation, ""}}},
	validOut{"name|channel|base/Bool|application/json", Out{IO{"name", "base/Bool", "application/json", "channel"}}},
}

// Test output args
func TestParseOutArgs(t *testing.T) {
	assert := assert.New(t)

	// Test valid cases
	for _, i := range validOuts {
		assert.Equal(parseOut(i.Arg), i.Expected)
	}

	// Test invalid cases
	for _, i := range invalidOuts {
		assert.Panics(
			func() {
				parseOut(i)
			},
			"It should panic!",
		)
	}
}
