package main

import (
	"flag"
	config "github.com/tombenke/axon-go/common/config"
)

const (
	actorName             = "axon-cron"
	defaultConfigFileName = "config.yml"
	defaultCronDef        = "@every 10s"
	defaultPrecision      = "ms"
)

// Config holds the configuration parameters of the actor node application
// It must inherit the configuration of the core `Node` object,
// and may contain other additional application-specific parameters.
type Config struct {
	// Node holds the config parameters that every actor node needs
	Node config.Node `yaml:"node"`

	// Show the help of the application
	ShowHelp bool `yaml:"showHelp"`

	// PrintConfig if true, then prints the resulting configuration to the console
	PrintConfig bool `yaml:"printConfig"`

	// CronDef is the settings for the cron-job
	CronDef string `yaml:"cronDef"`

	// The precision of the time value: ns, us, ms, s
	Precision string `yaml:"precision"`
}

// GetAppFlagSet returns with the flag-set of the application to parse the CLI parameters
func GetAppFlagSet(appName string, cfg *Config) *flag.FlagSet {
	fs := config.GetDefaultFlagSet(appName, &cfg.Node)

	fs.BoolVar(&cfg.ShowHelp, "h", false, "Show help message")
	fs.BoolVar(&cfg.ShowHelp, "help", false, "Show help message")

	fs.BoolVar(&cfg.PrintConfig, "p", false, "Print configuration parameters")
	fs.BoolVar(&cfg.PrintConfig, "print-config", false, "Print configuration parameters")

	fs.StringVar(&cfg.CronDef, "cron", cfg.CronDef, "Cron definition")
	fs.StringVar(&cfg.Precision, "precision", "ns", "The precision of time value: ns, us, ms, s")

	return fs
}

// builtInConfig returns with the built-in configuration of the application
func builtInConfig() Config {
	// Create the new, empty node with its name and configurability parameters
	node := config.NewNode(actorName, actorName, false, true)

	// Add I/O ports. The actor has no inputs.
	node.AddOutputPort("cron", "base/Any", "application/json", "axon.cron")

	return Config{
		Node:      node,
		CronDef:   "@every 10s",
		Precision: "ms",
	}
}

// mergeConfigs returns with the resulting config parameters set after merging them
//
// `builtInConfigContent` holds the configuration that the application defined as a baseline,
// `cliConfigContent` holds those configuration parameters, that origins from the default values,
// then extended by the config file, if there is any, then finally these were extended by the
// parameters from the environment and the CLI arguments.
// The most complex task of merging the I/O ports are done by the `config.MergeNodeConfigs()` function
// according to the values of `Extend` and `Modify` flags defined by the `builtInConfigContent`.
// The application needs to implement the merging of properties added by itself, on top of the Node parameters.
func mergeConfigs(builtInConfigContent Config, cliConfigContent Config) Config {
	resultingConfig := builtInConfigContent

	// Let the internal config module to manage the merging of the Node parameters
	resultingNode, err := config.MergeNodeConfigs(builtInConfigContent.Node, cliConfigContent.Node)
	if err != nil {
		panic(err.Error())
	}
	resultingConfig.Node = resultingNode

	//TODO: Add application-level merging tasks here if there is any
	resultingConfig.ShowHelp = cliConfigContent.ShowHelp
	resultingConfig.PrintConfig = cliConfigContent.PrintConfig
	resultingConfig.CronDef = cliConfigContent.CronDef
	resultingConfig.Precision = cliConfigContent.Precision

	return resultingConfig
}
