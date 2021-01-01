package testing

import (
	"github.com/sirupsen/logrus"
	"sync"
)

// ChecklistProcess starts a process that collects the reports of the processes of the test-bed,
// and checks the test was complete.
// It creates and returns with two channels, one channel for collecting test results,
// and another one that it will close, when the test is finished.
func ChecklistProcess(expected []string, done chan bool, wg *sync.WaitGroup, logger *logrus.Logger) (chan string, chan bool) {
	testCompletedCh := make(chan bool)
	reportCh := make(chan string)
	reported := make(map[string]bool)

	wg.Add(1)
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
				if checkReported(reported, expected) {
					return
				}
			}
		}
	}()

	logger.Infof("Checklist Process started.")
	return reportCh, testCompletedCh
}

// checkReported compares reported to expected,
// and returns true if all expected item exists in reported array, otherwise returns false.
func checkReported(reported map[string]bool, expected []string) bool {
	result := true
	for _, c := range expected {
		if _, exist := reported[c]; !exist {
			result = false
			break
		}
	}

	return result
}
