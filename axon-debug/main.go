package main

import (
	"fmt"
	"log"
	"runtime"

	"github.com/nats-io/nats.go"
	axon "github.com/tombenke/axon-go-common"
)

func main() {

	// Parse command line parameters
	parameters := *CliParse()

	// Connect to NATS
	nc, err := axon.ConnectToNats(*parameters.Urls, *parameters.UserCreds, "axon-debug")
	if err != nil {
		log.Fatal(err)
	}

	// Subscribe to the subject to observe and log
	nc.Subscribe(*parameters.Subject, func(msg *nats.Msg) {
		fmt.Printf("%s\n", string(msg.Data))
	})
	nc.Flush()

	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Listening on [%s]", *(parameters.Subject))
	if *parameters.ShowTime {
		log.SetFlags(log.LstdFlags)
	}

	runtime.Goexit()
}
