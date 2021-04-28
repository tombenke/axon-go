package main

import (
	"flag"
	"fmt"
	config "github.com/tombenke/axon-go-common/config"
	"regexp"
	"strconv"
	"time"
)

const (
	actorName             = "axon-inject"
	defaultConfigFileName = "config.yml"
	defaultPrecision      = "ms"
	defaultRepeat         = 1
	defaultMessage        = ""
	defaultDelay          = "0ms"
)

var defaultConfig = Config{
	Node:        config.GetDefaultNode(),
	PrintConfig: false,
	Precision:   defaultPrecision,
	Message:     defaultMessage,
	Repeat:      defaultRepeat,
	Delay:       defaultDelay,
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

	// Precision The precision of the time value: ns, us, ms, s
	Precision string `yaml:"precision"`

	// Message The message to inject into the network
	Message string `yaml:"message"`

	// Repeat The number of times the messages should be sent. Default is 1
	Repeat int `yaml:"repeat"`

	// Delay The delay the injector waits between two message sendings.
	// The format of the delay is <number><ns|us|ms|s>, for example: `150ms`.
	Delay string `yaml:"delay"`
}

// GetAppFlagSet returns with the flag-set of the application to parse the CLI parameters
func GetAppFlagSet(appName string, cfg *Config) *flag.FlagSet {
	fs := config.GetDefaultFlagSet(appName, &cfg.Node)

	fs.BoolVar(&cfg.ShowHelp, "h", false, "Show help message")
	fs.BoolVar(&cfg.ShowHelp, "help", false, "Show help message")

	fs.BoolVar(&cfg.PrintConfig, "p", false, "Print configuration parameters")
	fs.BoolVar(&cfg.PrintConfig, "print-config", false, "Print configuration parameters")

	fs.StringVar(&cfg.Precision, "precision", cfg.Precision, "The precision of time value: ns, us, ms, s")
	fs.StringVar(&cfg.Message, "m", cfg.Message, "The message to inject into the network")
	fs.StringVar(&cfg.Message, "message", cfg.Message, "The message to inject into the network")
	fs.IntVar(&cfg.Repeat, "repeat", cfg.Repeat, "The number of times the messages should be sent")
	fs.StringVar(&cfg.Delay, "delay", cfg.Delay, "The delay the injector waits between two message sendings. Format: <number><ns|us|ms|s>")

	return fs
}

// builtInConfig returns with the built-in configuration of the application
func builtInConfig() Config {
	// Create the new, empty node with its name and configurability parameters
	// nodeName, nodeType, extend, modify, presence, sync
	node := config.NewNode(actorName, actorName, false, true, true, false)

	// Add I/O ports.
	node.AddInputPort("inject", "base/Bytes", "text/plain", "", "")
	node.AddOutputPort("output", "base/Bytes", "text/plain", "axon-inject.output")

	return Config{
		Node:      node,
		Precision: defaultPrecision,
		Message:   defaultMessage,
		Repeat:    defaultRepeat,
		Delay:     defaultDelay,
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
	resultingConfig.Precision = cliConfigContent.Precision
	resultingConfig.Message = cliConfigContent.Message
	resultingConfig.Repeat = cliConfigContent.Repeat
	resultingConfig.Delay = cliConfigContent.Delay

	return resultingConfig
}

func parseDelay(delayStr string) (time.Duration, error) {
	re := regexp.MustCompile(`([0-9]+)([a-z]{1,2})`)
	results := re.FindStringSubmatch(delayStr)

	if len(results) != 3 {
		return 0, fmt.Errorf("Wrong delay string: `%s`", delayStr)
	}

	delayValue, err := strconv.ParseInt(results[1], 10, 64)
	if err != nil {
		return 0, fmt.Errorf("Wrong delay number format: `%s`", results[1])
	}

	switch results[2] {
	case "ns":
		return time.Duration(delayValue) * time.Nanosecond, nil
	case "us":
		return time.Duration(delayValue) * time.Microsecond, nil
	case "ms":
		return time.Duration(delayValue) * time.Millisecond, nil
	case "s":
		return time.Duration(delayValue) * time.Second, nil
	default:
		return time.Duration(0), fmt.Errorf("Wrong delay dimension: `%s`", results[2])
	}
}
