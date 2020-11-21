package main

import (
	"flag"
	"log"
	"os"

	"github.com/nats-io/nats.go"
)

// axon-file-reader -u demo.nats.io <subject>
// axon-file-reader -u demo.nats.io:4443 -s <subject> (TLS version)
// axon-file-reader -u demo.nats.io:4222 -s "axon.25b95691.log"


func usage() {
	log.Printf("Usage: axon-file-reader [-u nats-server] [-creds file] [-s] <subject> [-t] <message-type> [-f] <file-path> [-p] <precision>\n")
	flag.PrintDefaults()
}

func showUsageAndExit(exitcode int) {
	usage()
	os.Exit(exitcode)
}

type CliParams struct {
    Urls *string
    UserCreds *string
    FilePath *string
    MessageType *string
    Subject *string
    Precision *string
    ShowHelp *bool
}

func CliParse() *CliParams {

    parameters := CliParams{
        Urls: flag.String("u", nats.DefaultURL, "The nats server URLs (separated by comma)"),
        UserCreds: flag.String("creds", "", "User Credentials File"),
        FilePath: flag.String("f", "", "The path of the file"),
        MessageType: flag.String("t", "file-reader-output", "The type of the message"),
        Subject: flag.String("s", "axon.file-reader", "The subject to send the outbound messages"),
        Precision: flag.String("p", "ns", "The precision of time value: ns, us, ms, s"),
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

