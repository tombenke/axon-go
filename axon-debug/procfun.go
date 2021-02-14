package main

import (
	"encoding/json"
	"fmt"
	"github.com/tombenke/axon-go/common/actor/processor"
	"github.com/tombenke/axon-go/common/msgs/base"
	"gopkg.in/yaml.v2"
	"strings"
)

var (
	validDebugFormats = map[string]interface{}{
		"json":        nil,
		"json-indent": nil,
		"yaml":        nil,
		"yml":         nil,
	}
)

// getProcessorFun is the message processor function of the actor node
func getProcessorFun(config Config) func(ctx processor.Context) error {
	format := strings.ToLower(config.DebugFormat)
	if _, hasKey := validDebugFormats[format]; !hasKey {
		if format != "" {
			panic(fmt.Sprintf("Wrong debug-format value: '%s'", format))
		}
	}

	return func(ctx processor.Context) error {
		msg := ctx.GetInputMessage("input").(*base.Any)
		var msgStr []byte
		var err error
		switch format {
		case "yaml":
			fallthrough
		case "yml":
			msgStr, err = yaml.Marshal(msg)
			fmt.Printf("---\n%s", msgStr)
		case "json":
			msgStr, err = json.Marshal(msg)
			fmt.Printf("%s\n", msgStr)
		case "json-indent":
			fallthrough
		default:
			msgStr, err = json.MarshalIndent(msg, "", "  ")
			fmt.Printf("\n%s\n", msgStr)
		}
		if err != nil {
			panic(err)
		}

		return nil
	}
}
