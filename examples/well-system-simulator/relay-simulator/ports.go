package main

import (
	"github.com/tombenke/axon-go/common/config"
)

var inputsCfg = config.Inputs{
	config.In{IO: config.IO{
		Name:           "dt",
		Type:           "base/Float64",
		Representation: "application/json",
		Channel:        "",
	}, Default: `{"Body": {"Data": 1000}}`},
	config.In{IO: config.IO{
		Name:           "relay-state",
		Type:           "base/Bool",
		Representation: "application/json",
		Channel:        "",
	}, Default: ""},
	config.In{IO: config.IO{
		Name:           "power-input",
		Type:           "base/Float64",
		Representation: "application/json",
		Channel:        "well-pump-relay.electric-power-input",
	}, Default: ""},
	config.In{IO: config.IO{
		Name:           "power-need",
		Type:           "base/Float64",
		Representation: "application/json",
		Channel:        "well-pump-power-need",
	}, Default: ""},
}

var outputsCfg = config.Outputs{
	config.Out{IO: config.IO{
		Name:           "power-output",
		Type:           "base/Float64",
		Representation: "application/json",
		Channel:        "well-pump-power",
	}},
	config.Out{IO: config.IO{
		Name:           "power-need",
		Type:           "base/Float64",
		Representation: "application/json",
		Channel:        "well-pump-relay.electric-power-need",
	}},
}
