package main

import (
	"github.com/robfig/cron"
	"github.com/tombenke/axon-go-common/actor/node"
	"github.com/tombenke/axon-go-common/log"
	"sync"
)

func startCron(n node.Node, CronDef string, nodeWg *sync.WaitGroup, cronDoneCh chan bool) {

	inputs := n.NewInputs()

	cronFunc := func() {
		log.Logger.Infof("Cron emits the next trigger")
		n.Next(inputs)
	}

	nodeWg.Add(1)
	go func() {
		log.Logger.Debugf("Cron started.")
		defer log.Logger.Debugf("Cron stopped.")
		defer nodeWg.Done()

		c := cron.New()
		defer c.Stop()

		if err := c.AddFunc(CronDef, cronFunc); err != nil {
			panic(err)
		}

		c.Start()

		<-cronDoneCh
		log.Logger.Debugf("Cron is shutting down")
	}()
}
