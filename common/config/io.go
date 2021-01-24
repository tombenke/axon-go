package config

import (
	"fmt"
)

const (
	// DefaultType is the default message-type for IO ports
	DefaultType = "base/Any"
	// DefaultRepresentation the default representation for IO ports
	DefaultRepresentation = "application/json"
)

// IO defines the properties of a generic I/O port
type IO struct {
	Name           string
	Type           string
	Representation string
	Channel        string
}

// In defines the properties of an input descriptor CLI parameter
type In struct {
	IO      `yaml:",inline"`
	Default string
}

// WouldModify returns true if the modifiable properties of the `in` input
// differs from the corresponding properties of `mod`, otherwise returns false.
func (in In) WouldModify(mod In) bool {
	if in.Representation == mod.Representation &&
		in.Channel == mod.Channel &&
		in.Default == mod.Default {

		return false
	}

	return true
}

// ModifyWith replaces the configurable properties of the `in` input with the properties of `mod`
func (in *In) ModifyWith(mod In) {
	(*in).Representation = mod.Representation
	(*in).Channel = mod.Channel
	(*in).Default = mod.Default
}

// Inputs is an array of the input CLI parameters
type Inputs []In

func (inputs Inputs) FindByName(name string) (*In, bool) {
	for i := range inputs {
		if inputs[i].Name == name {
			return &(inputs[i]), true
		}
	}
	return nil, false
}

func (inputs *Inputs) ExtendWith(ext Inputs) {
	for e := range ext {
		if _, found := inputs.FindByName(ext[e].Name); !found {
			fmt.Println("before", *inputs)
			*inputs = append(*inputs, ext[e])
			fmt.Println("after", *inputs)
		}
	}
}

func (inputs Inputs) ModifyWith(mod Inputs) {
	for m := range mod {
		if i, found := inputs.FindByName(mod[m].Name); found {
			fmt.Println("before", inputs)
			i.ModifyWith(mod[m])
			fmt.Println("after", inputs)
		}
	}
}

// Out defines the properties of an output descriptor CLI parameter
type Out struct {
	IO `yaml:",inline"`
}

// WouldModify returns true if the modifiable properties of the `out` output
// differs from the corresponding properties of `mod`, otherwise returns false.
func (out Out) WouldModify(mod Out) bool {
	if out.Representation == mod.Representation &&
		out.Channel == mod.Channel {

		return false
	}

	return true
}

// ModifyWith replaces the configurable properties of the `out` input with the properties of `mod`
func (out *Out) ModifyWith(mod Out) {
	(*out).Representation = mod.Representation
	(*out).Channel = mod.Channel
}

// Outputs is an array of the output CLI parameters
type Outputs []Out

func (outputs Outputs) FindByName(name string) (*Out, bool) {
	for o := range outputs {
		if outputs[o].Name == name {
			return &(outputs[o]), true
		}
	}
	return nil, false
}

func (outputs *Outputs) ExtendWith(ext Outputs) {
	for e := range ext {
		if _, found := outputs.FindByName(ext[e].Name); !found {
			fmt.Println("before", *outputs)
			*outputs = append(*outputs, ext[e])
			fmt.Println("after", *outputs)
		}
	}
}

func (outputs Outputs) ModifyWith(mod Outputs) {
	for m := range mod {
		if o, found := outputs.FindByName(mod[m].Name); found {
			fmt.Println("before", outputs)
			o.ModifyWith(mod[m])
			fmt.Println("after", outputs)
		}
	}
}
