// Package nats package implements the basic functions to communicate through the messaging middleware.
// The communicating parties see the inbound and outbound messages as akind of message streams.
// This package implements the NATS version of streams functions.
package nats

import (
	"time"

	nats "github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	"github.com/tombenke/axon-go/common/messenger"
)

const (
	// DefaultNatsUserCreds holds the credentials to connect to the NATS server
	DefaultNatsUserCreds = ""
	// DefaultClientName holds the default value of the client name config parameter for the NATS connection
	DefaultClientName = "messenger"
	// DefaultNatsClusterID is the default config parameter value for connecting to the NATS cluster
	DefaultNatsClusterID = "test-cluster"
	// DefaultNatsClientID is the default config parameter value for the Messenger client that connects to the NATS cluster
	DefaultNatsClientID = "test-client"
)

// DefaultNatsURL is the default URL of the NATS messaging server
func DefaultNatsURL() string {
	return nats.DefaultURL
}

type connections struct {
	nc     *nats.Conn
	sc     stan.Conn
	logger *logrus.Logger
}

// NewMessenger creates a new Messenger instance using the configuration parameters
func NewMessenger(config messenger.Config) messenger.Messenger {
	nc, err := natsConnect(config)
	if err != nil {
		config.Logger.Fatal(err)
	}
	sc := stanConnect(nc, config)
	m := connections{nc, sc, config.Logger}
	return m
}

// setupOptions extends the options array with default configuration parameters
// that are useful to connect to the NATS server.
func setupDefaultConnOptions(opts []nats.Option, logger *logrus.Logger) []nats.Option {
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
func natsConnect(config messenger.Config) (*nats.Conn, error) {
	// Connect Options.
	opts := []nats.Option{nats.Name(config.ClientName)}
	opts = setupDefaultConnOptions(opts, config.Logger)

	// Use UserCredentials
	if config.UserCreds != "" {
		opts = append(opts, nats.UserCredentials(config.UserCreds))
	}

	return nats.Connect(config.Urls, opts...)
}

// Connect to the NATS streaming server and returns with a `stan.Conn` that can be used for further operations.
func stanConnect(nc *nats.Conn, config messenger.Config) stan.Conn {
	sc, err := stan.Connect(config.ClusterID, config.ClientID, stan.NatsConn(nc),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			config.Logger.Fatalf("Connection lost, reason: %v", reason)
		}))
	if err != nil {
		config.Logger.Fatalf("Can't connect: %v.\nMake sure a NATS Streaming Server is running at: %s", err, config.Urls)
	}
	config.Logger.Infof("Connected to %s clusterID: [%s] clientID: [%s]\n", config.Urls, config.ClusterID, config.ClientID)
	return sc
}

// Close both the Streaming and the NATS connections
func (s connections) Close() {
	// Close the Streaming connection
	s.sc.Close()

	// Close the NATS connection
	s.nc.Close()
}
