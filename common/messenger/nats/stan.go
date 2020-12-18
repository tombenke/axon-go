package nats

import (
	"github.com/nats-io/stan.go"
	"github.com/tombenke/axon-go/common/messenger"
)

// PublishDurable will publish to the cluster into the `channel` and wait for an ACK.
func (m connections) PublishDurable(channel string, data []byte) error {
	return m.sc.Publish(channel, data)
}

// PublishAsyncDurable will publish to the cluster and asynchronously process
// the ACK or error state. It will return the GUID for the message being sent.
func (m connections) PublishAsyncDurable(channel string, data []byte, ackHandler messenger.AckHandler) (string, error) {
	return m.sc.PublishAsync(channel, data, stan.AckHandler(ackHandler))
}

// Subscribe to the durable `channel`, and call `cb` with the received content.
// Automatically acknowledges to the channel the take-over of the message.
func (m connections) SubscribeDurable(channel string, cb func([]byte)) {
	_, err := m.sc.Subscribe(channel, func(msg *stan.Msg) {
		m.logger.Debugf("Received message from '%s'\n", channel)
		cb(msg.Data)
	})
	if err != nil {
		m.logger.Error(err)
	}
}

// Subscribe to the durable `channel`, and call `cb` with the received content.
// The second argument of the `cb` callback is the acknowledge callback function,
// that has to be called by the consumer of the content.
func (m connections) SubscribeDurableWithAck(channel string, cb func([]byte, func() error)) {
	_, err := m.sc.Subscribe(channel, func(msg *stan.Msg) {
		m.logger.Debugf("Received message from '%s'\n", channel)
		cb(msg.Data, func() error {
			if err := msg.Ack(); err != nil {
				m.logger.Errorf("Failed to ACK msg: %d", msg.Sequence)
				return err
			}
			return nil
		})
	}, stan.SetManualAckMode())
	if err != nil {
		m.logger.Error(err)
	}
}

// QueueSubscribeDurable will perform a queue subscription with the given options to the cluster.
//
// If no option is specified, DefaultSubscriptionOptions are used. The default start
// position is to receive new messages only (messages published after the subscription is
// registered in the cluster).
func (m connections) SubscribeQueueGroupDurable(channel, qgroup string, cb stan.MsgHandler, opts ...stan.SubscriptionOption) (stan.Subscription, error) {
	return m.sc.QueueSubscribe(channel, qgroup, cb, opts...)
}
