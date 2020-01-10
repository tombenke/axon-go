axon-function-js
================

The `axon-function-js` is a consumer / producer agent.
It can receive any type of message.
It executes a script written in JavaScript.
This script is exeuted every time a message arrives from the subject the agent is subscribed for.
The script gets the incoming message injected into a variable, named `message`.
The script also can get parameters, using the `scriptparams` agrument of the command. This is a string, that the script can get as a variable, named `parameters` and the script implementation decides how to evaluate it.

Inside the script the standard JavaScript built-in functions, variables and constants are usable, including `console.log()`.
The value of the last expression of the script will be the message that the agent will forward into the `target` subject.

The [`examples/th-sensor-simu/`](../examples/th-sensor-simu/) folder shows an example for the usage of this agent.

This `axon-function-js` agent, that loads the [`examples/th-sensor-simu/axon-sensor-th-simu.js` script](../examples/th-sensor-simu/axon-sensor-th-simulator.js) below, that expects a message, that contains a timestamp value in the `body` of the incoming message.

The JSON example below demonstrates such a message:

```JavaScript
    {
        "meta": {
            "timePrecision":"ns"
        },
        body: {
            "time": 1578655455000164640
        }
    }
```

This is the script, that will process the incoming message:

```JavaScript
    // The inbound message is placed into the 'message' variable as a string
    // Parse the JSON format string of the input message
    input = JSON.parse(message)

    // Create random values it it were measured
    temperature = 20. + Math.random() * 2
    humidity = 55. + Math.random() * 3

    // Creates the output message object
    output = {
        device: "143a0c9d-291a-4077-8fcc-d2aa259b8de2",
        time: input.body.time,
        meta: input.meta,
        body: {
            humidity: humidity,
            temperature: temperature
        }
    }

    // Serialize the JSON object to string
    result = JSON.stringify(output)

    // The last expression will be the outbound message as a string
    result
```

Then the following example below shows how the result, that will be send to the target, looks like:

```JavaScript
    {
        "body": {
            "device": "143a0c9d-291a-4077-8fcc-d2aa259b8de2",
            "humidity": 56.683517562784935,
            "temperature": 20.6242920550331,
            "time":1578654765000150272
        },
        "meta": {
            "timePrecision": "ns"
        }
    }
```

Execute the agent with the `-h` swith to get help:

```bash
$ go install && axon-function-js -h
Usage: axon-function-js [-u server] [-creds file] [-s] <source-subject> [-t] <target-subject> -scriptfile <script-filename> -scriptparams <script-parameters>
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
