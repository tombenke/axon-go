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
	}, Default: `{"Body": {"Data": 0}}`},
	config.In{IO: config.IO{
		Name:           "power-consumption",
		Type:           "base/Float64",
		Representation: "application/json",
		Channel:        "",
	}, Default: `{"Body": {"Data": 0}}`},
	config.In{IO: config.IO{
		Name:           "power-input",
		Type:           "base/Float64",
		Representation: "application/json",
		Channel:        "",
	}, Default: ""},
	config.In{IO: config.IO{
		Name:           "water-input",
		Type:           "base/Float64",
		Representation: "application/json",
		Channel:        "",
	}, Default: ""},
	config.In{IO: config.IO{
		Name:           "water-need",
		Type:           "base/Float64",
		Representation: "application/json",
		Channel:        "",
	}, Default: ""},
}

var outputsCfg = config.Outputs{
	config.Out{IO: config.IO{
		Name:           "power-need",
		Type:           "base/Float64",
		Representation: "application/json",
		Channel:        "",
	}},
	config.Out{IO: config.IO{
		Name:           "water-output",
		Type:           "base/Float64",
		Representation: "application/json",
		Channel:        "",
	}},
	config.Out{IO: config.IO{
		Name:           "water-need",
		Type:           "base/Float64",
		Representation: "application/json",
		Channel:        "",
	}},
}
