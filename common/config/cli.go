// Package config provides the implementation of the commonly used,
// generic CLI parameter parsing and configuration management routines
package config

import (
	"flag"
	"os"
)

// GetEnvWithDefault gets the value of the `envVarName` environment variable and return with it.
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
	defaultLogFormat = "text"

	messagingUrlsEnvVar = "MESSAGING_URL"
	messagingUrlsHelp   = "The Messaging server's URLs (separated by comma)"
	defaultMessagingURL = "localhost:4222"

	messagingUserCredsHelp    = "User Credentials"
	messagingUserCredsEnvVar  = "MESSAGING_CREDENTIALS"
	defaultMessagingUserCreds = ""

	inputsHelp  = "Input. Format: <name>[|<channel>[|<type>|<representation>|<default>]]"
	outputsHelp = "Output. Format: <name>[|<channel>[|<type>|<representation>]]"
)

// GetDefaultFlagSet returns with the default values of the generic configuration parameters
func GetDefaultFlagSet(defaultNodeName string, config *Node) *flag.FlagSet {
	fs := flag.NewFlagSet("fs-name", flag.ContinueOnError)

	fs.StringVar(&(*config).Name, "n", GetEnvWithDefault(nodeNameEnvVar, defaultNodeName), nodeNameHelp)
	fs.StringVar(&(*config).Name, "name", GetEnvWithDefault(nodeNameEnvVar, defaultNodeName), nodeNameHelp)

	fs.StringVar(&(*config).LogLevel, "l", GetEnvWithDefault(logLevelEnvVar, (*config).LogLevel), logLevelHelp)
	fs.StringVar(&(*config).LogLevel, "log-level", GetEnvWithDefault(logLevelEnvVar, (*config).LogLevel), logLevelHelp)

	fs.StringVar(&(*config).LogFormat, "f", GetEnvWithDefault(logFormatEnvVar, (*config).LogFormat), logFormatHelp)
	fs.StringVar(&(*config).LogFormat, "log-format", GetEnvWithDefault(logFormatEnvVar, (*config).LogFormat), logFormatHelp)

	fs.StringVar(&(*config).Messenger.Urls, "u", GetEnvWithDefault(messagingUrlsEnvVar, (*config).Messenger.Urls), messagingUrlsHelp)
	fs.StringVar(&(*config).Messenger.Urls, "messaging-urls", GetEnvWithDefault(messagingUrlsEnvVar, (*config).Messenger.Urls), messagingUrlsHelp)

	fs.StringVar(&(*config).Messenger.UserCreds, "c", GetEnvWithDefault(messagingUserCredsEnvVar, (*config).Messenger.UserCreds), messagingUserCredsHelp)
	fs.StringVar(&(*config).Messenger.UserCreds, "creds", GetEnvWithDefault(messagingUserCredsEnvVar, (*config).Messenger.UserCreds), messagingUserCredsHelp)

	fs.StringVar(&(*config).ConfigFileName, "config", "config.yml", "Config file name")

	fs.Var(&(*config).Ports.Inputs, "in", inputsHelp)
	fs.Var(&(*config).Ports.Outputs, "out", outputsHelp)

	return fs
}
