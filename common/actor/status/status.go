// Package status provides the functions to communicate the status of the actor with the orchestrator application
package status

import (
	"github.com/sirupsen/logrus"
	"github.com/tombenke/axon-go/common/messenger"
	"github.com/tombenke/axon-go/common/msgs"
	"github.com/tombenke/axon-go/common/msgs/orchestra"
	"sync"
)

// Status receives status request messages from the orchestrator application,
// sends responses to these requests, forwarding the actual status of the actor.
// This function runs as a standalone process, so it should be started as a go function.
func Status(actorName string, doneCh chan bool, wg *sync.WaitGroup, m messenger.Messenger, logger *logrus.Logger) chan bool {
	statusRequestCh := make(chan []byte)
	statusRequestSubs := m.ChanSubscribe("status-request", statusRequestCh)
	statusStoppedCh := make(chan bool)

	wg.Add(1)
	go func() {
		defer func() {
			if err := statusRequestSubs.Unsubscribe(); err != nil {
				panic(err)
			}
			close(statusRequestCh)
			wg.Done()

			logger.Infof("Status stopped.")
			close(statusStoppedCh)
		}()

		for {
			select {
			case <-doneCh:
				logger.Infof("Status shuts down.")
				return

			case <-statusRequestCh:
				logger.Infof("Status received status-request message")
				logger.Infof("Status sends status-report message")
				statusReportMsg := orchestra.NewStatusReportMessage(actorName)
				if err := m.Publish("status-report", statusReportMsg.Encode(msgs.JSONRepresentation)); err != nil {
					panic(err)
				}
				// TODO: Make orchestra message representations and channel names configurable
			}
		}
	}()
	logger.Infof("Status started")
	return statusStoppedCh
}
