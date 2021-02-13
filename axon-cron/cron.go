package main

import (
	"github.com/robfig/cron"
	"github.com/tombenke/axon-go/common/actor/node"
	"github.com/tombenke/axon-go/common/log"
	"sync"
)

func startCron(n node.Node, CronDef string, nodeWg *sync.WaitGroup, cronDoneCh chan bool) {

	nodeWg.Add(1)
	go func() {
		log.Logger.Infof("Cron started.")
		defer log.Logger.Infof("Cron stopped.")
		defer nodeWg.Done()

		c := cron.New()
		defer c.Stop()

		inputs := n.NewInputs()

		c.AddFunc(CronDef, func() {
			log.Logger.Infof("Cron emits the next trigger")
			n.Next(inputs)
		})
		c.Start()

		<-cronDoneCh
		log.Logger.Infof("Cron is shutting down")
	}()
}
