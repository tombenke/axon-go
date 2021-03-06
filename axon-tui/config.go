package main

import (
	"flag"
	"github.com/tombenke/axon-go-common/messenger"
	"os"
)

const (
	appName       = "axon-tui"
	appNameHelp   = "The name of the application"
	appNameEnvVar = "APP_NAME"

	defaultConfigFileName = "config.yml"

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

	messagingClusterIDHelp    = "Cluster ID of messaging"
	messagingClusterIDEnvVar  = "MESSAGING_CLUSTER_ID"
	defaultMessagingClusterID = ""

	EPNStatusChannelHelp    = "The name of the epn-status channel"
	EPNStatusChannelEnvVar  = "EPN_STATUS_CHANNEL"
	defaultEPNStatusChannel = "epn-status"
)

var defaultConfig = Config{
	EPNStatusChannel: defaultEPNStatusChannel,
	Name:             appName,
	ConfigFileName:   defaultConfigFileName,
	LogLevel:         defaultLogLevel,
	LogFormat:        defaultLogFormat,
	ShowHelp:         false,
	PrintConfig:      false,
	Messenger: messenger.Config{
		Urls:      defaultMessagingURL,
		UserCreds: defaultMessagingUserCreds,
		ClusterID: defaultMessagingClusterID,
	},
}

// Config holds the configuration parameters of the actor node application
// It must inherit the configuration of the core `Node` object,
// and may contain other additional application-specific parameters.
type Config struct {

	// EPNStatusChannel is the name of the epn-status channel
	EPNStatusChannel string `yaml:"epnStatusChannel"`

	// Name is the name of the node. It should be unique in a specific network
	Name string `yaml:"name"`

	// ConfigFileName is the name of the config file to load
	// the configuration parameters of the application.
	// Its default value is `config.yml`.
	// It is optional to use config files. When the node starts, it tries to find the config file
	// identified by this property, and loads it it it has been found.
	ConfigFileName string `yaml:"configFileName"`

	// LogLevel is the log level of the application
	LogLevel string `yaml:"logLevel"`

	// LogFormat the log format of the application
	LogFormat string `yaml:"logFormat"`

	// Show the help of the application
	ShowHelp bool `yaml:"showHelp"`

	// PrintConfig if true, then prints the resulting configuration to the console
	PrintConfig bool `yaml:"printConfig"`

	// Messenger holds the configuration parameters of the messaging middleware
	Messenger messenger.Config `yaml:"messenger"`
}

// GetAppFlagSet returns with the flag-set of the application to parse the CLI parameters
func GetAppFlagSet(appName string, cfg *Config) *flag.FlagSet {
	fs := flag.NewFlagSet(appName, flag.ContinueOnError)

	fs.BoolVar(&cfg.ShowHelp, "h", false, "Show help message")
	fs.BoolVar(&cfg.ShowHelp, "help", false, "Show help message")

	fs.BoolVar(&cfg.PrintConfig, "p", false, "Print configuration parameters")
	fs.BoolVar(&cfg.PrintConfig, "print-config", false, "Print configuration parameters")

	fs.StringVar(&cfg.LogLevel, "l", GetEnvWithDefault(logLevelEnvVar, cfg.LogLevel), logLevelHelp)
	fs.StringVar(&cfg.LogLevel, "log-level", GetEnvWithDefault(logLevelEnvVar, cfg.LogLevel), logLevelHelp)

	fs.StringVar(&cfg.LogFormat, "f", GetEnvWithDefault(logFormatEnvVar, cfg.LogFormat), logFormatHelp)
	fs.StringVar(&cfg.LogFormat, "log-format", GetEnvWithDefault(logFormatEnvVar, cfg.LogFormat), logFormatHelp)

	fs.StringVar(&cfg.Messenger.Urls, "u", GetEnvWithDefault(messagingUrlsEnvVar, cfg.Messenger.Urls), messagingUrlsHelp)
	fs.StringVar(&cfg.Messenger.Urls, "messaging-urls", GetEnvWithDefault(messagingUrlsEnvVar, cfg.Messenger.Urls), messagingUrlsHelp)

	fs.StringVar(&cfg.Messenger.UserCreds, "c", GetEnvWithDefault(messagingUserCredsEnvVar, cfg.Messenger.UserCreds), messagingUserCredsHelp)
	fs.StringVar(&cfg.Messenger.UserCreds, "creds", GetEnvWithDefault(messagingUserCredsEnvVar, cfg.Messenger.UserCreds), messagingUserCredsHelp)
	fs.StringVar(&cfg.Messenger.ClusterID, "messaging-cluster-id", GetEnvWithDefault(messagingClusterIDEnvVar, cfg.Messenger.ClusterID), messagingClusterIDHelp)

	fs.StringVar(&cfg.ConfigFileName, "config", "config.yml", "Config file name")

	fs.StringVar(&cfg.Name, "n", GetEnvWithDefault(appNameEnvVar, appName), appNameHelp)
	fs.StringVar(&cfg.Name, "name", GetEnvWithDefault(appNameEnvVar, appName), appNameHelp)

	fs.StringVar(&cfg.EPNStatusChannel, "epn-status-channel", GetEnvWithDefault(EPNStatusChannelEnvVar, defaultEPNStatusChannel), EPNStatusChannelHelp)

	return fs
}

// GetEnvWithDefault gets the value of the `envVarName` environment variable and return with it.
// If there is no such variable defined in the environment, then return with the `defaultValue`.
func GetEnvWithDefault(envVarName string, defaultValue string) string {
	value, ok := os.LookupEnv(envVarName)
	if !ok {
		value = defaultValue
	}
	return value
}
