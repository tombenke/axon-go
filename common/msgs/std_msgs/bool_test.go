package std_msgs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	shouldBeEqual = "The two values should be equal"
)

func TestNewBoolMessage(t *testing.T) {
	bp := NewBoolMessage(true)
	assert.Equal(t, bp.JSON(), (&Bool{true}).JSON())
}

func TestBoolGetType(t *testing.T) {
	assert.Equal(t, (&Bool{}).GetType(), BoolTypeName)
}

func TestBoolString(t *testing.T) {
	assert.Equal(t, (&Bool{}).String(), string(`{"Data":false}`))
	assert.Equal(t, (&Bool{false}).String(), string(`{"Data":false}`), shouldBeEqual)
	assert.Equal(t, (&Bool{true}).String(), string(`{"Data":true}`), shouldBeEqual)
}

func TestBoolJSON(t *testing.T) {
	assert.Equal(t, (&Bool{}).JSON(), []byte(`{"Data":false}`))
	assert.Equal(t, (&Bool{false}).JSON(), []byte(`{"Data":false}`), shouldBeEqual)
	assert.Equal(t, (&Bool{true}).JSON(), []byte(`{"Data":true}`), shouldBeEqual)
}

func TestBoolParseJSON(t *testing.T) {
	var b Bool
	bp := &b

	err := bp.ParseJSON([]byte(`{"Data":false}`))
	assert.Nil(t, err, "Should be nil")
	assert.Equal(t, Bool{}, b, shouldBeEqual)
	assert.Equal(t, Bool{false}, b, shouldBeEqual)

	err = bp.ParseJSON([]byte(`{"Data":true}`))
	assert.Nil(t, err, "Should be nil")
	assert.Equal(t, Bool{true}, b, shouldBeEqual)

	var w Bool
	err = (&w).ParseJSON([]byte(`{"Data":"wrong"}`))
	assert.NotNil(t, err, "Should be and error")
	assert.Equal(t, Bool{false}, w, shouldBeEqual)
}
