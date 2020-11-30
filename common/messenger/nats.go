package messenger

import (
	nats "github.com/nats-io/nats.go"
)

// Publish `msg` message to the `subject` topic
func (m messenger) Publish(subject string, msg []byte) error {
	subj := subject

	m.nc.Publish(subj, msg)
	m.nc.Flush()

	err := m.nc.LastError()
	if err != nil {
		logger.Error(err)
	}

	logger.Infof("Published message to '%s'\n", subj)
	return err
}

// Subscribe to the `subject` topic, and calls the `cb` call-back function with the inbound messages
func (m messenger) Subscribe(subject string, cb func([]byte)) {
	m.nc.Subscribe(subject, func(msg *nats.Msg) {
		logger.Debugf("Received message from '%s'\n", subject)
		cb(msg.Data)
	})
	m.nc.Flush()
}

// SubscribeSync to the `subject` topic, and calls the `cb` call-back function with the inbound messages
func (m messenger) SubscribeSync(subject string, cb func(*nats.Msg, func([]byte) error)) {
	m.nc.Subscribe(subject, func(msg *nats.Msg) {
		cb(msg, msg.Respond)
	})
	m.nc.Flush()
}
