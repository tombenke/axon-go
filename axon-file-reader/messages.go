package main

import (
	"encoding/json"
	axon "github.com/tombenke/axon-go-common"
)

// Message decribe the structure of a generic message
type Message struct {
	Type string `json:"type"`
	Time int64  `json:"time"`
	Meta Meta   `json:"meta"`
	Body string `json:"body"`
}

// Meta structure describes the Meta informations inside the Message
type Meta struct {
	TimePrecision string `json:"timePrecision"`
}

// NewMessage creates a new message with the timestamp of the actual time, and return the new message in JSON representation
func NewMessage(messageType string, timePrecision string, buf []byte) []byte {
	timestamp := axon.NowAsUnixWithPrecision(timePrecision)
	msg, _ := json.Marshal(Message{messageType, timestamp, Meta{timePrecision}, string(buf)})
	return []byte(msg)
}
