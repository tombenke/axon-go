package main

import (
	"flag"
	"log"
	"os"

	"github.com/nats-io/nats.go"
)

// axon-influxdb-writer -u demo.nats.io <subject>
// axon-influxdb-writer -u demo.nats.io:4443 <subject> (TLS version)
// axon-influxdb-writer -u demo.nats.io:4222 -s "axon.25b95691.log"

func usage() {
	log.Printf("Usage: axon-influxdb-writer [-u server] [-creds file] [-s] <subject> [-t]\n")
	flag.PrintDefaults()
}

func showUsageAndExit(exitcode int) {
	usage()
	os.Exit(exitcode)
}

// CliParams holds the configuration parameters of the application
type CliParams struct {
	Urls          *string
	UserCreds     *string
	InfluxDbURL   *string
	InfluxDbCreds *string
	InfluxDbName  *string
	ShowTime      *bool
	ShowHelp      *bool
	Subject       *string
}

// CliParse parses the CLI parameters and returns with the results
func CliParse() *CliParams {

	parameters := CliParams{
		Urls:          flag.String("u", nats.DefaultURL, "The nats server URLs (separated by comma)"),
		UserCreds:     flag.String("creds", "", "User Credentials"),
		InfluxDbURL:   flag.String("i", "http://localhost:8086", "InfluxDB URL"),
		InfluxDbCreds: flag.String("icreds", "", "User Credentials File for InfluxDB"),
		InfluxDbName:  flag.String("db", "axon", "The name of the InfluxDB database"),
		Subject:       flag.String("s", "axon.log", "The subject to subscribe for inbound messages"),
		ShowTime:      flag.Bool("t", false, "Display timestamps"),
		ShowHelp:      flag.Bool("h", false, "Show help message"),
	}

	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()

	if *parameters.ShowHelp {
		showUsageAndExit(0)
	}

	return &parameters
}
