// Package messenger package defines the interface for the basic functions to communicate through the messaging middleware.
// The communicating parties see the inbound and outbound messages as akind of message streams.
package messenger

import (
	"github.com/sirupsen/logrus"
	"time"
)

// Config holds the configuration parameters of the Messenger clients
type Config struct {
	Urls       string         `yaml:"urls"`
	UserCreds  string         `yaml:"credentials"`
	ClientName string         `yaml:"-"`
	ClusterID  string         `yaml:"-"`
	ClientID   string         `yaml:"-"`
	Logger     *logrus.Logger `yaml:"-"`
}

// AckHandler is used for Async Publishing to provide status of the ack.
// The func will be passed the GUID and any error state.
// No error means the message was successfully received by streaming channel.
type AckHandler func(string, error)

// Subscriber interface is returned by the Subscribe calls, and provide an `Unsubscribe()` method.
type Subscriber interface {
	Unsubscribe() error
}

// Messenger interface represents the messaging patterns
// that an underlying messaging middleware has to implement
type Messenger interface {
	// Non durable subjects
	Publish(string, []byte) error
	Subscribe(string, func([]byte)) Subscriber
	ChanSubscribe(string, chan []byte) Subscriber
	Request(subject string, msg []byte, timeout time.Duration) ([]byte, error)
	Response(subject string, service func([]byte) ([]byte, error))

	// Durable channels
	PublishDurable(string, []byte) error
	PublishAsyncDurable(string, []byte, AckHandler) (string, error)
	SubscribeDurable(string, func([]byte))
	SubscribeDurableWithAck(string, func([]byte, func() error))
	//SubscribeQueueGroupDurable(string, string, stan.MsgHandler, ...stan.SubscriptionOption) (stan.Subscription, error)

	// Close both non-durable, and durable connections
	Close()
}
