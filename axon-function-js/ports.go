package main

import (
	"github.com/tombenke/axon-go/common/config"
)

var inputsCfg = config.Inputs{
	config.In{
		IO: config.IO{
			Name:           "input",
			Type:           "base/Any",
			Representation: "application/json",
			Channel:        "axon-debug.input",
		},
		Default: "",
	},
}

var outputsCfg = config.Outputs{
	// This actor has NO outputs
	config.Out{
		IO: config.IO{
			Name:           "output",
			Type:           "base/Any",
			Representation: "application/json",
			Channel:        "axon-debug.output",
		},
	},
}
