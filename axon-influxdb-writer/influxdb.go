package main

import (
	"github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/tombenke/axon-go/common/log"
	"time"
)

// InfluxDb structure represents the client connection to the InfluxDB server.
type InfluxDb struct {
	client   influxdb2.Client
	config   InfluxDbConfig
	writeAPI api.WriteAPI
}

// NewInfluxDbConnection establishes a new connection to the InfluxDB database
func NewInfluxDbConnection(cfg InfluxDbConfig) InfluxDb {

	org := cfg.Organization
	bucket := cfg.Bucket
	client := influxdb2.NewClient(cfg.URL, cfg.Token)
	writeAPI := client.WriteAPI(org, bucket)

	return InfluxDb{client: client, config: cfg, writeAPI: writeAPI}
}

// WritePoint stores the given data point into the bucket of the time-series database
func (i InfluxDb) WritePoint(measurement string, fieldName string, data interface{}) {

	log.Logger.Debugf("WritePoint: %v\n", data)

	// create point
	p := influxdb2.NewPoint(
		measurement,
		map[string]string{
			//"tagName": "tagValue",
		},
		map[string]interface{}{
			fieldName: data,
		},
		time.Now())

	// write asynchronously
	i.writeAPI.WritePoint(p)

	// Force all unwritten data to be sent
	i.writeAPI.Flush()
}

// Close closes the InfluxDb client
func (i InfluxDb) Close() {
	i.client.Close()
}
