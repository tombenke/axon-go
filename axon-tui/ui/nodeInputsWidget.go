package ui

import (
	ui "github.com/gizak/termui/v3"
	"github.com/tombenke/axon-go-common/log"
	"github.com/tombenke/axon-go-common/msgs/orchestra"
	"sort"
)

type NodeInputsWidget struct {
	*Table
	Control
	eventsCh chan ui.Event
}

func NewNodeInputsWidget(eventsHub *EventsHub) *NodeInputsWidget {
	nodes := &NodeInputsWidget{
		Table:    NewTable(),
		eventsCh: eventsHub.Subscribe(),
	}

	nodes.Title = "Node Inputs"
	nodes.Header = []string{"Name", "Type", "Representation", "Channel"}
	nodes.ColGap = 2
	nodes.PadLeft = 1
	nodes.ShowCursor = false // true
	nodes.ShowLocation = true
	nodes.CursorColor = ui.ColorBlue
	nodes.UniqueCol = 0
	nodes.ColResizer = func() {
		tableWidth := nodes.Table.Inner.Dx()
		cw := tableWidth / 10
		nodes.ColWidths = []int{
			2 * cw, 2 * cw, 2 * cw, 3 * cw,
		}
	}

	nodes.Rows = make([][]string, 0)

	go nodes.controller()

	return nodes
}

// Draw draws the content of the widget into the `Buffer` object provided as an argument
func (nodes *NodeInputsWidget) Draw(buf *ui.Buffer) {
	if nodes.Visible {
		nodes.Table.Draw(buf)
	}
}

func (nodes *NodeInputsWidget) controller() {
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

func (nodes *NodeInputsWidget) update(ports []orchestra.Port) {
	numPorts := len(ports)
	nodes.Rows = make([][]string, numPorts)
	sort.Slice(ports, func(i, j int) bool {
		return ports[i].Name < ports[j].Name
	})

	if numPorts > 0 {
		for i, port := range ports {
			nodes.Rows[i] = make([]string, 4)
			nodes.Rows[i][0] = port.Name
			nodes.Rows[i][1] = port.Type
			nodes.Rows[i][2] = port.Representation
			nodes.Rows[i][3] = port.Channel.Name
		}
	}
	ui.Render(nodes)
	log.Logger.Debugf("nodeInputs.Table.Rows: %v, SelectedRow: %d", nodes.Rows, nodes.SelectedRow)
}
