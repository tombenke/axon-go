package main

import (
	"flag"
	"fmt"
	"github.com/tombenke/axon-go-common/file"
	"github.com/tombenke/axon-go-common/log"
	"gopkg.in/yaml.v2"
	"os"
)

// YAML converts the content of the Config structure to YAML format
func (c Config) YAML() ([]byte, error) {
	return yaml.Marshal(&c)
}

// GetConfig returns with the configuration parameters of the application
// It reads and parses the CLI parameters, loads the external configuration files if needed,
// then merges these parameters. Returns with the resulting configuration set.
func GetConfig(appName string, args []string) Config {

	// Get config file name from CLI parameter, or use the default one
	configFileName := getConfigFileName(appName, defaultConfig, args)
	log.Logger.Debugf("defaultConfig: %v", defaultConfig)

	// Read the configuration from config file, if it is found
	configFileContent, errLoadConfigfile := readConfigFromFile(defaultConfig, configFileName)
	if errLoadConfigfile != nil {
		log.Logger.Warning(errLoadConfigfile)
	}
	log.Logger.Debugf("configFileContent: %v", configFileContent)

	// Parse the CLI config parameters on top of the config-file content
	cliConfigContent := parseCliArgs(configFileContent, appName, args)
	log.Logger.Debugf("cliConfigContent: %v", cliConfigContent)

	// Merges the configurations into a resulting one
	resultingConfig := mergeConfigs(builtInConfig(), cliConfigContent)

	return resultingConfig
}

// getDefaultConfigFileName returns with the path to the config file
func getConfigFileName(appName string, defaultConfig Config, args []string) string {

	fs := GetAppFlagSet(appName, &defaultConfig)
	err := fs.Parse(args)
	if err != nil {
		log.Logger.Warningf(err.Error())
	}

	return defaultConfig.Node.ConfigFileName
}

// parseCliArgs parses the command line arguments and returns with the results as a structure
func parseCliArgs(configFileContent Config, appName string, args []string) Config {

	appConfig := configFileContent

	fs := GetAppFlagSet(appName, &appConfig)

	// Add usage printer function
	fs.Usage = usage(fs, appName)

	err := fs.Parse(args)
	if err != nil {
		log.Logger.Warningf(err.Error())
	}

	// Handle the -h flag
	if appConfig.ShowHelp {
		showUsageAndExit(fs, appName, 0)
	}

	return appConfig
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

// printResultingConfig prints out the actual configuration of the application to the console
func printResultingConfig(config Config) {
	configYAML, _ := config.YAML()
	fmt.Printf("Configuration:\n%s\n", configYAML)
}

// readConfigFromFile reads the config parameters from a file on top of the `defaultConfig`
//
// `defaultConfig` properties will provide the missing values, and the properties from the
// config file will overwrite the default values if they defined.
func readConfigFromFile(defaultConfig Config, path string) (Config, error) {
	c := defaultConfig
	var err error
	content, err := file.LoadFile(path)
	if err == nil {
		err = yaml.Unmarshal([]byte(content), &c)
	}

	return c, err
}
