package inputs

import (
	"github.com/sirupsen/logrus"
	"github.com/tombenke/axon-go/common/config"
	"github.com/tombenke/axon-go/common/io"
	"github.com/tombenke/axon-go/common/messenger"
	"github.com/tombenke/axon-go/common/msgs/base"
	"github.com/tombenke/axon-go/common/msgs/orchestra"
	at "github.com/tombenke/axon-go/common/testing"
	"sync"
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

var checklistAsync = []string{
	checkSendMsgToInput,
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

var syncInputsCfg = config.Inputs{
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

var syncInputs = at.TestCaseMsgs{
	"_RAP":                          orchestra.NewReceiveAndProcessMessage(1),
	"well-water-upper-level-state":  base.NewBoolMessage(false),
	"well-water-lower-level-state":  base.NewBoolMessage(false),
	"buffer-tank-upper-level-state": base.NewBoolMessage(false),
	"well-pump-controller-state":    base.NewStringMessage("REFILL-THE-WELL"),
}

var asyncInputsCfg = config.Inputs{
	config.In{IO: config.IO{
		Name:           "well-pump-controller-state",
		Type:           "base/String",
		Representation: "application/json",
		Channel:        "well-pump-controller-state",
	}, Default: `{"Body": {"Data": "REFILL-THE-WELL"}}`},
}

var asyncInputs = at.TestCaseMsgs{
	"well-pump-controller-state": base.NewStringMessage("REFILL-THE-WELL"),
}

// startMockProcessor starts a mock processor process that observes the `inputsCh` channel.
// If arrives an inputs data package, checks it content and reports the result to the Checklist process.
// Mock Processor will shut down if it receives a message via the `doneCh` channel.
func startMockProcessor(inputsCh chan io.Inputs, reportCh chan string, doneCh chan bool, wg *sync.WaitGroup, logger *logrus.Logger) chan bool {
	procStoppedCh := make(chan bool)

	wg.Add(1)
	go func() {
		logger.Infof("Mock Processor started.")
		defer logger.Infof("Mock Processor stopped.")
		defer close(procStoppedCh)
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
				logger.Infof("Mock Processor reported to received inputs to process.")
			}
		}
	}()

	return procStoppedCh
}

// sensInputmessages take the input port configurations and the test data
// and sends them to the Receiver through the channels defined to the corresponding ports.
func sendInputMessages(inputsCfg config.Inputs, inputs at.TestCaseMsgs, reportCh chan string, m messenger.Messenger, logger *logrus.Logger) {

	logger.Infof("Mock Actor start publishing messages to input channels.")
	inputPorts := io.NewInputs(inputsCfg)
	for p := range inputPorts {
		portName := p
		message := inputs[p]
		channel := inputPorts[portName].Channel
		representation := inputPorts[portName].Representation
		logger.Infof("Mock Actor publishes message to '%s' channel.", channel)
		//logger.Infof("Publish '%v' format message '%v' to '%s' channel.", representation, message, channel)
		if err := m.Publish(channel, message.Encode(representation)); err != nil {
			panic(err)
		}
		logger.Infof("Mock Actor published message to '%s' channel.", channel)
	}
	reportCh <- checkSendMsgToInput
	logger.Infof("Mock Actor reported to finish publishing messages to input channels.")
}
