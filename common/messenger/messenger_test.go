package messenger

import (
	"github.com/nats-io/nats.go"
	"testing"
)

var (
	testMessengerConfig = MessengerConfig{
		DefaultNatsURL(),
		DefaultNatsUserCreds,
		DefaultNatsName,
		DefaultNatsClusterID,
		DefaultNatsClientID,
		DefaultLogLevel,
	}
	testMsgContent = []byte("Some text to send...")
)

func TestSetupDefaultConnOptions(t *testing.T) {
	opts := []nats.Option{nats.Name("natsTest")}
	opts = setupDefaultConnOptions(opts)

	if l := len(opts); l != 6 {
		t.Error("setupConnOptions should return with 6 options")
	}
}
