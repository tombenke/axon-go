package processor

import (
	//"github.com/stretchr/testify/assert"
	"github.com/sirupsen/logrus"
	"github.com/tombenke/axon-go/common/config"
	"github.com/tombenke/axon-go/common/io"
	"github.com/tombenke/axon-go/common/msgs/base"
	at "github.com/tombenke/axon-go/common/testing"
	"sync"
	"testing"
	//"time"
)

const (
	checkReceiverSentInputs    = "inputs receiver sent inputs"
	checkSenderReceivedOutputs = "outputs sender received outputs"
)

var checklist = []string{
	checkReceiverSentInputs,
	checkSenderReceivedOutputs,
}

var inputsCfg = config.Inputs{
	config.In{IO: config.IO{
		Name:           "max-power",
		Type:           "base/Float64",
		Representation: "application/json",
		Channel:        "",
	}, Default: `{"Body": {"Data": 2000.0}}`},
	config.In{IO: config.IO{
		Name:           "power-need",
		Type:           "base/Float64",
		Representation: "application/json",
		Channel:        "well-pump-relay.electric-power-need",
	}, Default: ""},
}

var outputsCfg = config.Outputs{
	config.Out{IO: config.IO{
		Name:           "power-output",
		Type:           "base/Float64",
		Representation: "application/json",
		Channel:        "well-pump-relay.electric-power-input",
	}},
}

var testCase = at.TestCase{
	Inputs: at.TestCaseMsgs{
		"max-power":  base.NewFloat64Message(2000.0),
		"power-need": base.NewFloat64Message(4599.0),
	},
	Outputs: at.TestCaseMsgs{
		"power-output": base.NewFloat64Message(2000.0),
	},
}

func TestStartProcessor(t *testing.T) {

	var logger = logrus.New()

	// Use a WaitGroup to wait for the processes of the testbed to complete their mission
	wg := sync.WaitGroup{}

	// Create a trigger channel to start the test
	triggerCh := make(chan bool)

	// Start the processes of the test-bed
	doneCheckCh := make(chan bool)
	reportCh, testCompletedCh, checklistStoppedCh := at.ChecklistProcess(checklist, doneCheckCh, &wg, logger)

	// Create a channel to feed inputs to the Processor
	doneRcvCh := make(chan bool)
	inputsCh, mockRcvStoppedCh := StartMockReceiver(triggerCh, reportCh, doneRcvCh, &wg, logger)

	doneProcCh := make(chan bool)
	outputsCh, procStoppedCh := StartProcessor(ProcessorFun, outputsCfg, doneProcCh, &wg, inputsCh, logger)

	doneSndCh := make(chan bool)
	mockSndStoppedCh := StartMockSender(t, outputsCh, reportCh, doneSndCh, &wg, logger)

	// Start testing
	//time.Sleep(10 * time.Millisecond)
	triggerCh <- true

	// Wait until test is completed, then stop the processes
	logger.Infof("Wait until test is completed")
	<-testCompletedCh

	logger.Infof("Stops Sender")
	close(doneSndCh)
	logger.Infof("Wait Sender to stop")
	<-mockSndStoppedCh
	logger.Infof("Sender stopped")

	logger.Infof("Stops processor")
	close(doneProcCh)
	logger.Infof("Wait Processor to stop")
	<-procStoppedCh
	logger.Infof("Processor stopped")

	logger.Infof("Stops Receiver")
	close(doneRcvCh)
	logger.Infof("Wait Receiver to stop")
	<-mockRcvStoppedCh
	logger.Infof("Receiver stopped")

	logger.Infof("Stops Checklist")
	close(doneCheckCh)
	logger.Infof("Wait Checklist to stop")
	<-checklistStoppedCh
	logger.Infof("Checklist stopped")

	// Wait for the message to come in
	wg.Wait()
}

// ProcessorFun is the message processor function of the actor node
func ProcessorFun(ctx Context) error {
	maxPower := ctx.GetInputMessage("max-power").(*base.Float64).Body.Data
	powerNeed := ctx.GetInputMessage("power-need").(*base.Float64).Body.Data

	var powerOutput float64
	if powerNeed > maxPower {
		powerOutput = maxPower
	} else {
		powerOutput = powerNeed
	}

	ctx.SetOutputMessage("power-output", base.NewFloat64Message(powerOutput))
	return nil
}

func StartMockReceiver(triggerCh chan bool, reportCh chan string, doneCh chan bool, wg *sync.WaitGroup, logger *logrus.Logger) (chan io.Inputs, chan bool) {
	logger.Infof("Mock Receiver started.")
	inputsCh := make(chan io.Inputs)
	mockRcvStoppedCh := make(chan bool)
	defer close(mockRcvStoppedCh)
	inputs := io.NewInputs(inputsCfg)
	SetInputs(&inputs, testCase.Inputs)

	wg.Add(1)
	go func() {
		//defer close(inputsCh)
		defer logger.Infof("Mock Receiver stopped.")
		defer wg.Done()

		for {
			select {
			case <-doneCh:
				logger.Infof("Mock Receiver shuts down.")
				return

			case <-triggerCh:
				inputsCh <- inputs
				logger.Infof("Mock Receiver sent inputs")
				reportCh <- checkReceiverSentInputs
			}
		}
	}()

	return inputsCh, mockRcvStoppedCh
}

func StartMockSender(t *testing.T, outputsCh chan io.Outputs, reportCh chan string, doneCh chan bool, wg *sync.WaitGroup, logger *logrus.Logger) chan bool {
	logger.Infof("Mock Sender started.")
	mockSndStoppedCh := make(chan bool)

	wg.Add(1)
	go func() {
		defer close(mockSndStoppedCh)
		defer logger.Infof("Mock Sender stopped.")
		defer wg.Done()

		for {
			select {
			case <-doneCh:
				logger.Infof("Mock Sender shuts down.")
				return

			case outputs := <-outputsCh:
				logger.Infof("Mock Sender received outputs")
				CompareOutputsData(t, outputs, testCase)
				logger.Infof("Mock Sender reports that received outputs")
				reportCh <- checkSenderReceivedOutputs
			}
		}
	}()

	return mockSndStoppedCh
}
