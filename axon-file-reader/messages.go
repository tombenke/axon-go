package main

import (
    "encoding/json"
    axon "github.com/tombenke/axon-go-common"
)

type Message struct {
    Type string `json:"type"`
    Time int64 `json:"time"`
    Meta Meta `json:"meta"`
    Body string `json:"body"`
}

type Meta struct {
    TimePrecision string `json:"timePrecision"`
}

func NewMessage(messageType string, timePrecision string, buf []byte) []byte {
    timestamp := axon.NowAsUnixWithPrecision(timePrecision)
    msg, _ := json.Marshal(Message{messageType, timestamp, Meta{timePrecision}, string(buf)})
    return []byte(msg)
}
