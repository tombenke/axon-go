// Package outputs provides the functions to forward the results of the processing
package outputs

import (
	"github.com/sirupsen/logrus"
	"github.com/tombenke/axon-go/common/io"
	"github.com/tombenke/axon-go/common/messenger"
	"sync"
)

// AsyncSender receives outputs from the processor function via the `outputsCh` that it sends to
// the corresponding topics identified by the port.
// The outputs structures hold every details about the ports, the message itself, and the subject to send.
// This function runs as a standalone process, so it should be started as a go function.
func AsyncSender(actorName string, outputsCh chan io.Outputs, doneCh chan bool, wg *sync.WaitGroup, m messenger.Messenger, logger *logrus.Logger) chan bool {
	var outputs io.Outputs
	senderStoppedCh := make(chan bool)

	wg.Add(1)
	go func() {
		defer logger.Infof("Sender stopped")
		defer close(senderStoppedCh)
		defer wg.Done()

		for {
			select {
			case <-doneCh:
				logger.Infof("Sender shuts down.")
				return

			case outputs = <-outputsCh:
				logger.Infof("Sender received outputs")
				// In async mode it immediately sends the outputs whet it gets them
				asyncSendOutputs(actorName, outputs, m)
			}
		}
	}()

	logger.Infof("Sender started in async mode.")
	return senderStoppedCh
}

func asyncSendOutputs(actorName string, outputs io.Outputs, m messenger.Messenger) {
	for o := range outputs {
		channel := outputs[o].Channel
		representation := outputs[o].Representation
		message := outputs[o].Message
		messageType := outputs[o].Type
		logger.Infof("Sender sends '%v' type message of '%s' output port to '%s' channel in '%s' format\n", messageType, o, channel, representation)
		m.Publish(channel, message.Encode(representation))
	}
}
