console.log("Hello axon-function-js!")
input = GetInputMessage("input")
console.log("input: ", JSON.stringify(input))
output = input
SetOutputMessage("output", output)
