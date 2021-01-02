// Package outputs provides the functions to forward the results of the processing
package outputs

import (
	"github.com/sirupsen/logrus"
	"github.com/tombenke/axon-go/common/io"
	"github.com/tombenke/axon-go/common/messenger"
	"github.com/tombenke/axon-go/common/msgs"
	"github.com/tombenke/axon-go/common/msgs/orchestra"
	"sync"
)

// Sender receives outputs from the processor function via the `outputsCh` that it sends to
// the corresponding topics identified by the port.
// The outputs structures hold every details about the ports, the message itself, and the subject to send.
// This function runs as a standalone process, so it should be started as a go function.
func Sender(actorName string, outputsCh chan io.Outputs, doneCh chan bool, appWg *sync.WaitGroup, m messenger.Messenger, logger *logrus.Logger) {
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
			sendProcessingCompleted(actorName, m)

		case <-sendResultsCh:
			logger.Infof("Sender received trigger to send outputs")
			sendOutputs(actorName, outputs, m)
		}
	}
}

// sendProcessingCompleted sends a message to the orchestrator about that
// the agent completed the processing and it is ready to send outputs.
func sendProcessingCompleted(actorName string, m messenger.Messenger) {
	logger.Infof("Sender sends 'processing-completed' notification to orchestrator\n")
	processingCompletedMsg := orchestra.NewProcessingCompletedMessage(actorName)
	m.Publish("processing-completed", processingCompletedMsg.Encode(msgs.JSONRepresentation))
}

func sendOutputs(actorName string, outputs io.Outputs, m messenger.Messenger) {
	for o := range outputs {
		channel := outputs[o].Channel
		representation := outputs[o].Representation
		message := outputs[o].Message
		messageType := outputs[o].Type
		logger.Infof("Sender sends '%v' type message of '%s' output port to '%s' channel in '%s' format\n", messageType, o, channel, representation)
		m.Publish(channel, message.Encode(representation))
	}

	logger.Infof("Sender sends 'sending-completed' notification to orchestrator\n")
	sendingCompletedMsg := orchestra.NewSendingCompletedMessage(actorName)
	m.Publish("sending-completed", sendingCompletedMsg.Encode(msgs.JSONRepresentation))
}
