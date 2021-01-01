// Package outputs provides the functions to forward the results of the processing
package outputs

import (
	"github.com/sirupsen/logrus"
	"github.com/tombenke/axon-go/common/io"
	"github.com/tombenke/axon-go/common/messenger"
	"sync"
)

// Sender receives outputs from the processor function via the `outputsCh` that it sends to
// the corresponding topics identified by the port.
// The outputs structures hold every details about the ports, the message itself, and the subject to send.
// This function runs as a standalone process, so it should be started as a go function.
func Sender(outputsCh chan io.Outputs, doneCh chan bool, appWg *sync.WaitGroup, m messenger.Messenger, logger *logrus.Logger) {
	logger.Infof("Sender started.")
	sendResultsCh := make(chan []byte)
	sendResultsSubs := m.ChanSubscribe("send-results", sendResultsCh)
	var outputs io.Outputs

	defer sendResultsSubs.Unsubscribe()
	defer close(sendResultsCh)
	defer appWg.Done()

	for {
		select {
		case <-doneCh:
			logger.Infof("Sender shuts down.")
			return

		case outputs = <-outputsCh:
			logger.Infof("Sender received outputs")
			sendProcessingCompleted(m)

		case <-sendResultsCh:
			logger.Infof("Sender received trigger to send outputs")
			sendOutputs(outputs, m)
		}
	}
}

// sendProcessingCompleted sends a message to the orchestrator about that
// the agent completed the processing and it is ready to send outputs.
func sendProcessingCompleted(m messenger.Messenger) {
	// TODO: Add agent ID to te message
	logger.Infof("Sender sends 'processing-completed' notification to orchestrator\n")
	m.Publish("processing-completed", []byte("process-completed-msg"))
}

func sendOutputs(outputs io.Outputs, m messenger.Messenger) {
	for o := range outputs {
		channel := outputs[o].Channel
		representation := outputs[o].Representation
		message := outputs[o].Message
		messageType := outputs[o].Type
		logger.Infof("Sender sends '%v' type message of '%s' output port to '%s' channel in '%s' format\n", messageType, o, channel, representation)
		m.Publish(channel, message.Encode(representation))
	}
	// TODO: Add agent ID to te message
	logger.Infof("Sender sends 'sending-completed' notification to orchestrator\n")
	m.Publish("sending-completed", []byte("sending-completed-msg"))
}
