package ui

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/tombenke/axon-go-common/log"
	"github.com/tombenke/axon-go-common/msgs/orchestra"
	"github.com/tombenke/axon-go/axon-tui/epn"
)

type NodesTableWidget struct {
	*widgets.Table
	visible     bool
	inFocus     bool
	epnStatusCh chan orchestra.EPNStatus
}

func NewNodesTableWidget(epnStatus *epn.Status) *NodesTableWidget {
	epnStatusCh := epnStatus.Subscribe()
	nodes := &NodesTableWidget{
		Table:       widgets.NewTable(),
		visible:     true,
		epnStatusCh: epnStatusCh,
	}

	nodes.Table.Title = "Actor Nodes"
	nodes.Table.ColumnWidths = []int{30, 30, 4}
	nodes.Table.RowSeparator = false
	nodes.Table.TextStyle = ui.Style{Fg: ui.ColorYellow, Bg: ui.ColorBlue}
	nodes.Table.FillRow = true
	rows := make([][]string, 1)
	nodes.Table.Rows = rows
	//nodes.Table.Header = []string{"Name", "Type", "Synchronized"}
	//nodes.Table.ColGap = 2

	go func() {
		for {
			log.Logger.Debugf("wait for next status...%v", nodes.epnStatusCh)
			select {
			case status := <-nodes.epnStatusCh:
				numActors := len(status.Body.Actors)
				if numActors > 0 {
					nodes.Table.Rows = make([][]string, numActors)
					for i, actor := range status.Body.Actors {
						nodes.Table.Rows[i] = make([]string, 3)
						nodes.Table.Rows[i][0] = actor.Node.Name
						nodes.Table.Rows[i][1] = actor.Node.Type
						if actor.Node.Synchronization {
							nodes.Table.Rows[i][2] = "Y"
						} else {
							nodes.Table.Rows[i][2] = " "
						}
					}
					log.Logger.Debugf("node.Table.Rows: %v", nodes.Table.Rows)
					ui.Render(nodes)
				}
			}
		}
	}()

	return nodes
}
