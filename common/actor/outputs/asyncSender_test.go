package outputs

import (
	messengerImpl "github.com/tombenke/axon-go/common/messenger/nats"
	at "github.com/tombenke/axon-go/common/testing"
	"sync"
	"testing"
	"time"
)

func TestAsyncSender(t *testing.T) {
	// Connect to messaging
	m := messengerImpl.NewMessenger(messengerCfg)
	defer m.Close()

	// Use a WaitGroup to wait for the processes of the testbed to complete their mission
	wg := sync.WaitGroup{}

	// Create a trigger channel to start the test
	triggerCh := make(chan bool)

	// Start the processes of the test-bed
	doneChkCh := make(chan bool)
	reportCh, testCompletedCh, chkStoppedCh := at.ChecklistProcess(asyncChecklist, doneChkCh, &wg, logger)

	doneRcvCh := make(chan bool)
	rcvStoppedCh := startMockMessageReceivers(getOutputsData(), reportCh, doneRcvCh, &wg, logger, m)
	time.Sleep(200 * time.Millisecond)

	doneProcCh := make(chan bool)
	outputsCh, procStoppedCh := startMockProcessor(triggerCh, reportCh, doneProcCh, &wg, logger)

	// Start the sender process
	doneSndCh := make(chan bool)
	senderStoppedCh := AsyncSender(actorName, outputsCh, doneSndCh, &wg, m, logger)

	// Start testing
	logger.Infof("Send trigger to start testing")
	triggerCh <- true

	// Wait until test is completed, then stop the processes
	logger.Infof("Wait until test is completed")
	<-testCompletedCh

	logger.Infof("Stops Sender")
	close(doneSndCh)
	logger.Infof("Wait Sender to stop")
	<-senderStoppedCh
	logger.Infof("Sender stopped")

	logger.Infof("Stops Mock Processor")
	close(doneProcCh)
	logger.Infof("Wait Mock Processor to stop")
	<-procStoppedCh
	logger.Infof("Mock Processor stopped")

	logger.Infof("Stops Stops Mock Receiver")
	close(doneRcvCh)
	logger.Infof("Wait Mock Receiver to stop")
	<-rcvStoppedCh
	logger.Infof("Mock Receiver stopped")

	logger.Infof("Stops Checklist")
	close(doneChkCh)
	logger.Infof("Wait Checklist to stop")
	<-chkStoppedCh
	logger.Infof("Checklist stopped")

	// Wait for the message to come in
	wg.Wait()
}
