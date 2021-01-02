package orchestra

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/tombenke/axon-go/common/msgs"
	"github.com/tombenke/axon-go/common/msgs/common"
	"testing"
)

func TestSendResultsGetType(t *testing.T) {
	assert.Equal(t, NewSendResultsMessage().GetType(), SendResultsTypeName)
}

func TestSendResultsMessage(t *testing.T) {
	at := int64(1608732048980057025)
	prec := common.TimePrecision("ns")
	m := NewSendResultsMessageAt(at, prec)
	var n SendResults
	n.ParseJSON(m.JSON())
	n.ParseJSON([]byte(m.String()))
	assert.Equal(t, m, &n)
}

func TestSendResultsMessageCodec(t *testing.T) {
	at := int64(1608732048980057025)
	prec := common.TimePrecision("ns")
	m := NewSendResultsMessageAt(at, prec)
	var n SendResults
	n.Decode(msgs.JSONRepresentation, m.Encode(msgs.JSONRepresentation))
	assert.Equal(t, m, &n)
}

func TestSendResultsMessageCodecPanic(t *testing.T) {
	at := int64(1608732048980057025)
	prec := common.TimePrecision("ns")
	m := NewSendResultsMessageAt(at, prec)
	var n SendResults
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
