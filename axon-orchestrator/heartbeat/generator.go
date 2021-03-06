package heartbeat

import (
	"github.com/tombenke/axon-go-common/log"
	"sync"
	"time"
)

// Heartbeat represents the actual heartbeat sent by the generator through the heartbeatCh channel
type Heartbeat struct {
	Timestamp time.Time
}

// Generator is the Heartbeat Generator of the Orchestrator application
type Generator struct {
	heartbeat   time.Duration
	cronDoneCh  chan struct{}
	heartbeatCh chan Heartbeat
	ticker      *time.Ticker
}

// NewGenerator creates a new Heartbeat Generator
func NewGenerator(heartbeat time.Duration) (Generator, chan Heartbeat) {

	heartbeatCh := make(chan Heartbeat)
	generator := Generator{
		heartbeat:   heartbeat,
		cronDoneCh:  make(chan struct{}),
		heartbeatCh: heartbeatCh,
		ticker:      time.NewTicker(heartbeat),
	}

	return generator, heartbeatCh
}

// Start the Heartbeat Generator
func (g Generator) Start(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer close(g.heartbeatCh)

		for {
			select {
			case <-g.cronDoneCh:
				wg.Done()
				return

			case <-g.ticker.C:
				g.SendHeartbeat()
			}
		}
	}()

	log.Logger.Infof("Heartbeat Generator is started")
}

func (g Generator) SendHeartbeat() {
	log.Logger.Debugf("Heartbeat Generator sends Heartbeat.")
	hb := Heartbeat{Timestamp: time.Now()}
	g.heartbeatCh <- hb
}

// Shutdown stops the heartbeat generator process
func (g Generator) Shutdown() {
	log.Logger.Infof("Heartbeat Generator is shutting down")
	g.ticker.Stop()
	close(g.cronDoneCh)
}
