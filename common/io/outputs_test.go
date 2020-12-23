package io

import (
	"github.com/tombenke/axon-go/common/msgs/base"
	"testing"
)

func TestOutputsSetOutputMessage(t *testing.T) {
	bmsg := base.NewBoolMessage(true)
	o := Outputs{"State": Output{IO{Name: "State", Type: base.BoolTypeName, Message: bmsg}}}
	(o).SetOutputMessage("State", &base.Bool{true})
	//TODO: Write the test
}
