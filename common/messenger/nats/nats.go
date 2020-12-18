package nats

import (
	nats "github.com/nats-io/nats.go"
	"time"
)

// Publish `msg` message to the `subject` topic
func (m connections) Publish(subject string, msg []byte) error {
	subj := subject

	m.nc.Publish(subj, msg)
	m.nc.Flush()

	err := m.nc.LastError()
	if err != nil {
		m.logger.Error(err)
	}

	m.logger.Infof("Published message to '%s'\n", subj)
	return err
}

// Subscribe to the `subject` topic, and calls the `cb` call-back function with the inbound messages
func (m connections) Subscribe(subject string, cb func([]byte)) {
	m.nc.Subscribe(subject, func(msg *nats.Msg) {
		m.logger.Debugf("Received message from '%s'\n", subject)
		cb(msg.Data)
	})
	m.nc.Flush()
}

// Request `msg` message through the `subject` topic and expects a response until `timeout`.
func (m connections) Request(subject string, msg []byte, timeout time.Duration) ([]byte, error) {
	subj := subject

	m.logger.Infof("Request '%s' through '%s'\n", msg, subj)
	resp, err := m.nc.Request(subj, msg, timeout)
	if err != nil {
		m.logger.Error(err)
		return nil, err
	}

	m.logger.Infof("Response '%s'\n", resp.Data)
	return resp.Data, err
}

// Subscribe to the `subject` topic, and calls the `service` call-back function with the inbound messages,
// then respond with the return value of the `service` function through the `Reply` subject.
func (m connections) Response(subject string, service func([]byte) ([]byte, error)) {
	m.nc.Subscribe(subject, func(msg *nats.Msg) {
		resp, err := service(msg.Data)
		if err != nil {
			m.nc.Publish(msg.Reply, []byte(err.Error()))
		} else {
			m.nc.Publish(msg.Reply, resp)
		}
	})
	m.nc.Flush()
}
