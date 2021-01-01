package inputs

import (
	"github.com/sirupsen/logrus"
	"github.com/tombenke/axon-go/common/config"
	//"github.com/tombenke/axon-go/common/io"
	"github.com/tombenke/axon-go/common/messenger"
	messengerImpl "github.com/tombenke/axon-go/common/messenger/nats"
	"github.com/tombenke/axon-go/common/msgs/base"
	at "github.com/tombenke/axon-go/common/testing"
	"sync"
	"testing"
)

const (
	checkSendMsgToInput          = "mock '...' actor sent message to input '...' port"
	checkSendReceiveAndProcess   = "orchestrator sent 'receive-and-process' message"
	checkProcessorReceiveOutputs = "processor received outputs"
)

var checklist = []string{
	//checkSendMsgToInput,
	//checkSendReceiveAndProcess,
	//checkProcessorReceiveOutputs,
}

var logger = logrus.New()

var messengerCfg = messenger.Config{
	Urls:       "localhost:4222",
	UserCreds:  "",
	ClientName: "receiver-test-client",
	ClusterID:  "test-cluster",
	ClientID:   "receiver-test-client",
	Logger:     logger,
}

var inputsCfg = config.Inputs{
	config.In{IO: config.IO{
		Name:           "dt",
		Type:           "base/Float64",
		Representation: "application/json",
		Channel:        "dt",
	}, Default: `{"Body": {"Data": 1000}}`},
	config.In{IO: config.IO{
		Name:           "well-water-upper-level-state",
		Type:           "base/Bool",
		Representation: "application/json",
		Channel:        "well-water-upper-level-state",
	}, Default: `{"Body": {"Data": false}}`},
	config.In{IO: config.IO{
		Name:           "well-water-lower-level-state",
		Type:           "base/Bool",
		Representation: "application/json",
		Channel:        "well-water-lower-level-state",
	}, Default: `{"Body": {"Data": false}}`},
	config.In{IO: config.IO{
		Name:           "buffer-tank-upper-level-state",
		Type:           "base/Bool",
		Representation: "application/json",
		Channel:        "buffer-water-tank-upper-level-state",
	}, Default: `{"Body": {"Data": false}}`},
	config.In{IO: config.IO{
		Name:           "well-pump-controller-state",
		Type:           "base/String",
		Representation: "application/json",
		Channel:        "well-pump-controller-state",
	}, Default: `{"Body": {"Data": "REFILL-THE-WELL"}}`},
}

var inputs = at.TestCaseMsgs{
	"dt":                     base.NewFloat64Message(1000),
	"well-backfill-capacity": base.NewFloat64Message(0.6),
	"well-cross-section":     base.NewFloat64Message(0.0201056),
	"max-water-level":        base.NewFloat64Message(-30),
	"min-water-level":        base.NewFloat64Message(-36),
	"water-need":             base.NewFloat64Message(0),
	"water-level":            base.NewFloat64Message(-30),
}

// TestReceiveDefaultsOnly sets up the input ports, then gets a receive-and-process message,
// but receive no input messages via the ports, so it uses the default values defined to the ports.
// It sends the result inputs to the processor.
func TestReceiverDefaultsOnly(t *testing.T) {
	// Connect to messaging
	m := messengerImpl.NewMessenger(messengerCfg)
	defer m.Close()

	// Use a WaitGroup to wait for the processes of the testbed to complete their mission
	wg := sync.WaitGroup{}

	// Create a channel to shut down the processes if needed
	doneCh := make(chan bool)

	// Create a trigger channel to start the test
	//triggerCh := make(chan bool)

	// Start the processes of the test-bed
	//reportCh, testCompletedCh := at.ChecklistProcess(checklist, doneCh, &wg, logger)
	at.ChecklistProcess(checklist, doneCh, &wg, logger)

	//startMockOrchestrator(reportCh, doneCh, &wg, logger, m)

	//numMsgSenders := startMockMessageSenders(getInputsData(), reportCh, doneCh, &wg, logger, m)

	// Start the receiver process
	//inputsCh := Receiver(inputsCfg, doneCh, &wg, m, logger)
	Receiver(inputsCfg, doneCh, &wg, m, logger)

	//outputsCh := startMockProcessor(inputsCh, reportCh, doneCh, &wg, logger)

	// Start testing
	//triggerCh <- true

	// Wait until test is completed, then stop the processes
	//<-testCompletedCh
	close(doneCh)

	// Wait for the message to come in
	wg.Wait()
}

// TestReceiveDefaultsAndInputs sets up the input ports, gets some inputs, then a receive-and-process message,
// It uses the incoming messages as well as default values in case of ports there were no message arrived to.
// It sends the result inputs to the processor.
func TestReceiverDefaults(t *testing.T) {
}

// TestReceiveInputs sets up the input ports, and gets inputs to each ports, then a receive-and-process message,
// It uses the incoming messages that it sends as the result inputs to the processor.
func TestReceiverInputs(t *testing.T) {
}

// TestReceiveInputsBulk sets up the input ports, and gets inputs to each ports, then a receive-and-process message,
// In some ports more than one inputs arrive, so it uses the latest one arrived to the port,
// that it sends to the processor.
func TestReceiverInputsBulk(t *testing.T) {
}
