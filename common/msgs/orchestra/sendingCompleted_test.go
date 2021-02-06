package orchestra

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/tombenke/axon-go/common/msgs"
	"github.com/tombenke/axon-go/common/msgs/common"
	"testing"
)

func TestSendingCompletedGetType(t *testing.T) {
	assert.Equal(t, NewSendingCompletedMessage("Some text...").GetType(), SendingCompletedTypeName)
}

func TestSendingCompletedMessage(t *testing.T) {
	at := int64(1608732048980057025)
	prec := common.TimePrecision("ns")
	m := NewSendingCompletedMessageAt("Some text...", at, prec)
	var n SendingCompleted
	err := n.ParseJSON(m.JSON())
	assert.Nil(t, err)
	err = n.ParseJSON([]byte(m.String()))
	assert.Nil(t, err)
	assert.Equal(t, m, &n)
}

func TestSendingCompletedMessageCodec(t *testing.T) {
	at := int64(1608732048980057025)
	prec := common.TimePrecision("ns")
	m := NewSendingCompletedMessageAt("Some text...", at, prec)
	var n SendingCompleted
	err := n.Decode(msgs.JSONRepresentation, m.Encode(msgs.JSONRepresentation))
	assert.Nil(t, err)
	assert.Equal(t, m, &n)
}

func TestSendingCompletedMessageCodecPanic(t *testing.T) {
	at := int64(1608732048980057025)
	prec := common.TimePrecision("ns")
	m := NewSendingCompletedMessageAt("Some text...", at, prec)
	var n SendingCompleted
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
