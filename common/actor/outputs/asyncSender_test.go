package outputs

import (
	messengerImpl "github.com/tombenke/axon-go/common/messenger/nats"
	at "github.com/tombenke/axon-go/common/testing"
	"sync"
	"testing"
)

func TestAsyncSender(t *testing.T) {
	// Connect to messaging
	m := messengerImpl.NewMessenger(messengerCfg)
	defer m.Close()

	// Use a WaitGroup to wait for the processes of the testbed to complete their mission
	wg := sync.WaitGroup{}

	// Create a channel to shut down the processes if needed
	doneCh := make(chan bool)

	// Create a trigger channel to start the test
	triggerCh := make(chan bool)

	// Start the processes of the test-bed
	reportCh, testCompletedCh := at.ChecklistProcess(asyncChecklist, doneCh, &wg, logger)
	//startMockOrchestrator(t, reportCh, doneCh, &wg, logger, m)
	startMockMessageReceivers(getOutputsData(), reportCh, doneCh, &wg, logger, m)
	outputsCh := startMockProcessor(triggerCh, reportCh, doneCh, &wg, logger)

	// Start the sender process
	wg.Add(1)
	go AsyncSender(actorName, outputsCh, doneCh, &wg, m, logger)

	// Start testing
	triggerCh <- true

	// Wait until test is completed, then stop the processes
	<-testCompletedCh
	close(doneCh)

	// Wait for the message to come in
	wg.Wait()
}
