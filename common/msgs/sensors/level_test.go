package sensors

import (
	//	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	shouldBeEqual = "The two values should be equal"
)

func TestNewLevelMessage(t *testing.T) {
	//bp := NewLevelMessage(42.)
	//assert.Equal(t, bp.JSON(), (&Level{true}).JSON())
}

/*
func TestLevelGetType(t *testing.T) {
	assert.Equal(t, (&Level{}).GetType(), LevelTypeName)
}

func TestLevelString(t *testing.T) {
	assert.Equal(t, (&Level{}).String(), string(`{"Data":false}`))
	assert.Equal(t, (&Level{false}).String(), string(`{"Data":false}`), shouldBeEqual)
	assert.Equal(t, (&Level{true}).String(), string(`{"Data":true}`), shouldBeEqual)
}

func TestLevelJSON(t *testing.T) {
	assert.Equal(t, (&Level{}).JSON(), []byte(`{"Data":false}`))
	assert.Equal(t, (&Level{false}).JSON(), []byte(`{"Data":false}`), shouldBeEqual)
	assert.Equal(t, (&Level{true}).JSON(), []byte(`{"Data":true}`), shouldBeEqual)
}

func TestLevelParseJSON(t *testing.T) {
	var b Level
	bp := &b

	err := bp.ParseJSON([]byte(`{"Data":false}`))
	assert.Nil(t, err, "Should be nil")
	assert.Equal(t, Level{}, b, shouldBeEqual)
	assert.Equal(t, Level{false}, b, shouldBeEqual)

	err = bp.ParseJSON([]byte(`{"Data":true}`))
	assert.Nil(t, err, "Should be nil")
	assert.Equal(t, Level{true}, b, shouldBeEqual)

	var w Level
	err = (&w).ParseJSON([]byte(`{"Data":"wrong"}`))
	assert.NotNil(t, err, "Should be and error")
	assert.Equal(t, Level{false}, w, shouldBeEqual)
}
*/
