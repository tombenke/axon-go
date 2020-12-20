package io

import (
	"github.com/tombenke/axon-go/common/msgs/std_msgs"
	"testing"
)

func TestOutputsSetOutputMessage(t *testing.T) {
	bmsg := std_msgs.NewBoolMessage(true)
	o := Outputs{"State": Output{IO{Name: "State", Type: std_msgs.BoolTypeName, Message: bmsg}}}
	(o).SetOutputMessage("State", &std_msgs.Bool{true})
	//TODO: Write the test
}
