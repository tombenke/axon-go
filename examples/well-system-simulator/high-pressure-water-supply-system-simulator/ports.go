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
		Name:           "buffer-volume",
		Type:           "base/Float64",
		Representation: "application/json",
		Channel:        "",
	}, Default: `{"Body": {"Data": 0.01}}`}, // 10 l
	config.In{IO: config.IO{
		Name:           "max-pressure",
		Type:           "base/Float64",
		Representation: "application/json",
		Channel:        "",
	}, Default: `{"Body": {"Data": 3.5}}`}, // 3.5 bar
	config.In{IO: config.IO{
		Name:           "min-pressure",
		Type:           "base/Float64",
		Representation: "application/json",
		Channel:        "",
	}, Default: `{"Body": {"Data": 1.0}}`}, // 1.0 bar
	config.In{IO: config.IO{
		Name:           "water-output-need",
		Type:           "base/Float64",
		Representation: "application/json",
		Channel:        "",
	}, Default: ""},
	config.In{IO: config.IO{
		Name:           "water-input",
		Type:           "base/Float64",
		Representation: "application/json",
		Channel:        "well-water-buffer-tank-water-output",
	}, Default: ""},
	config.In{IO: config.IO{
		Name:           "water-buffer-tank-level",
		Type:           "base/Float64",
		Representation: "application/json",
		Channel:        "well-water-buffer-tank-level",
	}, Default: ""},
}

var outputsCfg = config.Outputs{
	config.Out{IO: config.IO{
		Name:           "water-output",
		Type:           "base/Float64",
		Representation: "application/json",
		Channel:        "",
	}},
	config.Out{IO: config.IO{
		Name:           "water-input-need",
		Type:           "base/Float64",
		Representation: "application/json",
		Channel:        "high-pressure-wss-input-need",
	}},
}
