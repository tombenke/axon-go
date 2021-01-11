package app

import (
	"flag"
	"fmt"
	commonConfig "github.com/tombenke/axon-go/common/config"
	"os"
)

type Config struct {
	// Node holds the config parameters that every actor node needs
	Node commonConfig.Node

	// TODO: Add the additional config parameters of the applications
}

// GetConfig returns with the configuration parameters of the application
// It reads and parses the CLI parameters, loads the external configuration files if needed,
// then merges these parameters. Returns with the resulting configuration set.
func GetConfig(appName string, args []string) Config {
	cliParams := parseCliArgs(appName, args)
	// TODO: configFileParams := parseConfigFile(configPath)
	// TODO: resultingConfig := mergeConfigs(predefinedParams, cliParams, fileParams)

	return cliParams
}

// parseCliArgs parses the command line arguments and returns with the results as a structure
func parseCliArgs(appName string, args []string) Config {

	config := Config{}

	// Get the dafault CLI args
	fs := commonConfig.GetDefaultFlagSet(appName, &config.Node)

	// Extend the default set with node specific arguments
	var showHelp bool
	fs.BoolVar(&showHelp, "h", false, "Show help message")
	fs.BoolVar(&showHelp, "help", false, "Show help message")
	//TODO Add additional CLI flags fs.StringVar(&config.Precision, "p", "ns", "The precision of time value: ns, us, ms, s")

	// Add usage printer function
	fs.Usage = usage(fs, appName)

	fs.Parse(args)
	fmt.Printf("CONFIG: %v\n%v\n", config, showHelp)

	// Handle the -h flag
	if showHelp {
		showUsageAndExit(fs, appName, 0)
	}

	return config
}

// Show usage info then exit
func showUsageAndExit(fs *flag.FlagSet, appName string, exitcode int) {
	usage(fs, appName)()
	os.Exit(exitcode)
}

// Print usage information
func usage(fs *flag.FlagSet, appName string) func() {
	return func() {
		fmt.Println("Usage: " + appName + " -h\n")
		fs.PrintDefaults()
	}
}
