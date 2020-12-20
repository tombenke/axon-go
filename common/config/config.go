package config

import (
	"flag"
	//"fmt"
	"github.com/tombenke/axon-go/common/messenger"
	"os"
)

// Get the value of the `envVarName` environment variable and return with it.
// If there is no such variable defined in the environment, then return with the `defaultValue`.
func GetEnvWithDefault(envVarName string, defaultValue string) string {
	value, ok := os.LookupEnv(envVarName)
	if !ok {
		value = defaultValue
	}
	return value
}

const (
	nodeNameHelp   = "The name of the node"
	nodeNameEnvVar = "NODE_NAME"

	logLevelHelp    = "The log level: panic | fatal | error | warning | info | debug | trace"
	logLevelEnvVar  = "LOG_LEVEL"
	defaultLogLevel = "info"

	logFormatHelp    = "The log format: json | text"
	logFormatEnvVar  = "LOG_FORMAT"
	defaultLogFormat = "json"

	natsUrlsEnvVar  = "NATS_URL"
	natsUrlsHelp    = "The NATS server's URLs (separated by comma)"
	defaultNatsUrls = "localhost:4222"

	natsUserCredsHelp    = "User Credentials"
	natsUserCredsEnvVar  = "NATS_CREDENTIALS"
	defaultNatsUserCreds = ""

	inputsHelp  = "Input. Format: [<topic>]:<name>:[<default-value>]"
	outputsHelp = "Output. Format: <name>:[<topic-name>]"
)

// The default config struct that every axon actor node inherits
type NodeConfig struct {
	messenger.Config

	Name      string
	LogLevel  string
	LogFormat string
	//Inputs    arrayFlags
	//Outputs   arrayFlags
	Inputs  Inputs
	Outputs Outputs
}

func GetDefaultFlagSet(defaultNodeName string, config *NodeConfig) *flag.FlagSet {
	fs := flag.NewFlagSet("fs-name", flag.PanicOnError)

	var showHelp bool
	fs.BoolVar(&showHelp, "h", false, "Show help message")

	fs.StringVar(&(*config).Name, "n", GetEnvWithDefault(nodeNameEnvVar, defaultNodeName), nodeNameHelp)
	fs.StringVar(&(*config).Name, "name", GetEnvWithDefault(nodeNameEnvVar, defaultNodeName), nodeNameHelp)

	fs.StringVar(&(*config).LogLevel, "l", GetEnvWithDefault(logLevelEnvVar, defaultLogLevel), logLevelHelp)
	fs.StringVar(&(*config).LogLevel, "log-level", GetEnvWithDefault(logLevelEnvVar, defaultLogLevel), logLevelHelp)

	fs.StringVar(&(*config).LogFormat, "f", GetEnvWithDefault(logFormatEnvVar, defaultLogFormat), logFormatHelp)
	fs.StringVar(&(*config).LogFormat, "log-format", GetEnvWithDefault(logFormatEnvVar, defaultLogFormat), logFormatHelp)

	fs.StringVar(&(*config).Urls, "u", GetEnvWithDefault(natsUrlsEnvVar, defaultNatsUrls), natsUrlsHelp)
	fs.StringVar(&(*config).Urls, "nats-urls", GetEnvWithDefault(natsUrlsEnvVar, defaultNatsUrls), natsUrlsHelp)

	fs.StringVar(&(*config).UserCreds, "c", GetEnvWithDefault(natsUserCredsEnvVar, defaultNatsUserCreds), natsUserCredsHelp)
	fs.StringVar(&(*config).UserCreds, "creds", GetEnvWithDefault(natsUserCredsEnvVar, defaultNatsUserCreds), natsUserCredsHelp)

	fs.Var(&(*config).Inputs, "in", inputsHelp)
	fs.Var(&(*config).Outputs, "out", outputsHelp)

	return fs
}
