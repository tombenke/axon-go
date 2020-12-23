package config

import (
	"strings"
)

// IO defines the properties of a generic I/O port
type IO struct {
	Topic string
	Name  string
}

// In defines the properties of an input descriptor CLI parameter
type In struct {
	IO
	DefaultValue interface{}
}

// Inputs is an array of the input CLI parameters
type Inputs []In

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
	parts := strings.Split(inStr, ":")

	switch len(parts) {
	case 1:
		result = In{IO: IO{Topic: parts[0], Name: parts[0]}, DefaultValue: ""}
	case 2:
		result = In{IO: IO{Topic: parts[0], Name: parts[0]}, DefaultValue: parts[1]}
	case 3:
		if parts[0] == "" {
			result = In{IO: IO{Topic: parts[1], Name: parts[1]}, DefaultValue: parts[2]}
		} else {
			result = In{IO: IO{Topic: parts[0], Name: parts[1]}, DefaultValue: parts[2]}
		}
	}

	if result.Name == "" {
		panic("Input name must be defined!")
	}
	return result
}

// Out defines the properties of an output descriptor CLI parameter
type Out struct {
	IO
}

// Outputs is an array of the output CLI parameters
type Outputs []Out

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
	parts := strings.Split(inStr, ":")

	switch len(parts) {
	case 1:
		result = Out{IO: IO{Topic: parts[0], Name: parts[0]}}
	case 2:
		if parts[1] == "" {
			result = Out{IO: IO{Topic: parts[0], Name: parts[0]}}
		} else {
			result = Out{IO: IO{Topic: parts[1], Name: parts[0]}}
		}
	}

	if result.Name == "" {
		panic("Input name must be defined!")
	}
	return result
}
