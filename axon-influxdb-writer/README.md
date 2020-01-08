axon-influxdb-writer
====================

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
