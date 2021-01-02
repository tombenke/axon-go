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
	n.ParseJSON(m.JSON())
	n.ParseJSON([]byte(m.String()))
	assert.Equal(t, m, &n)
}

func TestReceiveAndProcessMessageCodec(t *testing.T) {
	at := int64(1608732048980057025)
	prec := common.TimePrecision("ns")
	m := NewReceiveAndProcessMessageAt(42, at, prec)
	var n ReceiveAndProcess
	n.Decode(msgs.JSONRepresentation, m.Encode(msgs.JSONRepresentation))
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

func TestParseDefaultJSONValue(t *testing.T) {
	var m ReceiveAndProcess
	m.ParseJSON([]byte(`{"Body": { "Data": 42 }}`))
	assert.Equal(t, float64(42), m.Body.Data)
}
