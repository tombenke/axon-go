package epn

import (
	"github.com/tombenke/axon-go-common/log"
	"github.com/tombenke/axon-go-common/messenger"
	"github.com/tombenke/axon-go-common/msgs"
	"github.com/tombenke/axon-go-common/msgs/orchestra"
	"sync"
)

// Status represents the overall status of the EPN
type Status struct {
	epnStatusCh     chan []byte
	epnStatusSubs   messenger.Subscriber
	epnStatusDoneCh chan struct{}
	subscribers     []chan orchestra.EPNStatus
}

// NewStatus creates a new EPN Status Observer instance
func NewStatus(epnStatusChannelName string, messenger messenger.Messenger) Status {
	epnStatusCh := make(chan []byte)
	epnStatusSubs := messenger.ChanSubscribe(epnStatusChannelName, epnStatusCh)

	return Status{
		epnStatusCh:     epnStatusCh,
		epnStatusSubs:   epnStatusSubs,
		epnStatusDoneCh: make(chan struct{}),
		subscribers:     make([]chan orchestra.EPNStatus, 0),
	}
}

// Start the EPN Status Observer
func (s *Status) Start(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer close(s.epnStatusCh)

		for {
			select {

			case <-s.epnStatusDoneCh:
				wg.Done()
				return

			case epnStatus := <-s.epnStatusCh:
				s.ProcessEpnStatus(epnStatus)
			}
		}
	}()

	log.Logger.Debugf("EPN Status Observer is started")
}

func (s *Status) ProcessEpnStatus(epnStatus []byte) {
	log.Logger.Debugf("EPN Status Observer processes epn-status: %s", string(epnStatus))

	// Decode the received status message
	var epnStatusMsg orchestra.EPNStatus
	if err := epnStatusMsg.Decode(msgs.JSONRepresentation, epnStatus); err != nil {
		panic(err)
	}

	// Send the actual status message to each subscribers
	for _, subscriber := range s.subscribers {
		log.Logger.Debugf("EPN Status Observer sends status to subscribers: %v > %v", epnStatusMsg, subscriber)
		subscriber <- epnStatusMsg
	}
}

// Shutdown stops the EPN Status Observer process
func (s *Status) Shutdown() {
	log.Logger.Debugf("EPN Status Observer is shutting down")
	if err := s.epnStatusSubs.Unsubscribe(); err != nil {
		log.Logger.Errorf("Error in unsubscribe from epn-status channel: %s", err.Error())
	}
	defer close(s.epnStatusDoneCh)
}

// Subscribe returns a channel through which the Status Observer will send status changes
func (s *Status) Subscribe() chan orchestra.EPNStatus {
	subscriber := make(chan orchestra.EPNStatus)
	(*s).subscribers = append(s.subscribers, subscriber)

	log.Logger.Debugf("EPNStatus subscribers: %v", s.subscribers)
	return subscriber
}
