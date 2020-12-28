package config

import (
	"strings"
)

// String is a dummy implementation of the function
func (i *Inputs) String() string {
	return ""
}

// Set appends a new In CLI parameter to the inputs array
func (i *Inputs) Set(value string) error {
	*i = append(*i, parseIn(value))
	return nil
}

// parseIn parses the input CLI parameter and returns with an `In` object build from the parse results
func parseIn(inStr string) (result In) {
	parts := strings.Split(inStr, "|")

	switch len(parts) {
	case 1:
		result = In{IO: IO{Name: parts[0], Channel: "", Type: DefaultType, Representation: DefaultRepresentation}, Default: ""}
	case 2:
		result = In{IO: IO{Name: parts[0], Channel: parts[1], Type: DefaultType, Representation: DefaultRepresentation}, Default: ""}
	case 5:
		result = In{IO: IO{Name: parts[0], Channel: parts[1], Type: parts[2], Representation: parts[3]}, Default: parts[4]}
	default:
		panic("Wrong number of input port parameters")
	}

	if result.Name == "" {
		panic("Input port name must be defined!")
	}

	if result.Type == "" {
		result.Type = DefaultType
	}

	if result.Representation == "" {
		result.Representation = DefaultRepresentation
	}

	return result
}

// String is a dummy implementation of the function
func (o *Outputs) String() string {
	return ""
}

// Set appends a new out CLI parameter to the outputs array
func (o *Outputs) Set(value string) error {
	*o = append(*o, parseOut(value))
	return nil
}

// parseOut parses the output CLI parameter and returns with an `Out` object build from the parse results
func parseOut(inStr string) (result Out) {
	parts := strings.Split(inStr, "|")

	switch len(parts) {
	case 1:
		result = Out{IO: IO{Name: parts[0], Channel: "", Type: DefaultType, Representation: DefaultRepresentation}}
	case 2:
		result = Out{IO: IO{Name: parts[0], Channel: parts[1], Type: DefaultType, Representation: DefaultRepresentation}}
	case 4:
		result = Out{IO: IO{Name: parts[0], Channel: parts[1], Type: parts[2], Representation: parts[3]}}
	default:
		panic("Wrong number of output port parameters")
	}

	if result.Name == "" {
		panic("Output port name must be defined!")
	}

	if result.Type == "" {
		result.Type = DefaultType
	}

	if result.Representation == "" {
		result.Representation = DefaultRepresentation
	}

	return result
}
