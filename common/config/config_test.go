package config

import (
	//"flag"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Config struct {
	NodeConfig
	Precision string
}

// Parse the command line arguments and returns with the results as a structure
// It is a mock function that each of the implementations of the nodes includes.
func ParseCliArgs(defaultNodeName string, args []string) Config {

	config := Config{}

	// Get the dafault CLI args
	fs := GetDefaultFlagSet(defaultNodeName, &config.NodeConfig)

	// Extend the default set with node specific arguments
	fs.StringVar(&config.Precision, "p", "ns", "The precision of time value: ns, us, ms, s")

	// Add usage printer function
	// flag.Usage = func() { "Print usage..." }

	fs.Parse(args)

	// Handle the -h flag
	// if fs.showHelp {
	//     showUsageAndExit(0)
	// }

	fmt.Printf("parsed args: %v\n", fs.Args())
	fmt.Printf("config: %v\n", config)

	return config
}

func TestParseCliArgsWithDefaults(t *testing.T) {
	defaultNodeName := "node-name"
	c := ParseCliArgs(defaultNodeName, []string{})
	assert.Equal(t, c.Name, defaultNodeName)
	assert.Equal(t, c.LogLevel, defaultLogLevel)
	assert.Equal(t, c.LogFormat, defaultLogFormat)
	assert.Equal(t, c.Urls, defaultNatsUrls)
	assert.Equal(t, c.UserCreds, defaultNatsUserCreds)
	assert.Equal(t, c.Precision, "ns")
	assert.Equal(t, c.Inputs, *new(Inputs))
	assert.Equal(t, c.Outputs, *new(Outputs))
}

func TestConfigWithArgs(t *testing.T) {
	var args []string
	nodeName := "well-water-upper-level-sensor-simulator"
	logLevel := "error"

	args = append(args, "-name")
	args = append(args, nodeName)
	args = append(args, "-log-level")
	args = append(args, logLevel)
	args = append(args, "-in")
	args = append(args, "well-water-level:water-level:0.")
	args = append(args, "-in")
	args = append(args, "reference-water-level:well-water-upper-level:0.")
	args = append(args, "-out")
	args = append(args, "level-state:well-water-upper-level-state")

	c := ParseCliArgs(nodeName, args)
	assert.Equal(t, c.Name, nodeName)
	assert.Equal(t, c.LogLevel, logLevel)
	assert.Equal(t, c.Inputs, Inputs{In{IO: IO{Topic: "well-water-level", Name: "water-level"}, DefaultValue: "0."}, In{IO: IO{Topic: "reference-water-level", Name: "well-water-upper-level"}, DefaultValue: "0."}})
	assert.Equal(t, c.Outputs, Outputs{Out{IO: IO{Topic: "well-water-upper-level-state", Name: "level-state"}}})
}
