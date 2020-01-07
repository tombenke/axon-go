package main

import (
	"fmt"
	"log"
    "time"

    "github.com/robfig/cron"
    axon "github.com/tombenke/axon-go/common"
)

func nowAsUnixWithPrecision(precision string) int64 {
    nowNs := time.Now().UnixNano()
    switch precision {
    case "ns":
        return nowNs
    case "u", "us":
        return nowNs / 1e3
    case "ms":
        return nowNs / 1e6
    case "s":
        return nowNs / 1e9
    }
    return nowNs
}


func main() {

    // Parse command line parameters
    parameters := *CliParse()

    // Connect to NATS
	nc, err := axon.ConnectToNats(*parameters.Urls, *parameters.UserCreds, "axon-cron")
	if err != nil {
		log.Fatal(err)
	}

    c := cron.New()
    c.AddFunc(*parameters.CronDef, func() {
        timestamp := nowAsUnixWithPrecision(*parameters.Precision)
        subj := *parameters.Subject
        var msg []byte
        if *parameters.MessageType != "" {
            msg = []byte(fmt.Sprintf("{\"time\":%d,\"type\":\"%s\",\"meta\":{\"timePrecision\":\"%s\"}}", timestamp, *parameters.MessageType, *parameters.Precision))
        } else {
            msg = []byte(fmt.Sprintf("{\"time\":%d,\"meta\":{\"timePrecision\":\"%s\"}}", timestamp, *parameters.Precision))
        }
        nc.Publish(subj, msg)
        nc.Flush()

        if err := nc.LastError(); err != nil {
            log.Fatal(err)
        } else {
            log.Printf("Published [%s] : '%s'\n", subj, msg)
        }
    })

    c.Start()
    defer c.Stop()  // Stop the scheduler (does not stop any jobs already running).

    for {
    }
}
