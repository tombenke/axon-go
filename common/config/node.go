package config

import (
	"github.com/tombenke/axon-go/common/messenger"
)

// Node is the default config struct that every axon actor node inherits
type Node struct {
	messenger.Config

	Name      string
	LogLevel  string
	LogFormat string
	Inputs    Inputs
	Outputs   Outputs
}
