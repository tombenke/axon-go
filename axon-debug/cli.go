package main

import (
	"flag"
	"log"
	"os"

	"github.com/nats-io/nats.go"
)

// axon-debug -u demo.nats.io <subject>
// axon-debug -u demo.nats.io:4443 <subject> (TLS version)
// axon-debug -u demo.nats.io:4222 -s "axon.test.log"


func usage() {
	log.Printf("Usage: axon-debug [-u server] [-creds file] [-s] <subject>\n")
	flag.PrintDefaults()
}

func showUsageAndExit(exitcode int) {
	usage()
	os.Exit(exitcode)
}

type CliParams struct {
    Urls *string
    UserCreds *string
    ShowTime *bool
    ShowHelp *bool
    Subject *string
}

func CliParse() *CliParams {

    parameters := CliParams{
        Urls: flag.String("u", nats.DefaultURL, "The nats server URLs (separated by comma)"),
        UserCreds: flag.String("creds", "", "User Credentials File"),
        Subject: flag.String("s", "axon.log", "The subject to subscribe for inbound messages"),
        ShowTime: flag.Bool("t", false, "Display timestamps"),
        ShowHelp: flag.Bool("h", false, "Show help message"),
    }

	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()

	if *parameters.ShowHelp {
		showUsageAndExit(0)
	}

    return &parameters
}

