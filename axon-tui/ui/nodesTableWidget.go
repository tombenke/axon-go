package ui

import (
	ui "github.com/gizak/termui/v3"
	"github.com/tombenke/axon-go-common/log"
	"github.com/tombenke/axon-go-common/msgs/orchestra"
	"github.com/tombenke/axon-go/axon-tui/epn"
)

type NodesTableWidget struct {
	*Table
	visible     bool
	inFocus     bool
	epnStatusCh chan orchestra.EPNStatus
	eventsCh    chan ui.Event
}

func NewNodesTableWidget(epnStatus *epn.Status, eventsHub *EventsHub) *NodesTableWidget {
	epnStatusCh := epnStatus.Subscribe()
	nodes := &NodesTableWidget{
		Table:       NewTable(),
		visible:     true,
		epnStatusCh: epnStatusCh,
		eventsCh:    eventsHub.Subscribe(),
	}

	nodes.Title = "Actor Nodes"
	nodes.Header = []string{"Name", "Type", "Synchronized", "Valami"}
	nodes.ColGap = 2
	nodes.PadLeft = 1
	nodes.ShowCursor = true
	nodes.CursorColor = ui.ColorBlue
	nodes.UniqueCol = 0
	nodes.ColResizer = func() {
		tableWidth := nodes.Table.Inner.Dx()
		nameColWidth := (tableWidth - 5) / 3
		nodes.ColWidths = []int{
			nameColWidth, nameColWidth, nameColWidth, 5,
		}
	}

	nodes.Rows = make([][]string, 0)

	go nodes.controller()

	return nodes
}

func (nodes *NodesTableWidget) controller() {
	for {
		select {
		case epnStatus := <-nodes.epnStatusCh:
			nodes.update(epnStatus)

		case e := <-nodes.eventsCh:
			switch e.ID {
			case "<Left>", "<Up>":
				nodes.ScrollUp()
				ui.Render(nodes)

			case "<Right>", "<Down>":
				nodes.ScrollDown()
				ui.Render(nodes)
			}
		}
	}
}

func (nodes *NodesTableWidget) update(epnStatus orchestra.EPNStatus) {
	numActors := len(epnStatus.Body.Actors)
	if numActors > 0 {
		nodes.Rows = make([][]string, numActors)
		for i, actor := range epnStatus.Body.Actors {
			nodes.Rows[i] = make([]string, 3)
			nodes.Rows[i][0] = actor.Node.Name
			nodes.Rows[i][1] = actor.Node.Type
			if actor.Node.Synchronization {
				nodes.Rows[i][2] = "Y"
			} else {
				nodes.Rows[i][2] = "-"
			}
		}
		log.Logger.Debugf("node.Table.Rows: %v", nodes.Rows)
		ui.Render(nodes)
	}
}
