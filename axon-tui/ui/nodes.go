package ui

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/tombenke/axon-go-common/log"
	"github.com/tombenke/axon-go-common/msgs/orchestra"
	"github.com/tombenke/axon-go/axon-tui/epn"
)

type NodesWidget struct {
	*widgets.List
	data        []string
	visible     bool
	epnStatusCh chan orchestra.EPNStatus
}

func NewNodesWidget(epnStatus *epn.Status) *NodesWidget {
	epnStatusCh := epnStatus.Subscribe()
	nodes := &NodesWidget{
		List:        widgets.NewList(),
		visible:     true,
		data:        make([]string, 1),
		epnStatusCh: epnStatusCh,
	}

	nodes.List.Title = "Actor Nodes"
	nodes.List.Rows = []string{}

	go func() {
		for {
			log.Logger.Debugf("wait for next status...%v", nodes.epnStatusCh)
			select {
			case status := <-nodes.epnStatusCh:
				actors := []string{}
				for _, actor := range status.Body.Actors {
					actors = append(actors, actor.Name)
				}
				nodes.List.Rows = actors
				log.Logger.Debugf("node.List.Rows: %v", nodes.List.Rows)
				ui.Render(nodes)
			}
		}
	}()

	return nodes
}
