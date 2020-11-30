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

// Test with empty string
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
