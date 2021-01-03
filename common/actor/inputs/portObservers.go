package inputs

import (
	"github.com/sirupsen/logrus"
	"github.com/tombenke/axon-go/common/io"
	"github.com/tombenke/axon-go/common/messenger"
	"sync"
)

// startInPortsObservers starts one message observer for every port,
// and returns with the number of observers started.
func startInPortsObservers(inputs io.Inputs, inputsMuxCh chan io.Input, doneCh chan bool, wg *sync.WaitGroup, m messenger.Messenger, logger *logrus.Logger) {
	for p := range inputs {
		newPortObserver(inputs[p], inputsMuxCh, doneCh, wg, m, logger)
	}
}

// newPortObserver subscribes to an input channel with a go routine that observes the incoming messages.
// When a message arrives through the channel, the go routine forwards that through the `inCh` towards the aggregator.
// The newPortObserver creates and returns with the `inCh` channel that the aggregator can consume.
func newPortObserver(input io.Input, inputsMuxCh chan io.Input, doneCh chan bool, wg *sync.WaitGroup, m messenger.Messenger, logger *logrus.Logger) {
	inMsgCh := make(chan []byte)
	inMsgSubs := m.ChanSubscribe(input.Channel, inMsgCh)

	wg.Add(1)
	go func() {
		logger.Infof("Receiver's '%s' port observer started", input.Name)
		defer logger.Infof("Receiver's '%s' port observer stopped", input.Name)
		defer inMsgSubs.Unsubscribe()
		defer close(inMsgCh)
		defer wg.Done()

		for {
			select {
			case <-doneCh:
				logger.Infof("Receiver's '%s' port observer shut down", input.Name)
				return

			case inputMsg := <-inMsgCh:
				logger.Infof("Receiver's '%s' port observer received message", input.Name)
				//logger.Infof("Input message observer received a message via '%s' channel for '%s' port", input.Channel, input.Name)
				input.Message.Decode(input.Representation, inputMsg)
				inputsMuxCh <- input
				logger.Infof("Receiver's '%s' port observer sent message to inputMuxCh channel", input.Name)
			}
		}
	}()

	return
}
