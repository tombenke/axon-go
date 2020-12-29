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
		Name:           "max-power",
		Type:           "base/Float64",
		Representation: "application/json",
		Channel:        "",
	}, Default: `{"Body": {"Data": 2000.0}}`},
	config.In{IO: config.IO{
		Name:           "power-need",
		Type:           "base/Float64",
		Representation: "application/json",
		Channel:        "well-pump-relay.electric-power-need",
	}, Default: ""},
}

var outputsCfg = config.Outputs{
	config.Out{IO: config.IO{
		Name:           "power-output",
		Type:           "base/Float64",
		Representation: "application/json",
		Channel:        "well-pump-relay.electric-power-input",
	}},
}
