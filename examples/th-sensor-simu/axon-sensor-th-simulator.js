input = GetInputMessage("input")
//console.log("\ninput: ", JSON.stringify(input))

max = GetInputMessage("max")

// Create random values if it were measured
temperature = Math.random() * max.value

// Creates the output message object
output = {
    Header: input.Header,
    Body: {
        Data: temperature
    }
}

//console.log("\noutput: ", output.time)
SetOutputMessage("output", output)
