package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewHeaderAt(t *testing.T) {
	type testCase struct {
		Precision TimePrecision
		Time      int64
	}
	nsValue := int64(1608732048980057025)
	cases := []testCase{
		{"ns", nsValue},
		{"us", nsValue / 1e3},
		{"ms", nsValue / 1e6},
		{"s", nsValue / 1e9},
	}

	for _, c := range cases {
		h := NewHeaderAt(c.Time, c.Precision)
		assert.Equal(t, h, Header{c.Precision, c.Time})
	}
}

func TestNewHeader(t *testing.T) {
	nowNs := time.Now().UnixNano()
	h := NewHeader()
	assert.Equal(t, h.TimePrecision, TimePrecision("ns"))
	assert.Less(t, h.Timestamp-nowNs, int64(1000))
}

func TestNowAsUnixWithPrecision(t *testing.T) {
	type testCase struct {
		Precision TimePrecision
		Factor    int64
	}
	cases := []testCase{
		{"ns", 1e0},
		{"us", 1e3},
		{"ms", 1e6},
		{"s", 1e9},
	}

	for _, c := range cases {
		nowNs := time.Now().UnixNano()
		now := NowAsUnixWithPrecision(c.Precision)
		if c.Precision == TimePrecision("ns") {
			assert.Less(t, now-nowNs, int64(5000))
		} else {
			assert.Equal(t, now, nowNs/c.Factor)
		}
	}
}
