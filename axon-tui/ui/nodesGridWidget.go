package ui

import (
	ui "github.com/gizak/termui/v3"
	"github.com/tombenke/axon-go-common/msgs/orchestra"
	"github.com/tombenke/axon-go/axon-tui/epn"
)

type NodesGridWidget struct {
	*ui.Grid
	Control
	nodesTable  *NodesTableWidget
	nodeDetails *NodeDetailsWidget
	nodeInputs  *NodeInputsWidget
	nodeOutputs *NodeOutputsWidget
	epnStatusCh chan orchestra.EPNStatus
	eventsCh    chan ui.Event
	Height      int
}

func NewNodesGridWidget(width, height, headerHeight int, epnStatus *epn.Status, eventsHub *EventsHub) *NodesGridWidget {
	epnStatusCh := epnStatus.Subscribe()
	nodesGrid := &NodesGridWidget{
		Grid:        ui.NewGrid(),
		nodesTable:  NewNodesTableWidget(eventsHub),
		nodeInputs:  NewNodeInputsWidget(eventsHub),
		nodeOutputs: NewNodeOutputsWidget(eventsHub),
		nodeDetails: NewNodeDetailsWidget(),
		epnStatusCh: epnStatusCh,
		eventsCh:    eventsHub.Subscribe(),
		Height:      height,
	}

	nodesGrid.SetRect(0, headerHeight, width, height)

	nodesGrid.Set(
		ui.NewRow(1.0,
			ui.NewCol(1.0/2, nodesGrid.nodesTable),
			ui.NewCol(1.0/2,
				ui.NewRow(0.2, nodesGrid.nodeDetails),
				ui.NewRow(0.4, nodesGrid.nodeInputs),
				ui.NewRow(0.4, nodesGrid.nodeOutputs),
			),
		),
	)

	go nodesGrid.controller()

	return nodesGrid
}

// Draw draws the content of the widget into the `Buffer` object provided as an argument
func (n *NodesGridWidget) Draw(buf *ui.Buffer) {
	if n.Visible {
		n.nodesTable.Visible = true
		n.nodesTable.UseEvents = true
		n.nodeInputs.Visible = true
		n.nodeInputs.UseEvents = false
		n.nodeOutputs.Visible = true
		n.nodeOutputs.UseEvents = false

	} else {
		n.nodesTable.Visible = false
		n.nodesTable.UseEvents = false
		n.nodeInputs.Visible = false
		n.nodeInputs.UseEvents = false
		n.nodeOutputs.Visible = false
		n.nodeOutputs.UseEvents = false
	}
	n.Grid.Draw(buf)
}

// controller is the event handler and controller of the widget
func (n *NodesGridWidget) controller() {
	for {
		select {
		case epnStatus := <-n.epnStatusCh:
			actors := epnStatus.Body.Actors
			n.nodesTable.update(actors)
			n.nodesTable.CalcPos()
			selectedRow := n.nodesTable.SelectedRow
			numActors := len(actors)
			if selectedRow >= 0 && numActors > 0 {
				n.nodeDetails.Visible = true
				n.nodeDetails.update(epnStatus.Body.Actors[selectedRow])
				n.nodeInputs.update(epnStatus.Body.Actors[selectedRow].Node.Ports.Inputs)
				n.nodeOutputs.update(epnStatus.Body.Actors[selectedRow].Node.Ports.Outputs)
			} else {
				n.nodeDetails.Visible = false
			}
			ui.Render(n)

		case e := <-n.eventsCh:
			if n.UseEvents {
				switch e.ID {
				//NOTE: Add control logic here
				}
			}
		}
	}
}
