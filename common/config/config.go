package config

import (
	"flag"
	"github.com/tombenke/axon-go/common/messenger"
	"os"
)

type IO struct {
	Topic string
	Name  string
}

type In struct {
	IO
	DefaultValue interface{}
}

type Out struct {
	IO
}

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
	actorName = "axon-cron"

	logLevelHelp    = "The log level: panic | fatal | error | warning | info | debug | trace"
	logLevelEnvVar  = "LOG_LEVEL"
	defaultLogLevel = "info"

	logFormatHelp    = "The log format: json | text"
	logFormatEnvVar  = "LOG_FORMAT"
	defaultLogFormat = "json"

	natsUrlsEnvVar = "NATS_URL"
	natsUrlsHelp   = "The NATS server's URLs (separated by comma)"

	natsUserCredsHelp    = "User Credentials"
	natsUserCredsEnvVar  = "NATS_CREDENTIALS"
	defaultNatsUserCreds = ""
)

// The default config struct that every axon actor node inherits
type Config struct {
	messenger.MessengerConfig

	LogLevel  string
	LogFormat string
	//Inputs            []string
	//Outputs           []string
}

// Parse the command line arguments and returns with the results as a structure
func GetConfig() Config {

	config := Config{}

	flag.StringVar(&config.LogLevel, "l", GetEnvWithDefault(logLevelEnvVar, defaultLogLevel), logLevelHelp)
	flag.StringVar(&config.LogLevel, "log-level", GetEnvWithDefault(logLevelEnvVar, defaultLogLevel), logLevelHelp)

	flag.StringVar(&config.LogFormat, "f", GetEnvWithDefault(logFormatEnvVar, defaultLogFormat), logFormatHelp)
	flag.StringVar(&config.LogFormat, "log-format", GetEnvWithDefault(logFormatEnvVar, defaultLogFormat), logFormatHelp)

	flag.StringVar(&config.Urls, "u", GetEnvWithDefault(natsUrlsEnvVar, messenger.DefaultNatsURL()), natsUrlsHelp)
	flag.StringVar(&config.Urls, "nats-urls", GetEnvWithDefault(natsUrlsEnvVar, messenger.DefaultNatsURL()), natsUrlsHelp)

	flag.StringVar(&config.UserCreds, "c", GetEnvWithDefault(natsUserCredsEnvVar, defaultNatsUserCreds), natsUserCredsHelp)
	flag.StringVar(&config.UserCreds, "creds", GetEnvWithDefault(natsUserCredsEnvVar, defaultNatsUserCreds), natsUserCredsHelp)
	var showHelp bool
	flag.BoolVar(&showHelp, "h", false, "Show help message")

	/*
		flag.Usage = usage
		flag.Parse()

		if showHelp {
			showUsageAndExit(0)
		}
	*/
	return config
}
