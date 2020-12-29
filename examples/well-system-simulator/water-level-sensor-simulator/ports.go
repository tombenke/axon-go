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
		Name:           "reference-water-level",
		Type:           "base/Float64",
		Representation: "application/json",
		Channel:        "",
	}, Default: `{"Body": {"Data": 0.75}}`},
	config.In{IO: config.IO{
		Name:           "water-level",
		Type:           "base/Float64",
		Representation: "application/json",
		Channel:        "well-water-buffer-tank-level",
	}, Default: ""},
}

var outputsCfg = config.Outputs{
	config.Out{IO: config.IO{
		Name:           "water-level-state",
		Type:           "base/Bool",
		Representation: "application/json",
		Channel:        "buffer-water-tank-upper-level-state",
	}},
}
