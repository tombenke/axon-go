package outputs

import (
	"github.com/sirupsen/logrus"
	"github.com/tombenke/axon-go/common/config"
	"github.com/tombenke/axon-go/common/io"
	"github.com/tombenke/axon-go/common/messenger"
	"github.com/tombenke/axon-go/common/msgs/base"
	at "github.com/tombenke/axon-go/common/testing"
	"sync"
)

const (
	actorName                              = "test-actor"
	checkSendOutputs                       = "processor sent outputs to sender"
	checkSendingCompleted                  = "orchestrator received sending-completed"
	checkMsgArrivedWellPumpRelayState      = "well-pump-relay-state message arrived"
	checkMsgArrivedWellPumpControllerState = "well-pump-controller-state message arrived"
)

var syncChecklist = []string{
	checkSendOutputs,
	checkSendingCompleted,
	checkMsgArrivedWellPumpRelayState,
	checkMsgArrivedWellPumpControllerState,
}

var asyncChecklist = []string{
	checkSendOutputs,
	checkMsgArrivedWellPumpRelayState,
	checkMsgArrivedWellPumpControllerState,
}

var logger = logrus.New()

var messengerCfg = messenger.Config{
	Urls:       "localhost:4222",
	UserCreds:  "",
	ClientName: "sender-test-client",
	ClusterID:  "test-cluster",
	ClientID:   "sender-test-client",
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

func getOutputsData() io.Outputs {
	// Create Output ports
	outputs := io.NewOutputs(outputsCfg)

	// Set actual messages to the ports
	for portName := range outputMsgs {
		(&outputs).SetMessage(portName, outputMsgs[portName])
	}

	return outputs
}

// startMockProcessor starts a standalone process that emulates the behaviour of the actor's processor.
// It creates and returns an `outputs` channel for the SyncSender.
// MockProcessor waits for a trigger message via the `trigger` channel
// then sends `io.Outputs` test data package through the `outputs` channel to the SyncSender.
// MockProcessor will shut down if it receives a message via the `doneCh` channel.
func startMockProcessor(triggerCh chan bool, reportCh chan string, doneCh chan bool, wg *sync.WaitGroup, logger *logrus.Logger) (chan io.Outputs, chan bool) {
	outputsCh := make(chan io.Outputs)
	procStoppedCh := make(chan bool)

	wg.Add(1)
	go func() {
		defer logger.Infof("MockProcessor stopped.")
		defer close(outputsCh)
		defer close(procStoppedCh)
		defer wg.Done()

		for {
			select {
			case <-doneCh:
				logger.Infof("MockProcessor shuts down.")
				return

			case <-triggerCh:
				logger.Infof("MockProcessor sends outputs test data")
				outputs := getOutputsData()
				outputsCh <- outputs
				reportCh <- checkSendOutputs
			}
		}
	}()

	logger.Infof("Mock Processor started.")
	return outputsCh, procStoppedCh
}

// startMockMessageReceivers starts a process to each output message to send,
// and checks if the messages are really sent to the expected channel.
// Returns the number processes forked, that is actually the number of output ports.
func startMockMessageReceivers(outputs io.Outputs, reportCh chan string, doneCh chan bool, wg *sync.WaitGroup, logger *logrus.Logger, m messenger.Messenger) chan bool {
	rcvStoppedCh := make(chan bool)

	wg.Add(1)
	go func() {
		defer logger.Infof("Receiver master stopped.")
		defer close(rcvStoppedCh)
		defer wg.Done()

		// Create wait-group for the channel observer sub-processes
		rcvWg := sync.WaitGroup{}
		rcvDoneCh := make(chan bool)
		startSubReceivers(outputs, reportCh, rcvDoneCh, &rcvWg, logger, m)

		// Waits until receives a done signal
		<-doneCh
		logger.Infof("Receiver master shuts down.")
		close(rcvDoneCh)
		logger.Infof("Receiver master closed the 'rcvDoneCh'.")
		logger.Infof("Receiver master starts waiting for sub-receivers to stop")
		rcvWg.Wait()
		logger.Infof("Receiver master's sub-receivers stopped")
	}()

	defer logger.Infof("Receiver master started.")
	return rcvStoppedCh
}

func startSubReceivers(outputs io.Outputs, reportCh chan string, doneCh chan bool, wg *sync.WaitGroup, logger *logrus.Logger, m messenger.Messenger) int {

	for o := range outputs {
		wg.Add(1)

		go func(outName string) {
			name := outputs[outName].Name
			channel := outputs[outName].Channel
			logger.Infof("Start Message Receiver '%s >> %s'", name, channel)
			messageReceivedCh := make(chan []byte)
			messageReceivedSubs := m.ChanSubscribe(channel, messageReceivedCh)

			defer func() {
				logger.Infof("Message Receiver '%s' stopped.", channel)
				if err := messageReceivedSubs.Unsubscribe(); err != nil {
					panic(err)
				}
				close(messageReceivedCh)
				wg.Done()
			}()

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
