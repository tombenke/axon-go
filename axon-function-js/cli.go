package main

import (
	"flag"
	"log"
	"os"

	"github.com/nats-io/nats.go"
)

// axon-function-js -u demo.nats.io <subject>
// axon-function-js -u demo.nats.io:4443 <subject> (TLS version)
// axon-function-js -u demo.nats.io:4222 -s "axon.test.log"


func usage() {
	log.Printf("Usage: axon-function-js [-u server] [-creds file] [-s] <source-subject> [-t] <target-subject>\n")
	flag.PrintDefaults()
}

func showUsageAndExit(exitcode int) {
	usage()
	os.Exit(exitcode)
}

type CliParams struct {
    Urls *string
    UserCreds *string
    ShowHelp *bool
    Subject *string
    Target *string
    ScriptFile *string
}

func CliParse() *CliParams {

    parameters := CliParams{
        Urls: flag.String("u", nats.DefaultURL, "The nats server URLs (separated by comma)"),
        UserCreds: flag.String("creds", "", "User Credentials File"),
        Subject: flag.String("s", "axon.func.in", "The subject to subscribe for inbound messages"),
        Target: flag.String("t", "axon.func.out", "The subject to send the outbound messages into"),
        ScriptFile: flag.String("scriptfile", "function.js", "The name of the JavaScript file that holds the function implementation."),
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

