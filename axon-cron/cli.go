package main

import (
	"flag"
	"log"
	"os"

	"github.com/nats-io/nats.go"
)

// axon-cron -u demo.nats.io <subject>
// axon-cron -u demo.nats.io:4443 -s <subject> (TLS version)
// axon-cron -u demo.nats.io:4222 -s "axon.25b95691.log"

func usage() {
	log.Printf("Usage: axon-cron [-u nats-server] [-creds file] [-s] <subject> [-t] <message-type> [-cron] <cron definition> [-p] <precision>\n")
	flag.PrintDefaults()
}

func showUsageAndExit(exitcode int) {
	usage()
	os.Exit(exitcode)
}

// CliParams holds the configuration parameters of the application
type CliParams struct {
	Urls        *string
	UserCreds   *string
	CronDef     *string
	MessageType *string
	Subject     *string
	Precision   *string
	ShowHelp    *bool
}

// CliParse parses the CLI parameters and returns with the results
func CliParse() *CliParams {

	parameters := CliParams{
		Urls:        flag.String("u", nats.DefaultURL, "The nats server URLs (separated by comma)"),
		UserCreds:   flag.String("creds", "", "User Credentials File"),
		CronDef:     flag.String("cron", "@every 10s", "Cron definition"),
		MessageType: flag.String("t", "", "The type of the message"),
		Subject:     flag.String("s", "axon.cron", "The subject to send the outbound messages"),
		Precision:   flag.String("p", "ns", "The precision of time value: ns, us, ms, s"),
		ShowHelp:    flag.Bool("h", false, "Show help message"),
	}

	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()

	if *parameters.ShowHelp {
		showUsageAndExit(0)
	}

	return &parameters
}
