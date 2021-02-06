package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Config struct {
	Node
	Precision string
}

// parseCliArgs parses the command line arguments and returns with the results as a structure
// It is a mock function that each of the implementations of the nodes includes.
func parseCliArgs(defaultNodeName string, args []string) Config {

	config := Config{
		Node:      GetDefaultNode(),
		Precision: "",
	}

	// Get the dafault CLI args
	fs := GetDefaultFlagSet(defaultNodeName, &config.Node)

	// Extend the default set with node specific arguments
	fs.StringVar(&config.Precision, "p", "ns", "The precision of time value: ns, us, ms, s")

	if err := fs.Parse(args); err != nil {
		panic(err)
	}

	return config
}

func TestParseCliArgsWithDefaults(t *testing.T) {
	defaultNodeName := "node-name"
	c := parseCliArgs(defaultNodeName, []string{})
	assert.Equal(t, defaultNodeName, c.Name)
	assert.Equal(t, defaultLogLevel, c.LogLevel)
	assert.Equal(t, defaultLogFormat, c.LogFormat)
	assert.Equal(t, defaultMessagingURL, c.Messenger.Urls)
	assert.Equal(t, defaultMessagingUserCreds, c.Messenger.UserCreds)
	assert.Equal(t, "ns", c.Precision)
	assert.Equal(t, *new(Inputs), c.Ports.Inputs)
	assert.Equal(t, *new(Outputs), c.Ports.Outputs)
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
	args = append(args, `water-level|well-water-level|||{"Body": {"Data": 0.}}`)
	args = append(args, "-in")
	args = append(args, `reference-water-level|well-water-upper-level|||{"Body": {"Data": 0.}}`)
	args = append(args, "-out")
	args = append(args, "level-state|well-water-upper-level-state")

	c := parseCliArgs(nodeName, args)
	assert.Equal(t, nodeName, c.Name)
	assert.Equal(t, logLevel, c.LogLevel)
	assert.Equal(t,
		Inputs{
			In{IO: IO{Name: "water-level", Channel: "well-water-level", Type: DefaultType, Representation: DefaultRepresentation}, Default: `{"Body": {"Data": 0.}}`},
			In{IO: IO{Name: "reference-water-level", Channel: "well-water-upper-level", Type: DefaultType, Representation: DefaultRepresentation}, Default: `{"Body": {"Data": 0.}}`}},
		c.Ports.Inputs)
	assert.Equal(t,
		Outputs{
			Out{IO: IO{Name: "level-state", Channel: "well-water-upper-level-state", Type: DefaultType, Representation: DefaultRepresentation}}},
		c.Ports.Outputs)
}
