// Package msgs provides a generic `Message` interface for all the messages can be used by the axon-go actors.
package msgs

// Representation tells the mime-type of the message that is used for encoding and decoding
type Representation string

const (
	// JSONRepresentation Representation enum value
	JSONRepresentation Representation = "application/json"
)

// Codec interface declares the methods that Encodes and Decodes the message to and from `Representation` format.
type Codec interface {
	Encode(Representation) []byte
	Decode(Representation, []byte) error
}

// JSONConverter interface declares the method that Marshals and Unmarshals the message to and from JSON representation.
type JSONConverter interface {
	JSON() []byte
	ParseJSON([]byte) error
}

// YamlConverter interface declares the method that Marshals and Unmarshals the message to and from YAML representation.
type YamlConverter interface {
	YAML() []byte
	ParseYAML([]byte) error
}

// GobConverter interface declares the method that Marshals and Unmarshals the message to and from Golangs GOB representation.
type GobConverter interface {
	EncodeGob() []byte
	DecodeGob([]byte) error
}

// ProtobufConverter interface declares the method that Marshals and Unmarshals the message to and from Protobuf representation.
type ProtobufConverter interface {
}

// ROSConverter interface declares the method that Marshals and Unmarshals the message to and from ROS message format representation.
type ROSConverter interface {
}

// Message is the generic interface of all messages used by the axon-go actors
type Message interface {
	// GetType() returns the Message-Type in string representation
	GetType() string

	// String() Returns the message content as a string in JSON format.
	String() string

	// Codec interface declares the member functions to encode and decode the message to and from a selected representation format
	Codec

	// JSONConverter interface declares the member functions for encoding and decoding the message in JSON representation
	JSONConverter

	//	YamlConverter
	//	GobConverter
	//	ProtobufConverter
	//	ROSConverter
}
