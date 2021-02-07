package main

import (
	"github.com/tombenke/axon-go/common/config"
)

var inputsCfg = config.Inputs{
	// This actor has NO inputs
}

var outputsCfg = config.Outputs{
	config.Out{IO: config.IO{
		Name:           "cron",
		Type:           "base/Any",
		Representation: "application/json",
		Channel:        "axon.cron",
	}},
}
