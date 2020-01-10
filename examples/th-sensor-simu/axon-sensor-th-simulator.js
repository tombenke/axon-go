// The inbound message is placed into the 'message' variable as a string
// Parse the JSON format string of the input message
input = JSON.parse(message)
device = "default-device-id"
if (parameters !== "") {
    scriptParameters = JSON.parse(parameters)
    if (scriptParameters["device"]) {
        device = scriptParameters.device
    }
}

// Create random values it it were measured
temperature = 20. + Math.random() * 2
humidity = 55. + Math.random() * 3


// Creates the output message object
output = {
    device: device,
    time: input.time,
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

