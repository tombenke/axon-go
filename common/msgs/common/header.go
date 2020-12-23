package common

import (
	"time"
)

type TimePrecisionType string

const (
	Nanoseconds  TimePrecisionType = "ns"
	Microseconds                   = "us"
	Milliseconds                   = "ms"
	Seconds                        = "s"
)

type Header struct {
	/* ROS example
	seq      uint32
	stamp    time
	frame_id string
	*/

	MessageType   string
	TimePrecision TimePrecisionType
	Timestamp     time.Time
}

func nowAsUnixWithPrecision(precision string) int64 {
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
