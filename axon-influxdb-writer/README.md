axon-influxdb-writer
====================

The `axon-influxdb-writer` is a subscribe-only agent, that receives messages that contain at least a `time` field, and any number of other fields.
The agent connects to an [`InfluxDB`](https://docs.influxdata.com/influxdb/v1.7/) time-series database, and writes a record into the database, that contains the data properties arrived in the `body`, and uses the `time` property as a time-stamp.

It is configured, what the database name has to be, where the records will be stored.

This is an example that the `axon-cron` agent emits:

```JavaScript
    {
        "type":"measurement",
        "meta":{
            "timePrecision":"ns"
        },
        "body":{
            "time":1578563461000510976,
            "device":"6cfde020-fcd8-493f-9f6c-d8415b4a3fd5",
            "temperature":19.30,
            "humidity":57.40
        }
    }
```

Execute the agent with the `-h` swith to get help:

```bash
$ axon-influxdb-writer -h

Usage: axon-influxdb-writer [-u server] [-creds file] [-s] <subject> [-t]
  -creds string
    	User Credentials File
  -db string
    	The name of the InfluxDB database (default "axon")
  -h	Show help message
  -i string
    	InfluxDB URL (default "http://localhost:8086")
  -icreds string
    	User Credentials File for InfluxDB
  -s string
    	The subject to subscribe for inbound messages (default "axon.log")
  -t	Display timestamps
  -u string
    	The nats server URLs (separated by comma) (default "nats://127.0.0.1:4222")
```
