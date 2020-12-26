package base

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/tombenke/axon-go/common/msgs"
	"testing"
)

func TestAnyGetType(t *testing.T) {
	data := new(map[string]interface{})
	assert.Equal(t, NewAnyMessage(*data).GetType(), AnyTypeName)
}

func TestAnyMessage(t *testing.T) {
	data := new(map[string]interface{})
	m := NewAnyMessage(*data)
	var n Any
	n.ParseJSON(m.JSON())
	n.ParseJSON([]byte(m.String()))
	assert.Equal(t, m, &n)
}

func TestAnyMessageCodec(t *testing.T) {
	data := new(map[string]interface{})
	m := NewAnyMessage(*data)
	var n Any
	n.Decode(msgs.JSONRepresentation, m.Encode(msgs.JSONRepresentation))
	assert.Equal(t, m, &n)
}

func TestAnyMessageCodecPanic(t *testing.T) {
	data := new(map[string]interface{})
	m := NewAnyMessage(*data)
	var n Any
	func() {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, r, errors.New("Decode error: unknown representational format 'wrong-representation'"))
			}
		}()
		n.Decode(msgs.Representation("wrong-representation"), m.Encode(msgs.JSONRepresentation))
	}()
	func() {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, r, errors.New("Encode error: unknown representational format 'wrong-representation'"))
			}
		}()
		n.Decode(msgs.JSONRepresentation, m.Encode(msgs.Representation("wrong-representation")))
	}()
}

func TestComplexJSONToAnyMessage(t *testing.T) {
	m := NewAnyMessage(Any{"Meta": map[string]interface{}{"TimePrecision": "ns"}, "Time": 1.608732048980057e+18, "Type": "heartbeat"})
	jsonText := []byte(`{"Meta": {"TimePrecision": "ns"}, "Time": 1608732048980057025, "Type": "heartbeat"}`)
	var n Any
	n.Decode(msgs.JSONRepresentation, jsonText)
	fmt.Printf("%v\n", n)
	assert.Equal(t, m, &n)
}
