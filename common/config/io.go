package config

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

// Inputs is an array of the input CLI parameters
type Inputs []In

// Out defines the properties of an output descriptor CLI parameter
type Out struct {
	IO `yaml:",inline"`
}

// Outputs is an array of the output CLI parameters
type Outputs []Out
