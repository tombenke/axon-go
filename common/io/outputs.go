package io

import (
	"errors"
	"fmt"
	"github.com/tombenke/axon-go/common/msgs"
	"reflect"
)

type Output struct {
	IO
}

type Outputs map[string]Output

type OutputsHandler interface {
	SetOutputMessage(string, msgs.Message) error
}

func (outputs *Outputs) SetOutputMessage(name string, outMsg msgs.Message) error {
	if _, ok := (*outputs)[name]; ok {
		fmt.Printf("outputs: %s", reflect.TypeOf((*outputs)[name]))
		(*outputs)[name] = Output{IO{Name: name, Type: outMsg.GetType(), Message: outMsg}}
		return nil
	}
	return errors.New("There is no output port named to " + name)
}
