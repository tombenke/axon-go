package main

import (
    "fmt"
	"log"
	"runtime"
    "time"
    "encoding/json"

	"github.com/nats-io/nats.go"
    axon "github.com/tombenke/axon-go/common"
	influx "github.com/influxdata/influxdb1-client/v2"
)

func connectToInfluxDb(url string, creds string) influx.Client {
	c, err := influx.NewHTTPClient(influx.HTTPConfig{
		Addr: url,
	})
	if err != nil {
		fmt.Println("Error creating InfluxDB Client: ", err.Error())
	}
	defer c.Close()

    return c
}

func logToInfluxDb(client influx.Client, influxDbName string, payload string) {
    fmt.Printf("\n%s\n", payload)

	jsonMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(payload), &jsonMap)
	if err != nil {
		panic(err)
	}

    // Create a new point batch
	bp, _ := influx.NewBatchPoints(influx.BatchPointsConfig{
		Database:  influxDbName,
		Precision: "s",
	})

	// Create a point and add to batch
	tags := map[string]string{"room": "bedroom"}
    body := jsonMap["body"].(map[string]interface{})
    /*
	fields := map[string]interface{}{
		"temperature": body["temperature"],
		"humidity": body["humidity"],
	}
    */
	pt, err := influx.NewPoint(influxDbName, tags, body/*fields*/, time.Unix(0, int64(jsonMap["time"].(float64))))
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}
	bp.AddPoint(pt)
    fmt.Print("pt: ", pt)

	// Write the batch
	client.Write(bp)
}

func main() {
    // Parse command line parameters
    parameters := *CliParse()

    // Connect to InfluxDB
    c := connectToInfluxDb(*parameters.InfluxDbUrl, *parameters.InfluxDbCreds)

    // Connect to NATS
	nc, err := axon.ConnectToNats(*parameters.Urls, *parameters.UserCreds, "axon-influxdb-writer")
	if err != nil {
		log.Fatal(err)
	}

    // Subscribe to the subject to observe and log
	nc.Subscribe(*parameters.Subject, func(msg *nats.Msg) {
        logToInfluxDb(c, *parameters.InfluxDbName, string(msg.Data))
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

