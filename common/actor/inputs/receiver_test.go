package inputs

import (
	"github.com/sirupsen/logrus"
	"github.com/tombenke/axon-go/common/config"
	"github.com/tombenke/axon-go/common/io"
	"github.com/tombenke/axon-go/common/messenger"
	messengerImpl "github.com/tombenke/axon-go/common/messenger/nats"
	"github.com/tombenke/axon-go/common/msgs/base"
	at "github.com/tombenke/axon-go/common/testing"
	"sync"
	"testing"
	"time"
)

const (
	checkSendMsgToInput          = "mock '...' actor sent message to input '...' port"
	checkSendReceiveAndProcess   = "orchestrator sent 'receive-and-process' message"
	checkProcessorReceiveOutputs = "processor received outputs"
)

var checklistDefaultsOnly = []string{
	checkSendReceiveAndProcess,
	checkProcessorReceiveOutputs,
}

var checklistFull = []string{
	checkSendMsgToInput,
	checkSendReceiveAndProcess,
	checkProcessorReceiveOutputs,
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
		Name:           "_timestamp",
		Type:           "base/Int64",
		Representation: "application/json",
		Channel:        "_timestamp",
	}, Default: `{"Body": {"Data": 1608732048980057025}}`},
	config.In{IO: config.IO{
		Name:           "_dt",
		Type:           "base/Float64",
		Representation: "application/json",
		Channel:        "_dt",
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
	"_timestamp":                    base.NewInt64Message(1000),
	"_dt":                           base.NewFloat64Message(1000),
	"well-water-upper-level-state":  base.NewBoolMessage(false),
	"well-water-lower-level-state":  base.NewBoolMessage(false),
	"buffer-tank-upper-level-state": base.NewBoolMessage(false),
	"well-pump-controller-state":    base.NewStringMessage("REFILL-THE-WELL"),
}

// TestReceiveStartStop sets up the input ports, then stops.
func TestReceiverStartStop(t *testing.T) {
	// Connect to messaging
	m := messengerImpl.NewMessenger(messengerCfg)
	defer m.Close()

	// Use a WaitGroup to wait for the processes of the testbed to complete their mission
	wg := sync.WaitGroup{}

	// Create a channel to shut down the processes if needed
	doneCh := make(chan bool)

	// Start the receiver process
	Receiver(inputsCfg, doneCh, &wg, m, logger)

	// Wait until test is completed, then stop the processes
	close(doneCh)

	// Wait for the message to come in
	wg.Wait()
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

	// Start the processes of the test-bed
	reportCh, testCompletedCh := at.ChecklistProcess(checklistDefaultsOnly, doneCh, &wg, logger)

	// Create a trigger channel to start the test
	triggerOrchCh := startMockOrchestrator(reportCh, doneCh, &wg, m, logger)

	// Start the receiver process
	inputsCh := Receiver(inputsCfg, doneCh, &wg, m, logger)

	startMockProcessor(inputsCh, reportCh, doneCh, &wg, logger)

	// Start testing
	time.Sleep(10 * time.Millisecond)
	triggerOrchCh <- true

	// Wait until test is completed, then stop the processes
	<-testCompletedCh
	close(doneCh)

	// Wait for the message to come in
	wg.Wait()
}

// TestReceiveInputs sets up the input ports, and gets inputs to each ports, then a receive-and-process message,
// It uses the incoming messages that it sends as the result inputs to the processor.
func TestReceiverInputs(t *testing.T) {
	// Connect to messaging
	m := messengerImpl.NewMessenger(messengerCfg)
	defer m.Close()

	// Use a WaitGroup to wait for the processes of the testbed to complete their mission
	wg := sync.WaitGroup{}

	// Create a channel to shut down the processes if needed
	doneCh := make(chan bool)

	// Start the processes of the test-bed
	reportCh, testCompletedCh := at.ChecklistProcess(checklistFull, doneCh, &wg, logger)

	// Create a trigger channel to start the test
	triggerOrchCh := startMockOrchestrator(reportCh, doneCh, &wg, m, logger)

	// Start the receiver process
	inputsCh := Receiver(inputsCfg, doneCh, &wg, m, logger)

	startMockProcessor(inputsCh, reportCh, doneCh, &wg, logger)

	// Start testing
	sendInputMessages(inputsCfg, inputs, reportCh, m, logger)
	time.Sleep(10 * time.Millisecond)
	triggerOrchCh <- true

	// Wait until test is completed, then stop the processes
	<-testCompletedCh
	close(doneCh)

	// Wait for the message to come in
	wg.Wait()
}

// TestReceiveInputsBulk sets up the input ports, and gets inputs to each ports, then a receive-and-process message,
// In some ports more than one inputs arrive, so it uses the latest one arrived to the port,
// that it sends to the processor.
func TestReceiverInputsBulk(t *testing.T) {
	// TODO
}

// startMockOrchestrator starts a standalone process that emulates
// the behaviour of an external orchestrator application.
// Orchestrator waits for a trigger to send `receive-and-process` message to the input aggregator.
// The Mock Orchestrator reports every relevant event to the Checklist process.
// Mock Orchestrator will shut down if it receives a message via the `doneCh` channel.
func startMockOrchestrator(reportCh chan string, doneCh chan bool, wg *sync.WaitGroup, m messenger.Messenger, logger *logrus.Logger) chan bool {
	triggerOrchCh := make(chan bool)

	wg.Add(1)
	go func() {
		defer close(triggerOrchCh)
		defer wg.Done()

		for {
			select {
			case <-doneCh:
				logger.Infof("Mock Orchestrator shuts down.")
				return

			case <-triggerOrchCh:
				// TODO: Define and use message type with timestamp, and dt properties
				m.Publish("receive-and-process", []byte("receive-and-process-msg"))
				logger.Infof("Mock Orchestrator sent 'receive-and-process' message.")
				reportCh <- checkSendReceiveAndProcess
			}
		}
	}()

	logger.Infof("Mock Orchestrator started.")
	return triggerOrchCh
}

// startMockProcessor starts a mock processor process that observes the `inputsCh` channel.
// If arrives an inputs data package, checks it content and reports the result to the Checklist process.
// Mock Processor will shut down if it receives a message via the `doneCh` channel.
func startMockProcessor(inputsCh chan io.Inputs, reportCh chan string, doneCh chan bool, wg *sync.WaitGroup, logger *logrus.Logger) {
	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			select {
			case <-doneCh:
				logger.Infof("Mock Processor shuts down.")
				return

			case <-inputsCh:
				// TODO: Compare inputs got to expected default values
				logger.Infof("Mock Processor received inputs to process.")
				reportCh <- checkProcessorReceiveOutputs
			}
		}
	}()
	logger.Infof("Mock Processor started.")
}

// sensInputmessages take the input port configurations and the test data
// and sends them to the Receiver through the channels defined to the corresponding ports.
func sendInputMessages(inputsCfg config.Inputs, inputs at.TestCaseMsgs, reportCh chan string, m messenger.Messenger, logger *logrus.Logger) {

	inputPorts := io.NewInputs(inputsCfg)
	for p := range inputPorts {
		portName := p
		message := inputs[p]
		channel := inputPorts[portName].Channel
		representation := inputPorts[portName].Representation
		logger.Infof("Publish '%v' format message '%v' to '%s' channel.", representation, message, channel)
		m.Publish(channel, message.Encode(representation))
	}
	reportCh <- checkSendMsgToInput
}
