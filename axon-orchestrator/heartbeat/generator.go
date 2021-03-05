package heartbeat

import (
	"github.com/tombenke/axon-go-common/log"
	"github.com/tombenke/axon-go-common/messenger"
	"github.com/tombenke/axon-go-common/msgs"
	"github.com/tombenke/axon-go-common/msgs/orchestra"
	"sync"
	"time"
)

// Heartbeat represents the actual heartbeat sent by the generator through the heartbeatCh channel
type Heartbeat struct{}

// Generator is the Heartbeat Generator of the Orchestrator application
type Generator struct {
	messenger                messenger.Messenger
	statusRequestChannelName string
	heartbeat                time.Duration
	cronDoneCh               chan struct{}
	heartbeatCh              chan Heartbeat
	ticker                   *time.Ticker
}

// NewGenerator creates a new Heartbeat Generator
func NewGenerator(heartbeat time.Duration, channelName string, messenger messenger.Messenger) (Generator, chan Heartbeat) {

	heartbeatCh := make(chan Heartbeat)
	generator := Generator{
		messenger:                messenger,
		statusRequestChannelName: channelName,
		heartbeat:                heartbeat,
		cronDoneCh:               make(chan struct{}),
		heartbeatCh:              heartbeatCh,
		ticker:                   time.NewTicker(heartbeat),
	}

	return generator, heartbeatCh
}

// Start the Heartbeat Generator
func (g Generator) Start(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		for {
			select {
			case <-g.cronDoneCh:
				wg.Done()
				return

			case <-g.ticker.C:
				g.SendStatusRequest()
			}
		}
	}()

	log.Logger.Infof("Heartbeat Generator is started")
}

func (g Generator) SendStatusRequest() {
	log.Logger.Debugf("Heartbeat Generator sends 'status-request' message.")
	var hb Heartbeat
	g.heartbeatCh <- hb
	statusRequestMsg := orchestra.NewStatusRequestMessage()
	if err := g.messenger.Publish(g.statusRequestChannelName, statusRequestMsg.Encode(msgs.JSONRepresentation)); err != nil {
		panic(err)
	}
}

// Shutdown stops the heartbeat generator process
func (g Generator) Shutdown() {
	log.Logger.Infof("Heartbeat Generator is shutting down")
	g.ticker.Stop()
	close(g.cronDoneCh)
	close(g.heartbeatCh)
}
