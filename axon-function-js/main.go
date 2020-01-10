package main

import (
	"log"
	"runtime"
    "io/ioutil"

    "github.com/robertkrimen/otto"
	"github.com/nats-io/nats.go"
    axon "github.com/tombenke/axon-go/common"
)

func main() {

    // Parse command line parameters
    parameters := *CliParse()

    script, err := loadScript(*parameters.ScriptFile)
    axon.CheckFatal(err)

    // Connect to NATS
	nc, err := axon.ConnectToNats(*parameters.Urls, *parameters.UserCreds, "axon-function-js")
    axon.CheckFatal(err)

    // Subscribe to the subject to observe and process
	nc.Subscribe(*parameters.Subject, handleInboundMessage(nc, &parameters, script))
	nc.Flush()
    axon.CheckFatal(nc.LastError())

	log.Printf("Listening on [%s]", *parameters.Subject)

	runtime.Goexit()
}

// Load the JavaScript implementation of the function from the `scriptFileName` file
func loadScript(scriptFileName string) (*string, error) {
    dat, err := ioutil.ReadFile(scriptFileName)
    script := string(dat)
    axon.CheckFatal(err)

    return &script, nil
}

// Receive the incoming message, call the script and sends the result of the script to the target subject
func handleInboundMessage(nc *nats.Conn, parameters *CliParams, script *string) func (*nats.Msg) {

    return func(msg *nats.Msg) {
        outMsg, err := runScript(script, parameters.ScriptParameters, string(msg.Data))
        axon.CheckFatal(err)
        sendOutboundMessage(nc, *parameters.Target, outMsg)
	}
}

// Execute the JavaScript implementation of the function with the incoming message
func runScript(script *string, scriptParameters *string, inMsg string) (string, error) {

    vm := otto.New()
    vm.Set("message", inMsg)
    vm.Set("parameters", scriptParameters)
    vm.Run(*script)
    value, err := vm.Run(*script)
    // err = ReferenceError: abcdefghijlmnopqrstuvwxyz is not defined
    // If there is an error, then value.IsUndefined() is true
    axon.Check("SCRIPT ERROR: ", err)

    return value.String(), nil
}

// Send the result of the function as a message to the target subject
func sendOutboundMessage(nc *nats.Conn, subj string, outMsg string) error {
    msg := []byte(outMsg)

    nc.Publish(subj, msg)
    nc.Flush()
    axon.CheckFatal(nc.LastError())

    log.Printf("Published [%s] : '%s'\n", subj, msg)

    return nil
}

