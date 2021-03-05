package actorStatus

import (
	"github.com/tombenke/axon-go-common/log"
	"github.com/tombenke/axon-go-common/messenger"
	"github.com/tombenke/axon-go/axon-orchestrator/heartbeat"
	"sync"
)

// Actor represents the actual status of an actor node of the EPN
type Actor struct {
	name     string
	nodeType string
}

// Actors holds the actual status infos of all the actors of the EPN
type Actors map[string]Actor

// Registry represents the Actor Registry process and status registry
type Registry struct {
	actors         Actors
	messenger      messenger.Messenger
	heartbeatCh    chan heartbeat.Heartbeat
	registryDoneCh chan struct{}
	//statusReportCh chan struct{}
}

// NewRegistry creates a new Actor Registry
func NewRegistry(heartbeatCh chan heartbeat.Heartbeat, messenger messenger.Messenger) Registry {

	return Registry{
		actors:         make(Actors, 0),
		messenger:      messenger,
		heartbeatCh:    heartbeatCh,
		registryDoneCh: make(chan struct{}),
	}
}

// Start the Actor Registry
func (r Registry) Start(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		for {
			select {

			case <-r.registryDoneCh:
				wg.Done()
				return

			case hb := <-r.heartbeatCh:
				r.ProcessHeartbeat(hb)

				//			case statusReport := <-r.statusReportCh:
				//				r.ProcessStatusReport(statusReport)
			}
		}
	}()

	log.Logger.Infof("Actor Registry is started")
}

func (r Registry) ProcessHeartbeat(hb heartbeat.Heartbeat) {
	log.Logger.Debugf("Actor Registry processes Heartbeat")
}

// Shutdown stops the Actor Registry process
func (r Registry) Shutdown() {
	log.Logger.Infof("Actor Registry is shutting down")
	close(r.registryDoneCh)
}