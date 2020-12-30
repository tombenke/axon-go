// Package outputs provides the functions to forward the results of the processing
package outputs

import (
	"github.com/tombenke/axon-go/common/io"
	//	"github.com/tombenke/axon-go/common/messenger"
	"sync"
)

// Sender receives outputs from the processor function via the `outputsCh` that it sends to
// the corresponding topics identified by the port.
// The outputs structures hold every details about the ports, the message itself, and the subject to send.
// This function runs as a standalone process, so it should be started as a go function.
func Sender(outputsCh chan io.Outputs, done chan bool, appWg *sync.WaitGroup) {
	defer appWg.Done()
	for {
		select {
		case outputs := <-outputsCh:
			sendOutputs(outputs)

		case <-done:
			close(outputsCh)
			return
		}
	}
}

func sendOutputs(outputs io.Outputs) {
	// TODO: Build Context
	//outputs := make(io.Outputs, 1)
}
