package common

import (
	"time"
)

// TimePrecision is the type of precision of the Timestamp value hold by the message Header
type TimePrecision string

const (
	// Nanoseconds TimePrecision enum value
	Nanoseconds TimePrecision = "ns"
	// Microseconds TimePrecision enum value
	Microseconds TimePrecision = "us"
	// Milliseconds TimePrecision enum value
	Milliseconds TimePrecision = "ms"
	// Seconds TimePrecision enum value
	Seconds TimePrecision = "s"
	// DefaultTimePrecision is the default value for Timeprecision
	DefaultTimePrecision TimePrecision = Nanoseconds
)

// Header is the generic message header structure
type Header struct {
	TimePrecision TimePrecision
	Timestamp     int64
}

// NewHeader creates a new generic message header with the timestamp of the current time in "ns" precision
func NewHeader() Header {
	return NewHeaderAt(NowAsUnixWithPrecision(DefaultTimePrecision), DefaultTimePrecision)
}

// NewHeaderAt creates a new generic message header with the `at` timestamp value and `withPrecision` precision
func NewHeaderAt(at int64, withPrecision TimePrecision) Header {
	return Header{
		TimePrecision: withPrecision,
		Timestamp:     at,
	}
}

// NowAsUnixWithPrecision returns with the current time with `precision` precision
func NowAsUnixWithPrecision(precision TimePrecision) int64 {
	nowNs := time.Now().UnixNano()
	switch precision {
	case "ns":
		return nowNs
	case "u", "us":
		return nowNs / 1e3
	case "ms":
		return nowNs / 1e6
	case "s":
		return nowNs / 1e9
	}
	return nowNs
}
