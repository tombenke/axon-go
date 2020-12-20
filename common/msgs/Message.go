package msgs

// ProtobufConverter interface declares the method that Marshals and Unmarshals the message to and from JSON representation.
type JSONConverter interface {
	JSON() []byte
	ParseJSON([]byte) error
}

// ProtobufConverter interface declares the method that Marshals and Unmarshals the message to and from YAML representation.
type YamlConverter interface {
	YAML() []byte
	ParseYAML([]byte) error
}

// ProtobufConverter interface declares the method that Marshals and Unmarshals the message to and from Golangs GOB representation.
type GobConverter interface {
	EncodeGob() []byte
	DecodeGob([]byte) error
}

// ProtobufConverter interface declares the method that Marshals and Unmarshals the message to and from Protobuf representation.
type ProtobufConverter interface {
}

// ProtobufConverter interface declares the method that Marshals and Unmarshals the message to and from ROS message format representation.
type ROSConverter interface {
}

// Message is the generic interface of all messages used by the axon-go actors
type Message interface {
	// String() Returns the message content as a string in JSON format.
	String() string
	// GetType() returns the Message-Type in string representation
	GetType() string

	JSONConverter
	//	YamlConverter
	//	GobConverter
	//	ProtobufConverter
	//	ROSConverter
}
