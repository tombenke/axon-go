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
// and may contain other additiona application-specific parameters.
type Config struct {
	// Node holds the config parameters that every actor node needs
	Node config.Node `yaml:"node"`

	// TODO: Add the additional config parameters of the applications
}

// YAML converts the content of the Config structure to YAML format
func (c Config) YAML() ([]byte, error) {
	return yaml.Marshal(&c)
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

	appConfig := Config{}

	// Get the dafault CLI args
	fs := config.GetDefaultFlagSet(appName, &appConfig.Node)

	// Extend the default set with node specific arguments
	var showHelp bool
	fs.BoolVar(&showHelp, "h", false, "Show help message")
	fs.BoolVar(&showHelp, "help", false, "Show help message")
	//TODO Add additional CLI flags fs.StringVar(&appConfig.Precision, "p", "ns", "The precision of time value: ns, us, ms, s")

	// Add usage printer function
	fs.Usage = usage(fs, appName)

	fs.Parse(args)
	fmt.Printf("CONFIG: %v\n%v\n", appConfig, showHelp)

	// Handle the -h flag
	if showHelp {
		showUsageAndExit(fs, appName, 0)
	}

	// TODO: Move to config
	printConfig := true
	if printConfig {
		PrintConfig(appConfig)
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

// PrintConfig prints out the actual configuration of the application to the console
func PrintConfig(config Config) {
	configYAML, _ := config.YAML()
	fmt.Printf("Configuration:\n%s\n", configYAML)
}

// ReadConfigFromFile Reads the the config parameters from a file
func ReadConfigFromFile(path string) (Config, error) {
	c := Config{}
	var err error
	content, err := LoadFile(path)
	if err == nil {
		err = yaml.Unmarshal([]byte(content), &c)
	}
	return c, err
}

// LoadFile loads []byte content from a file
func LoadFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	return content, nil
}
