package actorStatus

import (
	"github.com/tombenke/axon-go-common/log"
	"github.com/tombenke/axon-go-common/messenger"
	"github.com/tombenke/axon-go-common/msgs"
	"github.com/tombenke/axon-go-common/msgs/orchestra"
	"github.com/tombenke/axon-go/axon-orchestrator/heartbeat"
	"sync"
	"time"
)

// Actor represents the actual status of an actor node of the EPN
type Actor struct {
	Node               orchestra.StatusReportBody
	LastStatusResponse time.Time
	ResponseTime       time.Duration
}

// Actors holds the actual status infos of all the actors of the EPN
type Actors map[string]Actor

// Registry represents the Actor Registry process and status registry
type Registry struct {
	lastStatusRequest        time.Time
	actors                   Actors
	maxResponseTime          time.Duration
	messenger                messenger.Messenger
	heartbeatCh              chan heartbeat.Heartbeat
	statusRequestChannelName string
	epnStatusChannelName     string
	statusReportCh           chan []byte
	statusReportSubs         messenger.Subscriber
	registryDoneCh           chan struct{}
}

// NewRegistry creates a new Actor Registry
func NewRegistry(
	heartbeatCh chan heartbeat.Heartbeat,
	statusRequestChannelName string,
	statusReportChannelName string,
	epnStatusChannelName string,
	maxResponseTime time.Duration,
	messenger messenger.Messenger) Registry {

	statusReportCh := make(chan []byte)
	statusReportSubs := messenger.ChanSubscribe(statusReportChannelName, statusReportCh)

	return Registry{
		lastStatusRequest:        time.Now(),
		actors:                   make(Actors),
		statusRequestChannelName: statusRequestChannelName,
		epnStatusChannelName:     epnStatusChannelName,
		maxResponseTime:          maxResponseTime,
		messenger:                messenger,
		heartbeatCh:              heartbeatCh,
		statusReportCh:           statusReportCh,
		statusReportSubs:         statusReportSubs,
		registryDoneCh:           make(chan struct{}),
	}
}

// Start the Actor Registry
func (r *Registry) Start(wg *sync.WaitGroup) {
	wg.Add(1)
	go func(r *Registry) {
		defer close(r.statusReportCh)

		for {
			select {

			case <-r.registryDoneCh:
				wg.Done()
				return

			case hb := <-r.heartbeatCh:
				r.ProcessHeartbeat(hb)
				r.SendEPNStatus()

			case statusReport := <-r.statusReportCh:
				r.ProcessStatusReport(statusReport)
			}
		}
	}(r)

	log.Logger.Infof("Actor Registry is started")
}

// ProcessHeartbeat method of the Registry processes the heartbeat event
func (r *Registry) ProcessHeartbeat(hb heartbeat.Heartbeat) {
	log.Logger.Debugf("Actor Registry processes Heartbeat")
	for aName, actor := range r.actors {

		timeSinceLastResponse := hb.Timestamp.Sub(actor.LastStatusResponse)
		// If Response time exceeds the limit, then remove the actor from the registry
		if timeSinceLastResponse >= r.maxResponseTime {
			log.Logger.Debugf("Remove '%s' actor from Registry due too long response time", aName)
			delete(r.actors, aName)
		}
	}

	r.lastStatusRequest = hb.Timestamp
	log.Logger.Debugf("Actor Registry sends 'status-request' message")
	statusRequestMsg := orchestra.NewStatusRequestMessageAt(hb.Timestamp.UnixNano(), "ns")
	if err := r.messenger.Publish(r.statusRequestChannelName, statusRequestMsg.Encode(msgs.JSONRepresentation)); err != nil {
		panic(err)
	}
}

// SendEPNStatus sends a status message to the epn-status channel on the overall status of the EPN of actors
func (r *Registry) SendEPNStatus() {
	log.Logger.Debugf("Actor Registry sends EPN Status")
	actors := make([]orchestra.Actor, 0)
	for _, actor := range r.actors {
		actors = append(actors, orchestra.Actor{
			Node:         actor.Node,
			ResponseTime: actor.ResponseTime,
		})
	}
	epnStatusMsg := orchestra.NewEPNStatusMessage(orchestra.EPNStatusBody{
		Actors: actors,
	})
	if err := r.messenger.Publish(r.epnStatusChannelName, epnStatusMsg.Encode(msgs.JSONRepresentation)); err != nil {
		panic(err)
	}
}

// ProcessStatusReport method of the Registry processes the incoming status-report message
func (r *Registry) ProcessStatusReport(statusReportBytes []byte) {
	log.Logger.Debugf("Actor Registry processes status-report: %s", string(statusReportBytes))

	// Determine arrival and response time
	arrivedAt := time.Now()
	responseTime := arrivedAt.Sub(r.lastStatusRequest)

	// Parse the status-report message
	var statusReportMsg orchestra.StatusReport
	if err := (&statusReportMsg).Decode(msgs.JSONRepresentation, statusReportBytes); err != nil {
		panic(err)
	}
	actorName := statusReportMsg.Body.Name

	// Upsert into the registry
	if actor, isInRegistry := r.actors[actorName]; isInRegistry {
		actor.LastStatusResponse = arrivedAt
		actor.ResponseTime = responseTime
		r.actors[actorName] = actor
		log.Logger.Debugf("Update the status of '%s' Actor: %v", actorName, r.actors[actorName])
	} else {
		actor := Actor{
			Node:               statusReportMsg.Body,
			LastStatusResponse: arrivedAt,
			ResponseTime:       responseTime,
		}
		r.actors[actorName] = actor
		log.Logger.Debugf("Put the '%s' Actor into the Registry: %v", actorName, r.actors[actorName])
	}
}

// Shutdown stops the Actor Registry process
func (r *Registry) Shutdown() {
	log.Logger.Infof("Actor Registry is shutting down")
	if err := r.statusReportSubs.Unsubscribe(); err != nil {
		log.Logger.Errorf("Error in unsubscribe from status-report channel: %s", err.Error())
	}
	defer close(r.registryDoneCh)
}
