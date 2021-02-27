package main

import (
	"flag"
	config "github.com/tombenke/axon-go-common/config"
)

const (
	actorName                   = "axon-influxdb-writer"
	defaultConfigFileName       = "config.yml"
	defaultInfluxDbURL          = ""
	defaultInfluxDbToken        = ""
	defaultInfluxDbBucket       = ""
	defaultInfluxDbOrganization = ""
	defaultInfluxDbMeasurement  = ""
)

var defaultConfig = Config{
	Node:        config.GetDefaultNode(),
	PrintConfig: false,
}

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

	// InfluxDb holds the configuration parameters of the influxdb client
	InfluxDb InfluxDbConfig `yaml:"influxDb"`
}

// InfluxDbConf holds the configuration parameters for the time-series database server connection
type InfluxDbConfig struct {
	// URL is the URL of the InfluxDB server
	URL string `yaml:"url"`

	// Token holds the access token for the client to connect to the InfluxDB server
	Token string `yaml:"token"`

	// Bucket is the name of the database the records will be stored into
	Bucket string `yaml:"bucket"`

	// Organization is the organization of the client that uses the database
	Organization string `yaml:"organization"`

	// Measurement describes the data stored in with the associated fields
	Measurement string `yaml:"measurement"`
}

// GetAppFlagSet returns with the flag-set of the application to parse the CLI parameters
func GetAppFlagSet(appName string, cfg *Config) *flag.FlagSet {
	fs := config.GetDefaultFlagSet(appName, &cfg.Node)

	fs.BoolVar(&cfg.ShowHelp, "h", false, "Show help message")
	fs.BoolVar(&cfg.ShowHelp, "help", false, "Show help message")

	fs.BoolVar(&cfg.PrintConfig, "p", false, "Print configuration parameters")
	fs.BoolVar(&cfg.PrintConfig, "print-config", false, "Print configuration parameters")

	fs.StringVar(&cfg.InfluxDb.URL, "influxdb-url", defaultInfluxDbURL, "The URL of the influxdb server")
	fs.StringVar(&cfg.InfluxDb.Token, "influxdb-token", defaultInfluxDbToken, "The access token to the influxdb server")
	fs.StringVar(&cfg.InfluxDb.Bucket, "influxdb-bucket", defaultInfluxDbBucket, "The name of the bucket to store the data into")
	fs.StringVar(&cfg.InfluxDb.Organization, "influxdb-organization", defaultInfluxDbOrganization, "The name of the organization the client belongs to")
	fs.StringVar(&cfg.InfluxDb.Measurement, "influxdb-measurement", defaultInfluxDbMeasurement, "The name of the measurement")

	return fs
}

// builtInConfig returns with the built-in configuration of the application
func builtInConfig() Config {
	// Create the new, empty node with its name and configurability parameters:
	// nodeName, nodeType, extend, modify, presence, sync
	node := config.NewNode(actorName, actorName, true, true, true, false)

	// Add I/O ports. The actor has no outputs.
	//node.AddInputPort("input", "base/Float64", "application/json", "axon-influxdb-writer.input", "")

	return Config{
		Node: node,
		InfluxDb: InfluxDbConfig{
			URL:          defaultInfluxDbURL,
			Token:        defaultInfluxDbToken,
			Bucket:       defaultInfluxDbBucket,
			Organization: defaultInfluxDbOrganization,
			Measurement:  defaultInfluxDbMeasurement,
		},
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
	resultingConfig.InfluxDb = cliConfigContent.InfluxDb

	return resultingConfig
}
