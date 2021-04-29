package main

import (
	"github.com/tombenke/axon-go-common/actor/node"
	"github.com/tombenke/axon-go-common/log"
	"github.com/tombenke/axon-go-common/msgs/base"
	"os"
	"sync"
	"syscall"
	"time"
)

func startInjector(n node.Node, config Config, appWg *sync.WaitGroup, injectDoneCh chan interface{}, sigsCh chan os.Signal) {

	inputs := n.NewInputs()

	appWg.Add(1)
	go func() {
		log.Logger.Debugf("Injector started.")
		defer log.Logger.Debugf("Inject stopped.")
		defer appWg.Done()

		delayDuration, err := parseDelay(config.Delay)
		if err != nil {
			log.Logger.Errorf("%s", err)
		} else {
			for r := 0; r < config.Repeat; r++ {
				log.Logger.Infof("Injector emits the next message")
				newMessage := base.NewBytesMessage([]byte(config.Message))
				(*inputs).SetMessage("inject", newMessage)
				n.Next(inputs)
				if r < config.Repeat-1 {
					time.Sleep(delayDuration)
					log.Logger.Debugf("%d/%d iteration", config.Repeat, r)
				}
			}
		}

		sigsCh <- syscall.SIGTERM
		<-injectDoneCh
		log.Logger.Debugf("Injector is shutting down")
	}()
}
