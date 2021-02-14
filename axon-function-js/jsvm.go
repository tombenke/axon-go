package main

import (
	"encoding/json"
	"fmt"
	"github.com/robertkrimen/otto"
	"github.com/tombenke/axon-go/common/actor/processor"
	"github.com/tombenke/axon-go/common/file"
	"github.com/tombenke/axon-go/common/log"
	"github.com/tombenke/axon-go/common/msgs/base"
)

// JSVM is holds a JavaScript Virtual Machine to execute the processor function written in JavaScript
type JSVM struct {
	vm     *otto.Otto
	script string
}

// NewJSVM creates and returns with a new JavaScript Virtual Machine, and loads the `scriptFile` to run.
func NewJSVM(scriptFile string) *JSVM {
	jsvm := JSVM{
		vm:     otto.New(),
		script: loadScript(scriptFile),
	}
	return &jsvm
}

func loadScript(scriptFile string) string {
	script, errLoadScript := file.LoadFile(scriptFile)
	if errLoadScript != nil {
		panic(fmt.Sprintf("Could not load script-file: '%s'", scriptFile))
	}
	return string(script)
}

// Run runs the preloaded JavaScript script withing the `ctx` context
func (jsvm *JSVM) Run(ctx processor.Context) error {
	// Set the context
	if errVmSet := jsvm.vm.Set("GetInputMessage", func(call otto.FunctionCall) otto.Value {
		fmt.Printf("Called GetInputMessage(%s)\n", call.Argument(0).String())
		inputMsg := ctx.GetInputMessage("input").(*base.Any)
		inputMsgJSON, errJSON := json.Marshal(inputMsg)
		if errJSON != nil {
			panic(fmt.Sprintf("Fatal error in marshalling input message: %s", errJSON.Error()))
		}

		obj, errToObj := jsvm.vm.Object("(" + string(inputMsgJSON) + ")")
		if errToObj != nil {
			panic(fmt.Sprintf("Fatal error in creating JS.Object: %s", errToObj.Error()))
		}

		value, errToValue := jsvm.vm.ToValue(obj)
		if errToValue != nil {
			panic(fmt.Sprintf("Fatal error in converting obj to value: %s", errToValue.Error()))
		}

		return value
	}); errVmSet != nil {
		panic(fmt.Sprintf("Fatal error in set the 'GetInputMessage()' to the JSVM: %s", errVmSet.Error()))
	}

	if errVmSet := jsvm.vm.Set("SetOutputMessage", func(call otto.FunctionCall) otto.Value {
		output := call.Argument(0).String()
		outputMsg, errExport := call.Argument(1).Object().Value().Export()
		if errExport != nil {
			panic(fmt.Sprintf("Fatal error in exporting '%s' output message: %s", output, errExport.Error()))
		}
		fmt.Printf("Called SetOutputMessage(%s, %v)\n", call.Argument(0).String(), outputMsg)

		ctx.SetOutputMessage(output, base.NewAnyMessage(outputMsg.(map[string]interface{})))
		return otto.Value{}
	}); errVmSet != nil {
		panic(fmt.Sprintf("Fatal error in set the 'GetInputMessage()' to the JSVM: %s", errVmSet.Error()))
	}

	// Run the script within the context
	res, errRun := jsvm.vm.Run(jsvm.script)
	if errRun != nil {
		log.Logger.Errorf("JS script run error: %s", errRun.Error())
	}
	log.Logger.Debugf("JS script returned: %v", res)

	return errRun
}
