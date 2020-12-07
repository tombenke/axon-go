package config

import (
	"strings"
)

// Generic IO properties
type IO struct {
	Topic string
	Name  string
}

// Inputs
type In struct {
	IO
	DefaultValue interface{}
}

type Inputs []In

func (i *Inputs) String() string {
	return ""
}

func (i *Inputs) Set(value string) error {
	*i = append(*i, parseIn(value))
	return nil
}

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

// Outputs
type Out struct {
	IO
}

type Outputs []Out

func (o *Outputs) String() string {
	return ""
}

func (o *Outputs) Set(value string) error {
	*o = append(*o, parseOut(value))
	return nil
}

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
