package orchestra

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/tombenke/axon-go/common/msgs"
	"github.com/tombenke/axon-go/common/msgs/common"
	"testing"
)

func TestReceiveAndProcessGetType(t *testing.T) {
	assert.Equal(t, NewReceiveAndProcessMessage(42).GetType(), ReceiveAndProcessTypeName)
}

func TestReceiveAndProcessMessage(t *testing.T) {
	at := int64(1608732048980057025)
	prec := common.TimePrecision("ns")
	m := NewReceiveAndProcessMessageAt(42, at, prec)
	var n ReceiveAndProcess
	err := n.ParseJSON(m.JSON())
	assert.Nil(t, err)
	err = n.ParseJSON([]byte(m.String()))
	assert.Nil(t, err)
	assert.Equal(t, m, &n)
}

func TestReceiveAndProcessMessageCodec(t *testing.T) {
	at := int64(1608732048980057025)
	prec := common.TimePrecision("ns")
	m := NewReceiveAndProcessMessageAt(42, at, prec)
	var n ReceiveAndProcess
	err := n.Decode(msgs.JSONRepresentation, m.Encode(msgs.JSONRepresentation))
	assert.Nil(t, err)
	assert.Equal(t, m, &n)
}

func TestReceiveAndProcessMessageCodecPanic(t *testing.T) {
	at := int64(1608732048980057025)
	prec := common.TimePrecision("ns")
	m := NewReceiveAndProcessMessageAt(42, at, prec)
	var n ReceiveAndProcess
	func() {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, r, errors.New("Decode error: unknown representational format 'wrong-representation'"))
			}
		}()
		err := n.Decode(msgs.Representation("wrong-representation"), m.Encode(msgs.JSONRepresentation))
		assert.Nil(t, err)
	}()
	func() {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, r, errors.New("Encode error: unknown representational format 'wrong-representation'"))
			}
		}()
		err := n.Decode(msgs.JSONRepresentation, m.Encode(msgs.Representation("wrong-representation")))
		assert.Nil(t, err)
	}()
}

func TestParseDefaultJSONValue(t *testing.T) {
	var m ReceiveAndProcess
	err := m.ParseJSON([]byte(`{"Body": { "Data": 42 }}`))
	assert.Nil(t, err)
	assert.Equal(t, float64(42), m.Body.Data)
}
