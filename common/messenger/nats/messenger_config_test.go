package nats

import (
	"github.com/nats-io/nats.go"
	"github.com/tombenke/axon-go/common/log"
	"github.com/tombenke/axon-go/common/messenger"
	"testing"
)

var (
	testConfig = messenger.Config{
		Urls:       DefaultNatsURL(),
		UserCreds:  DefaultNatsUserCreds,
		ClientName: DefaultClientName,
		ClusterID:  DefaultNatsClusterID,
		ClientID:   DefaultNatsClientID,
		Logger:     log.Logger,
	}
)

func TestSetupDefaultConnOptions(t *testing.T) {
	opts := []nats.Option{nats.Name("natsTest")}
	opts = setupDefaultConnOptions(opts, log.Logger)

	if l := len(opts); l != 6 {
		t.Error("setupConnOptions should return with 6 options")
	}
}
