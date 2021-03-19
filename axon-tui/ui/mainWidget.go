package ui

import (
	ui "github.com/gizak/termui/v3"
	"github.com/tombenke/axon-go/axon-tui/epn"
	"syscall"
)

// MainWidget represents the main widget of the application
type MainWidget struct {
	*ui.Block
	Control
	header    *HeaderWidget
	nodesGrid *NodesGridWidget
	eventsCh  chan ui.Event
}

// Draw draws the content of the widget into the `Buffer` object provided as an argument
func (m *MainWidget) Draw(buf *ui.Buffer) {
	if m.Visible {
		ui.Clear()
		m.header.Visible = true
		m.nodesGrid.Visible = true
		ui.Render(m.header, m.nodesGrid)
	} else {
		m.header.Visible = false
		m.nodesGrid.Visible = false
	}
}

// NewMainWidget create a new instance of the main widget of the application
func NewMainWidget(epnStatus *epn.Status, eventsHub *EventsHub) *MainWidget {

	termWidth, termHeight := ui.TerminalDimensions()
	headerHeight := 3

	main := &MainWidget{
		Block:     ui.NewBlock(),
		header:    NewHeaderWidget(termWidth, headerHeight),
		nodesGrid: NewNodesGridWidget(termWidth, termHeight, headerHeight, epnStatus, eventsHub),
		eventsCh:  eventsHub.Subscribe(),
	}

	main.Border = false

	go main.controller()

	return main
}

// controller is the event handler and controller of the widget
func (m *MainWidget) controller() {
	for e := range m.eventsCh {
		if m.UseEvents {
			switch e.ID {
			case "q", "<Escape>", "<C-c>":
				if err := syscall.Kill(syscall.Getpid(), syscall.SIGTERM); err != nil {
					panic(err)
				}
				return

			case "<Resize>":
				m.resize(e.Payload.(ui.Resize))
			}
		}
	}
}

// resize change the size of the widget according to the `size` parameters
func (m *MainWidget) resize(size ui.Resize) {
	m.header.SetRect(0, 0, size.Width, m.header.Height)
	m.nodesGrid.SetRect(0, m.header.Height, size.Width, size.Height)

	ui.Render(m)
}
