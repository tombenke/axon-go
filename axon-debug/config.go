package main

import (
	"flag"
	config "github.com/tombenke/axon-go/common/config"
)

const (
	actorName             = "axon-debug"
	defaultConfigFileName = "config.yml"
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

	// The printing format of the output. One of: json, json-indent, yaml, yml.
	DebugFormat string `yaml:"debugFormat"`
}

// GetAppFlagSet returns with the flag-set of the application to parse the CLI parameters
func GetAppFlagSet(appName string, cfg *Config) *flag.FlagSet {
	fs := config.GetDefaultFlagSet(appName, &cfg.Node)

	fs.BoolVar(&cfg.ShowHelp, "h", false, "Show help message")
	fs.BoolVar(&cfg.ShowHelp, "help", false, "Show help message")

	fs.BoolVar(&cfg.PrintConfig, "p", false, "Print configuration parameters")
	fs.BoolVar(&cfg.PrintConfig, "print-config", false, "Print configuration parameters")

	fs.StringVar(&cfg.DebugFormat, "debug-format", "json-indent", "The printing format of the output. One of: json, json-indent, yaml, yml")

	return fs
}

// builtInConfig returns with the built-in configuration of the application
func builtInConfig() Config {
	// Create the new, empty node with its name and configurability parameters
	node := config.NewNode(actorName, actorName, false, true)

	// Add I/O ports. The actor has no outputs.
	node.AddInputPort("input", "base/Any", "application/json", "axon-debug.input", "")

	return Config{
		Node: node,
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

	// Add application-level merging tasks here if there is any
	resultingConfig.ShowHelp = cliConfigContent.ShowHelp
	resultingConfig.PrintConfig = cliConfigContent.PrintConfig
	resultingConfig.DebugFormat = cliConfigContent.DebugFormat

	return resultingConfig
}
