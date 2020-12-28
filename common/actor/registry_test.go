package actor

import (
	"github.com/stretchr/testify/assert"
	"github.com/tombenke/axon-go/common/msgs"
	"github.com/tombenke/axon-go/common/msgs/base"
	"testing"
)

// Test if registry has registered entries
func TestRegistry(t *testing.T) {
	assert.True(t, msgs.IsMessageTypeRegistered(base.AnyTypeName))
	assert.True(t, msgs.DoesMessageTypeImplementsRepresentation(base.AnyTypeName, msgs.JSONRepresentation))
}

// Try to re-register the `Any` message-type
func TestRegistryReRegisterAnyType(t *testing.T) {
	assert.Panics(t,
		func() {
			msgs.RegisterMessageType(base.AnyTypeName, []msgs.Representation{msgs.JSONRepresentation}, func() msgs.Message {
				return base.NewAnyMessage(map[string]interface{}{})
			})
		},
		"It should panic",
	)
}
