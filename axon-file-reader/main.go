package main

import (
	"fmt"
    "os"
    "io/ioutil"
	"log"

    axon "github.com/tombenke/axon-go-common"
)

func main() {

    // Parse command line parameters
    parameters := *CliParse()

    // Connect to NATS
	nc, err := axon.ConnectToNats(*parameters.Urls, *parameters.UserCreds, "axon-file-reader")
	if err != nil {
		log.Fatal(err)
	}

    // Read the file content to send
    buf, err := ioutil.ReadFile(*parameters.FilePath)
    if err != nil {
        fmt.Fprintf(os.Stderr, "File read error: %s\n", err)
        panic(err.Error())
    }

    // Create the message to send
    msg := NewMessage(*parameters.MessageType, *parameters.Precision, buf)
    axon.SendMessage(nc, *parameters.Subject, msg)
}

