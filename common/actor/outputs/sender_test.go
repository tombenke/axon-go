package outputs

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
)

const (
	checkSendOutputs                       = "processor sent outputs to sender"
	checkSendingCompleted                  = "orchestrator received sending-completed"
	checkMsgArrivedWellPumpRelayState      = "well-pump-relay-state message arrived"
	checkMsgArrivedWellPumpControllerState = "well-pump-controller-state message arrived"
)

var checklist = []string{
	checkSendOutputs,
	checkSendingCompleted,
	checkMsgArrivedWellPumpRelayState,
	checkMsgArrivedWellPumpControllerState,
}

var logger = logrus.New()

var messengerCfg = messenger.Config{
	Urls:       "localhost:4222",
	UserCreds:  "",
	ClientName: "test-client",
	ClusterID:  "test-cluster",
	ClientID:   "test-client",
	Logger:     logger,
}

var outputsCfg = config.Outputs{
	config.Out{IO: config.IO{
		Name:           "well-pump-relay-state",
		Type:           "base/Bool",
		Representation: "application/json",
		Channel:        "well-pump-relay-state",
	}},
	config.Out{IO: config.IO{
		Name:           "well-pump-controller-state",
		Type:           "base/String",
		Representation: "application/json",
		Channel:        "well-pump-controller-state",
	}},
}

var outputMsgs = at.TestCaseMsgs{
	"well-pump-relay-state":      base.NewBoolMessage(false),
	"well-pump-controller-state": base.NewStringMessage("REFILL-THE-WELL"),
}

func TestSender(t *testing.T) {
	// Connect to messaging
	m := messengerImpl.NewMessenger(messengerCfg)
	defer m.Close()

	// Use a WaitGroup to wait for the processes of the testbed to complete their mission
	wg := sync.WaitGroup{}
	wg.Add(4)

	// Create a channel to shut down the processes if needed
	doneCh := make(chan bool)

	// Create a trigger channel to start the test
	triggerCh := make(chan bool)

	// Start the processes of the test-bed
	reportCh, testCompletedCh := ChecklistProcess(doneCh, &wg, logger)
	StartMockOrchestrator(reportCh, doneCh, &wg, logger, m)
	numMsgReceivers := StartMockMessageReceivers(GetOutputsData(), reportCh, doneCh, &wg, logger, m)
	wg.Add(numMsgReceivers)
	outputsCh := StartMockProcessor(triggerCh, reportCh, doneCh, &wg, logger)

	// Start the sender process
	go Sender(outputsCh, doneCh, &wg, m, logger)

	// Start testing
	triggerCh <- true

	// Wait until test is completed, then stop the processes
	<-testCompletedCh
	close(doneCh)

	// Wait for the message to come in
	wg.Wait()
}

func GetOutputsData() io.Outputs {
	// Create Output ports
	outputs := io.NewOutputs(outputsCfg)

	// Set actual messages to the ports
	for portName := range outputMsgs {
		(&outputs).SetMessage(portName, outputMsgs[portName])
	}

	return outputs
}

// StartMockProcessor starts a standalone process that emulates the behaviour of the actor's processor.
// It creates and returns an `outputs` channel for the Sender.
// MockProcessor waits for a trigger message via the `trigger` channel
// then sends `io.Outputs` test data package through the `outputs` channel to the Sender.
// MockProcessor will shut down if it receives a message via the `doneCh` channel.
func StartMockProcessor(triggerCh chan bool, reportCh chan string, doneCh chan bool, wg *sync.WaitGroup, logger *logrus.Logger) chan io.Outputs {
	outputsCh := make(chan io.Outputs)

	go func() {
		defer close(outputsCh)
		defer wg.Done()

		for {
			select {
			case <-doneCh:
				logger.Infof("MockProcessor shuts down.")
				return

			case <-triggerCh:
				logger.Infof("MockProcessor sends outputs test data")
				outputs := GetOutputsData()
				outputsCh <- outputs
				reportCh <- checkSendOutputs
			}
		}
	}()

	logger.Infof("Mock Processor started.")
	return outputsCh
}

// StartMockOrchestrator starts a standalone process that emulates the behaviour of an external orchestrator application.
// Orchestrator waits for an incoming message via the `processing-completed` messaging channel,
// then sends a trigger message to the Sender process via the `send-outputs` messaging channel.
// MockProcessor will shut down if it receives a message via the `doneCh` channel.
func StartMockOrchestrator(reportCh chan string, doneCh chan bool, wg *sync.WaitGroup, logger *logrus.Logger, m messenger.Messenger) {
	processingCompletedCh := make(chan []byte)
	processingCompletedSubs := m.ChanSubscribe("processing-completed", processingCompletedCh)

	sendingCompletedCh := make(chan []byte)
	sendingCompletedSubs := m.ChanSubscribe("sending-completed", sendingCompletedCh)

	go func() {
		defer processingCompletedSubs.Unsubscribe()
		defer close(processingCompletedCh)
		defer sendingCompletedSubs.Unsubscribe()
		defer close(sendingCompletedCh)
		defer wg.Done()

		for {
			select {
			case <-doneCh:
				logger.Infof("MockOrchestrator shuts down.")
				return

			case <-processingCompletedCh:
				logger.Infof("MockOrchestrator received 'processing-completed' message.")
				logger.Infof("MockOrchestrator sends 'send-results' message.")
				m.Publish("send-results", []byte("send-results-msg"))

			case <-sendingCompletedCh:
				logger.Infof("MockOrchestrator received 'sending-completed' message.")
				reportCh <- checkSendingCompleted
			}
		}
	}()
	logger.Infof("Mock Orchestrator started.")
}

// ChecklistProcess starts a process that collects the reports of the processes of the test-bed,
// and checks the test was complete.
// It creates and returns with two channels, one channel for collecting test results,
// and another one that it will close, when the test is finished.
func ChecklistProcess(done chan bool, wg *sync.WaitGroup, logger *logrus.Logger) (chan string, chan bool) {
	testCompletedCh := make(chan bool)
	reportCh := make(chan string)
	reported := make(map[string]bool)

	go func() {
		defer close(testCompletedCh)
		defer close(reportCh)
		defer wg.Done()

		for {
			select {
			case <-done:
				// Test either was completed or it was shut down
				logger.Infof("Checklist Process shuts down.")
				return
			case report := <-reportCh:
				logger.Infof("Checklist received report: '%s'", report)
				reported[report] = true
				if CheckReported(reported) {
					return
				}
			}
		}
	}()

	logger.Infof("Checklist Process started.")
	return reportCh, testCompletedCh
}

// CheckReported compares reported to expected,
// and returns true if all expected item exists in reported array, otherwise returns false.
func CheckReported(reported map[string]bool) bool {
	result := true
	for _, c := range checklist {
		if _, exist := reported[c]; !exist {
			result = false
			break
		}
	}

	return result
}

// StartMockMessageReceivers starts a process to each output message to send,
// and checks if the messages are really sent to the expected channel.
// Returns the number processes forked, that is actually the number of output ports.
func StartMockMessageReceivers(outputs io.Outputs, reportCh chan string, doneCh chan bool, wg *sync.WaitGroup, logger *logrus.Logger, m messenger.Messenger) int {
	for o := range outputs {
		go func(outName string) {
			name := outputs[outName].Name
			channel := outputs[outName].Channel
			logger.Infof("Start Message Receiver '%s >> %s'", name, channel)
			messageReceivedCh := make(chan []byte)
			messageReceivedSubs := m.ChanSubscribe(channel, messageReceivedCh)

			defer messageReceivedSubs.Unsubscribe()
			defer close(messageReceivedCh)
			defer wg.Done()

			for {
				select {
				case <-doneCh:
					logger.Infof("Message Receiver '%s' shuts down.", channel)
					return
				case <-messageReceivedCh:
					logger.Infof("Message Receiver received '%s'", name)
					reportCh <- name + " message arrived"
				}
			}
		}(o)
	}
	return len(outputs)
}
