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
		Name:           "pump-capacity",
		Type:           "base/Float64",
		Representation: "application/json",
		Channel:        "",
	}, Default: `{"Body": {"Data": 2.0}}`}, // [m3/h]
	config.In{IO: config.IO{
		Name:           "power-consumption",
		Type:           "base/Float64",
		Representation: "application/json",
		Channel:        "",
	}, Default: `{"Body": {"Data": 800}}`}, // [W]
	config.In{IO: config.IO{
		Name:           "power-input",
		Type:           "base/Float64",
		Representation: "application/json",
		Channel:        "well-pump-power",
	}, Default: ""},
	config.In{IO: config.IO{
		Name:           "water-input",
		Type:           "base/Float64",
		Representation: "application/json",
		Channel:        "well-water-output",
	}, Default: ""},
	config.In{IO: config.IO{
		Name:           "water-need",
		Type:           "base/Float64",
		Representation: "application/json",
		Channel:        "well-water-buffer-tank-consumption-need",
	}, Default: ""},
}

var outputsCfg = config.Outputs{
	config.Out{IO: config.IO{
		Name:           "power-need",
		Type:           "base/Float64",
		Representation: "application/json",
		Channel:        "well-pump-power-need",
	}},
	config.Out{IO: config.IO{
		Name:           "water-output",
		Type:           "base/Float64",
		Representation: "application/json",
		Channel:        "well-pump-water-output",
	}},
	config.Out{IO: config.IO{
		Name:           "water-need",
		Type:           "base/Float64",
		Representation: "application/json",
		Channel:        "well-pump-water-need",
	}},
}
