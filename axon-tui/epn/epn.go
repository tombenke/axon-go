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
}

// NewStatus creates a new EPN Status Observer instance
func NewStatus(epnStatusChannelName string, messenger messenger.Messenger) Status {
	epnStatusCh := make(chan []byte)
	epnStatusSubs := messenger.ChanSubscribe(epnStatusChannelName, epnStatusCh)

	return Status{
		epnStatusCh:     epnStatusCh,
		epnStatusSubs:   epnStatusSubs,
		epnStatusDoneCh: make(chan struct{}),
	}
}

// Start the EPN Status Observer
func (s Status) Start(wg *sync.WaitGroup) {
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

	log.Logger.Infof("EPN Status Observer is started")
}

func (s Status) ProcessEpnStatus(epnStatus []byte) {
	log.Logger.Debugf("EPN Status Observer processes epn-status: %s", string(epnStatus))
	var epnStatusMsg orchestra.EPNStatus
	if err := epnStatusMsg.Decode(msgs.JSONRepresentation, epnStatus); err != nil {
		panic(err)
	}

	// TODO: Implement visualization
}

// Shutdown stops the EPN Status Observer process
func (s Status) Shutdown() {
	log.Logger.Infof("EPN Status Observer is shutting down")
	if err := s.epnStatusSubs.Unsubscribe(); err != nil {
		log.Logger.Errorf("Error in unsubscribe from epn-status channel: %s", err.Error())
	}
	defer close(s.epnStatusDoneCh)
}
