// Package messenger package implements the basic functions to communicate through the messaging middleware.
// The agents see the inbound and outbound messages as akind of message streams.
// This package implements the NATS version of streams functions.
package messenger

import (
	"time"

	nats "github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	"github.com/tombenke/axon-go/common/log"
)

var logger *logrus.Logger = log.Logger

const (
	DefaultNatsUserCreds = ""
	DefaultNatsName      = "messenger"
	DefaultNatsClusterID = "test-cluster"
	DefaultNatsClientID  = "test-client"
	DefaultLogLevel      = "info"
)

type MessengerConfig struct {
	Urls      string
	UserCreds string
	NatsName  string
	ClusterID string
	ClientID  string
	LogLevel  string
}

// DefaultNatsURL is the default URL of the NATS messaging server
func DefaultNatsURL() string {
	return nats.DefaultURL
}

type Messenger interface {
	// Non durable subjects (simple NATS)
	Publish(string, []byte) error
	Subscribe(string, func([]byte))
	SubscribeSync(subject string, cb func(*nats.Msg, func([]byte) error))

	// Durable channels (STAN)
	PublishDurable(string, []byte) error
	PublishAsyncDurable(string, []byte, stan.AckHandler) (string, error)
	SubscribeDurable(string, func([]byte))
	SubscribeDurableWithAck(string, func([]byte, func() error))
	SubscribeQueueGroupDurable(string, string, stan.MsgHandler, ...stan.SubscriptionOption) (stan.Subscription, error)

	// Close both NATS and STAN
	Close()
}

type messenger struct {
	nc *nats.Conn
	sc stan.Conn
}

// Create a new Messenger using the configuration parameters
func NewMessenger(config MessengerConfig) Messenger {
	log.SetLevelStr(config.LogLevel)
	nc, err := natsConnect(config)
	if err != nil {
		logger.Fatal(err)
	}
	sc := stanConnect(nc, config)
	m := messenger{nc, sc}
	return m
}

// setupOptions extends the options array with default configuration parameters
// that are useful to connect to the NATS server.
func setupDefaultConnOptions(opts []nats.Option) []nats.Option {
	totalWait := 10 * time.Minute
	reconnectDelay := time.Second

	opts = append(opts, nats.ReconnectWait(reconnectDelay))
	opts = append(opts, nats.MaxReconnects(int(totalWait/reconnectDelay)))
	opts = append(opts, nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
		if err != nil {
			logger.Errorf("disconnect error: %s", err)
		}
		logger.Errorf("disconnected: will attempt reconnects for %.0fm", totalWait.Minutes())
	}))
	opts = append(opts, nats.ReconnectHandler(func(nc *nats.Conn) {
		logger.Infof("reconnected [%s]", nc.ConnectedUrl())
	}))
	opts = append(opts, nats.ClosedHandler(func(nc *nats.Conn) {
		logger.Errorf("exiting: %s", nc.LastError())
	}))
	return opts
}

// Connect to the NATS server and returns with a `*nats.Conn` that can be used for further operations.
func natsConnect(config MessengerConfig) (*nats.Conn, error) {
	// Connect Options.
	opts := []nats.Option{nats.Name(config.NatsName)}
	opts = setupDefaultConnOptions(opts)

	// Use UserCredentials
	if config.UserCreds != "" {
		opts = append(opts, nats.UserCredentials(config.UserCreds))
	}

	return nats.Connect(config.Urls, opts...)
}

// Connect to the NATS streaming server and returns with a `stan.Conn` that can be used for further operations.
func stanConnect(nc *nats.Conn, config MessengerConfig) stan.Conn {
	sc, err := stan.Connect(config.ClusterID, config.ClientID, stan.NatsConn(nc),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			logger.Fatalf("Connection lost, reason: %v", reason)
		}))
	if err != nil {
		logger.Fatalf("Can't connect: %v.\nMake sure a NATS Streaming Server is running at: %s", err, config.Urls)
	}
	logger.Infof("Connected to %s clusterID: [%s] clientID: [%s]\n", config.Urls, config.ClusterID, config.ClientID)
	return sc
}

// Close both the Streaming and the NATS connections
func (s messenger) Close() {
	// Close the Streaming connection
	s.sc.Close()

	// Close the NATS connection
	s.nc.Close()
}
