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

	// Show the help of the application
	ShowHelp bool `yaml:"showHelp"`

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
	configFileName := getConfigFileName(appName, defaultConfig, args)

	// Read the configuration from config file, if it is found
	configFileContent, _ := readConfigFromFile(defaultConfig, configFileName)

	// Parse the CLI config parameters on top of the config-file content
	cliConfigContent := parseCliArgs(configFileContent, appName, args)

	// Merges the configurations into a resulting one
	resultingConfig := mergeConfigs(hardCodedConfigContent, cliConfigContent)

	return resultingConfig
}

// getDefaultConfigFileName returns with the path to the config file
func getConfigFileName(appName string, defaultConfig Config, args []string) string {

	fs := GetAppFlagSet(appName, &defaultConfig)
	fs.Parse(args)

	return defaultConfig.Node.ConfigFileName
}

// parseCliArgs parses the command line arguments and returns with the results as a structure
func parseCliArgs(configFileContent Config, appName string, args []string) Config {

	appConfig := configFileContent

	fs := GetAppFlagSet(appName, &appConfig)

	// Add usage printer function
	fs.Usage = usage(fs, appName)

	fs.Parse(args)

	// Handle the -h flag
	if appConfig.ShowHelp {
		showUsageAndExit(fs, appName, 0)
	}

	return appConfig
}

// GetAppFlagSet returns with the flag-set of the application to parse the CLI parameters
func GetAppFlagSet(appName string, cfg *Config) *flag.FlagSet {
	fs := config.GetDefaultFlagSet(appName, &cfg.Node)

	fs.BoolVar(&cfg.ShowHelp, "h", false, "Show help message")
	fs.BoolVar(&cfg.ShowHelp, "help", false, "Show help message")

	fs.BoolVar(&cfg.PrintConfig, "p", false, "Print configuration parameters")
	fs.BoolVar(&cfg.PrintConfig, "print-config", false, "Print configuration parameters")

	//TODO Add additional CLI flags if needed here

	return fs
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

// mergeConfigs returns with the resulting config parameters set after merging them
//
// `hardCodedConfigContent` holds the configuration that the application defined as a baseline,
// `cliConfigContent` holds those configuration parameters, that origins from the default values,
// then extended by the config file, if there is any, then finally these were extended by the
// parameters from the environment and the CLI arguments.
// The most complex task of merging the I/O ports are done by the `config.MergeNodeConfigs()` function
// according to the values of `Extend` and `Modify` flags defined by the `hardCodedConfigContent`.
// The application needs to implement the merging of properties added by itself, on top of the Node parameters.
func mergeConfigs(hardCodedConfigContent Config, cliConfigContent Config) Config {
	resultingConfig := hardCodedConfigContent

	// Let the internal config module to manage the merging of the Node parameters
	resultingNode, err := config.MergeNodeConfigs(hardCodedConfigContent.Node, cliConfigContent.Node)
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}
	resultingConfig.Node = resultingNode

	//TODO: Add application-level merging tasks here if there is any
	resultingConfig.ShowHelp = cliConfigContent.ShowHelp
	resultingConfig.PrintConfig = cliConfigContent.PrintConfig

	return resultingConfig
}
