package app

import (
	"flag"
	"fmt"
	config "github.com/tombenke/axon-go/common/config"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

// Config holds the configuration parameters of the actor node application
// It must inherit the configuration of the core `Node` object,
// and may contain other additional application-specific parameters.
type Config struct {
	// Node holds the config parameters that every actor node needs
	Node config.Node `yaml:"node"`

	// PrintConfig if true, then prints the resulting configuration to the console
	PrintConfig bool `yaml:"printConfig"`

	// TODO: Add the additional config parameters of the applications
}

// YAML converts the content of the Config structure to YAML format
func (c Config) YAML() ([]byte, error) {
	return yaml.Marshal(&c)
}

// GetConfig returns with the configuration parameters of the application
// It reads and parses the CLI parameters, loads the external configuration files if needed,
// then merges these parameters. Returns with the resulting configuration set.
func GetConfig(appName string, hardCodedConfigContent Config, args []string) Config {
	defaultConfig := Config{
		Node:        config.GetDefaultNode(),
		PrintConfig: false,
	}

	// Get config file name from CLI parameter, or use the default one
	configFileName := getConfigFileName(args)

	// Read the configuration from config file, if it is found
	configFileContent, _ := readConfigFromFile(defaultConfig, configFileName)

	// Parse the CLI config parameters on top of the config-file content
	cliConfigContent := parseCliArgs(configFileContent, appName, args)

	// Merges the content of the three sources of configuration parameters into one
	resultingConfig := mergeConfigs(hardCodedConfigContent, cliConfigContent, configFileContent)

	return resultingConfig
}

// getDefaultConfigFileName returns with the path to the config file
func getConfigFileName(args []string) string {
	// Check the CLI parameter
	// TODO

	var configFileName string
	fs := flag.NewFlagSet("find-configfile-name", flag.ContinueOnError)
	fs.StringVar(&configFileName, "config", "config.yml", "Config file name")
	fs.Parse(args)

	return configFileName
}

// parseCliArgs parses the command line arguments and returns with the results as a structure
func parseCliArgs(configFileContent Config, appName string, args []string) Config {

	appConfig := configFileContent

	// Get the dafault CLI args
	fs := config.GetDefaultFlagSet(appName, &appConfig.Node)

	// Extend the default set with node specific arguments
	var showHelp bool
	fs.BoolVar(&showHelp, "h", false, "Show help message")
	fs.BoolVar(&showHelp, "help", false, "Show help message")

	fs.BoolVar(&appConfig.PrintConfig, "p", false, "Print configuration parameters")
	fs.BoolVar(&appConfig.PrintConfig, "print-config", false, "Print configuration parameters")

	//TODO Add additional CLI flags if needed here

	// Add usage printer function
	fs.Usage = usage(fs, appName)

	fs.Parse(args)

	// Handle the -h flag
	if showHelp {
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
	content, err := LoadFile(path)
	if err == nil {
		err = yaml.Unmarshal([]byte(content), &c)
	}
	// TODO: Write warning about config not found
	return c, err
}

// LoadFile loads []byte content from a file
func LoadFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return []byte(""), err
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return []byte(""), err
	}

	return content, nil
}

// mergeConfigs returns with the resulting config parameters set after merging the coming from the three sources
func mergeConfigs(hardCodedConfigContent Config, cliConfigContent Config, configFileContent Config) Config {
	resultingConfig := hardCodedConfigContent

	//TODO: Implement
	resultingConfig.PrintConfig = cliConfigContent.PrintConfig
	resultingConfig.Node = config.MergeNodeConfigs(hardCodedConfigContent.Node, cliConfigContent.Node, configFileContent.Node)

	//fmt.Println("mergeConfig:", hardCodedConfigContent, cliConfigContent, configFileContent, "=>", resultingConfig)
	return resultingConfig
}
