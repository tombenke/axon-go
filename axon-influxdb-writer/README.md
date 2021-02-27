# axon-influxdb-writer

The `axon-influxdb-writer` is a consumer agent.
It can receive any type of message that holds a numeric or boolean value.
It saves the value into an InfluxDB time-series database.

## Usage

Execute the agent with the `-h` switch to get help:

```bash
$ go install && axon-influxdb-writer -h
Usage: axon-influxdb-writer [-u server] [-creds file] [-s] <source-subject> [-t] <target-subject> -scriptfile <script-filename> -scriptparams <script-parameters>
  -creds string
    	User Credentials File
  -h	Show help message
  -s string
    	The subject to subscribe for inbound messages (default "axon.func.in")
  -scriptfile string
    	The name of the JavaScript file that holds the function implementation. (default "function.js")
  -scriptparams string
    	THe parameters of the script.
  -t string
    	The subject to send the outbound messages into (default "axon.func.out")
  -u string
    	The nats server URLs (separated by comma) (default "nats://127.0.0.1:4222")
```

