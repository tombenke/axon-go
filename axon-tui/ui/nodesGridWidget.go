package ui

import (
	"fmt"
	ui "github.com/gizak/termui/v3"
	"github.com/tombenke/axon-go/axon-tui/epn"
)

type NodesGridWidget struct {
	*ui.Grid
	nodesTable  *NodesTableWidget
	nodeDetails *NodeDetailsWidget
	eventsCh    chan ui.Event
	visible     bool
	inFocus     bool
	Height      int
}

func NewNodesGridWidget(width, height, headerHeight int, epnStatus *epn.Status, eventsHub *EventsHub) *NodesGridWidget {
	nodesGrid := &NodesGridWidget{
		Grid:        ui.NewGrid(),
		nodesTable:  NewNodesTableWidget(epnStatus),
		nodeDetails: NewNodeDetailsWidget(),
		eventsCh:    eventsHub.Subscribe(),
		visible:     true,
		inFocus:     false,
		Height:      height,
	}

	nodesGrid.SetRect(0, headerHeight, width, height)

	nodesGrid.Set(
		ui.NewRow(1.0,
			ui.NewCol(1.0/2, nodesGrid.nodesTable),
			ui.NewCol(1.0/2, nodesGrid.nodeDetails),
		),
	)

	go nodesGrid.controller()

	return nodesGrid
}

// controller is the event handler and controller of the widget
func (n *NodesGridWidget) controller() {
	for e := range n.eventsCh {
		switch e.ID {
		case "u", "<Up>":
			fmt.Println("UP")

		case "d", "<Down>":
			fmt.Println("DOWN")
		}
	}
}
