// Package status provides the functions to communicate the status of the actor with the orchestrator application
package status

import (
	"github.com/sirupsen/logrus"
	"github.com/tombenke/axon-go/common/messenger"
	"sync"
)

// Sender receives status request messages from the orchestrator application,
// send responses to these requests, forwarding the actual status of the actor.
// This function runs as a standalone process, so it should be started as a go function.
func Status(actorName string, doneCh chan bool, wg *sync.WaitGroup, m messenger.Messenger, logger *logrus.Logger) {
	statusRequestCh := make(chan []byte)
	statusRequestSubs := m.ChanSubscribe("status-request", statusRequestCh)

	defer statusRequestSubs.Unsubscribe()
	defer close(statusRequestCh)
	defer wg.Done()

	for {
		select {
		case <-doneCh:
			logger.Infof("Status shuts down.")
			return

		case <-statusRequestCh:
			logger.Infof("Status received status-request message")
			logger.Infof("Status sends status-report message")
			m.Publish("status-report", []byte(actorName))
		}
	}
}
