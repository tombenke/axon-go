package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// invalid input strings
var invalidIns []string = []string{
	"",           // no name string
	":",          // empty name string
	"::",         // empty name string
	"topic::",    // empty name string
	"::0.1",      // empty name string
	"topic::0.1", // empty name string
}

type validIn struct {
	Arg      string
	Expected In
}

var validIns []validIn = []validIn{
	validIn{"name", In{IO{"name", "name"}, ""}},               // name only
	validIn{"name:0.1", In{IO{"name", "name"}, "0.1"}},        // name and default value
	validIn{":name:0.1", In{IO{"name", "name"}, "0.1"}},       // name and default value
	validIn{"topic:name:", In{IO{"topic", "name"}, ""}},       // topic and name
	validIn{"topic:name:0.1", In{IO{"topic", "name"}, "0.1"}}, // full
}

const notEqualMsg string = "The two objects should be the equal!"

// Test input args
func TestParseInArgs(t *testing.T) {
	assert := assert.New(t)

	// Test valid cases
	for _, i := range validIns {
		assert.Equal(parseIn(i.Arg), i.Expected, notEqualMsg)
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
	"",       // no name string
	":",      // empty name string
	":topic", // empty name string
}

type validOut struct {
	Arg      string
	Expected Out
}

var validOuts []validOut = []validOut{
	validOut{"name", Out{IO{"name", "name"}}},        // name only
	validOut{"name:", Out{IO{"name", "name"}}},       // name and default value
	validOut{"name:topic", Out{IO{"topic", "name"}}}, // name and topic name
}

// Test output args
func TestParseOutArgs(t *testing.T) {
	assert := assert.New(t)

	// Test valid cases
	for _, i := range validOuts {
		assert.Equal(parseOut(i.Arg), i.Expected, notEqualMsg)
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
