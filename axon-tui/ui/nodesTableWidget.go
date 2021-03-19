package ui

import (
	ui "github.com/gizak/termui/v3"
	"github.com/tombenke/axon-go-common/log"
	"github.com/tombenke/axon-go-common/msgs/orchestra"
	"sort"
)

type NodesTableWidget struct {
	*Table
	Control
	eventsCh chan ui.Event
}

func NewNodesTableWidget(eventsHub *EventsHub) *NodesTableWidget {
	nodes := &NodesTableWidget{
		Table:    NewTable(),
		eventsCh: eventsHub.Subscribe(),
	}

	nodes.Title = "Actor Nodes"
	nodes.Header = []string{"Name", "Type", "Synchronized", "Valami"}
	nodes.ColGap = 2
	nodes.PadLeft = 1
	nodes.ShowCursor = true
	nodes.ShowLocation = true
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

// Draw draws the content of the widget into the `Buffer` object provided as an argument
func (nodes *NodesTableWidget) Draw(buf *ui.Buffer) {
	if nodes.Visible {
		nodes.Table.Draw(buf)
	}
}

func (nodes *NodesTableWidget) controller() {
	for e := range nodes.eventsCh {
		if nodes.UseEvents {
			switch e.ID {
			case "<Left>", "<Up>":
				nodes.ScrollUp()

			case "<Right>", "<Down>":
				nodes.ScrollDown()

			case "<PageUp>":
				nodes.ScrollPageUp()

			case "<PageDown>":
				nodes.ScrollPageDown()

			case "<Home>":
				nodes.ScrollTop()

			case "<End>":
				nodes.ScrollBottom()

			case "<MouseLeft>":
				nodes.HandleClick(e.Payload.(ui.Mouse).X, e.Payload.(ui.Mouse).Y)
			}
			ui.Render(nodes)
		}
	}
}

func (nodes *NodesTableWidget) update(actors []orchestra.Actor) {
	numActors := len(actors)
	nodes.Rows = make([][]string, numActors)
	sort.Slice(actors, func(i, j int) bool {
		return actors[i].Node.Name < actors[j].Node.Name
	})

	if numActors > 0 {
		for i, actor := range actors {
			nodes.Rows[i] = make([]string, 3)
			nodes.Rows[i][0] = actor.Node.Name
			nodes.Rows[i][1] = actor.Node.Type
			if actor.Node.Synchronization {
				nodes.Rows[i][2] = "Y"
			} else {
				nodes.Rows[i][2] = "-"
			}
		}
	}
	ui.Render(nodes)
	log.Logger.Debugf("node.Table.Rows: %v, SelectedRow: %d", nodes.Rows, nodes.SelectedRow)
}
