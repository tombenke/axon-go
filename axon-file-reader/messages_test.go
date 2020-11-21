package main

import (
    "testing"
    "fmt"
    "encoding/json"
)

var (
    jsonMsg string = `{"type":"file-content","time":1580923185913448282,"meta":{"timePrecision":"ns"},"body":"Hello There!"}`
    binMsg Message = Message{"file-content", 1580923185913448282, Meta{"ns"}, "Hello There!"}
)

func TestMessageMarshal(t *testing.T) {

    marshalled, err := json.Marshal(binMsg)
    if (err != nil) {
        t.Error("JSON marshalling error", err)
    }
    if jsonMsg != string(marshalled) {
        t.Errorf("%s expected to be %s", marshalled, jsonMsg)
    }
    fmt.Printf("%s", marshalled)
}

func TestMessageUnmarshal(t *testing.T) {
    var unmarshalled Message
    err := json.Unmarshal([]byte(jsonMsg), &unmarshalled)
    if err != nil {
        t.Error("JSON unmarshal error", err)
    }
    if unmarshalled != binMsg {
        t.Errorf("%v expected to be %v", unmarshalled, binMsg)
    }
}
