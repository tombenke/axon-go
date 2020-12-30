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
		Name:           "well-backfill-capacity",
		Type:           "base/Float64",
		Representation: "application/json",
		Channel:        "",
	}, Default: `{"Body": {"Data": 0.6}}`}, // 0.6 [m3/h]
	config.In{IO: config.IO{
		Name:           "well-cross-section",
		Type:           "base/Float64",
		Representation: "application/json",
		Channel:        "",
	}, Default: `{"Body": {"Data": 0.0201056 }}`}, // Math.pow(0.16 / 2., 2) * Math.PI [m2]
	config.In{IO: config.IO{
		Name:           "max-water-level",
		Type:           "base/Float64",
		Representation: "application/json",
		Channel:        "",
	}, Default: `{"Body": {"Data": -30 }}`}, // -30 [m]
	config.In{IO: config.IO{
		Name:           "min-water-level",
		Type:           "base/Float64",
		Representation: "application/json",
		Channel:        "",
	}, Default: `{"Body": {"Data": -36 }}`}, // -36 [m]
	config.In{IO: config.IO{
		Name:           "water-need",
		Type:           "base/Float64",
		Representation: "application/json",
		Channel:        "well-pump-water-need",
	}, Default: ""},
	config.In{IO: config.IO{
		Name:           "water-level",
		Type:           "base/Float64",
		Representation: "application/json",
		Channel:        "well-water-level",
	}, Default: `{"Body": {"Data": -30 }}`}, // maxWaterLevel (the well is full by default)
}

var outputsCfg = config.Outputs{
	config.Out{IO: config.IO{
		Name:           "water-level",
		Type:           "base/Float64",
		Representation: "application/json",
		Channel:        "well-water-level",
	}},
	config.Out{IO: config.IO{
		Name:           "water-output",
		Type:           "base/Float64",
		Representation: "application/json",
		Channel:        "well-water-output",
	}},
}
