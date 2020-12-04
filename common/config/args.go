package config

import (
	"strings"
)

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
