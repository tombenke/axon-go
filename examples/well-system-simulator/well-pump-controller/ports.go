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
		Name:           "well-water-upper-level-state",
		Type:           "base/Bool",
		Representation: "application/json",
		Channel:        "",
	}, Default: `{"Body": {"Data": false}}`},
	config.In{IO: config.IO{
		Name:           "well-water-lower-level-state",
		Type:           "base/Bool",
		Representation: "application/json",
		Channel:        "",
	}, Default: `{"Body": {"Data": false}}`},
	config.In{IO: config.IO{
		Name:           "buffer-tank-upper-level-state",
		Type:           "base/Bool",
		Representation: "application/json",
		Channel:        "",
	}, Default: `{"Body": {"Data": false}}`},
	config.In{IO: config.IO{
		Name:           "well-pump-controller-state",
		Type:           "base/String",
		Representation: "application/json",
		Channel:        "",
	}, Default: `{"Body": {"Data": "STANDBY"}}`},
}

var outputsCfg = config.Outputs{
	config.Out{IO: config.IO{
		Name:           "well-pump-relay-state",
		Type:           "base/Bool",
		Representation: "application/json",
		Channel:        "",
	}},
	config.Out{IO: config.IO{
		Name:           "well-pump-controller-state",
		Type:           "base/String",
		Representation: "application/json",
		Channel:        "",
	}},
}
