input = GetInputMessage("input")
// console.log("\ninput: ", input.time, input.meta.timePrecision)

device = GetInputMessage("device")
// console.log("\ndevice: ", JSON.stringify(device))

// Create random values if it were measured
temperature = 20. + Math.random() * 2
humidity = 55. + Math.random() * 3

// Creates the output message object
output = {
    device: device.id,
    time: input.time,
    meta: input.meta,
    body: {
        humidity: humidity,
        temperature: temperature
    }
}

//console.log("\noutput: ", output.time)
SetOutputMessage("output", output)
