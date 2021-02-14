console.log("Hello axon-function-js!")
input = GetInputMessage("input")
console.log("input: ", input.time, input.meta.timePrecision)
output = input
SetOutputMessage("output", output)
