package ui

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/tombenke/axon-go-common/log"
	"github.com/tombenke/axon-go/axon-tui/epn"
	"sync"
	"syscall"
)

type UI struct {
	events <-chan ui.Event
	header *widgets.Paragraph
	grid   *ui.Grid
	nodes  *NodesWidget
}

func New(epnStatus *epn.Status) UI {
	// Initialize the UI
	if err := ui.Init(); err != nil {
		log.Logger.Fatalf("failed to initialize termui: %v", err)
	}

	u := UI{
		nodes: NewNodesWidget(epnStatus),
	}
	log.Logger.Debugf("UI: %v", u)

	// Setup the main screen
	u.setupScreen()

	return u
}

func (u *UI) Start(appWg *sync.WaitGroup) {
	appWg.Add(1)
	defer appWg.Done()

	// Start the polling and processing of UI events
	u.events = ui.PollEvents()

	for e := range u.events {
		switch e.ID {
		case "q", "<C-c>":
			if err := syscall.Kill(syscall.Getpid(), syscall.SIGTERM); err != nil {
				panic(err)
			}
			return

		case "<Resize>":
			payload := e.Payload.(ui.Resize)
			u.header.SetRect(0, 0, payload.Width, 3)
			u.grid.SetRect(0, 3, payload.Width, payload.Height)
			ui.Clear()
			ui.Render(u.header, u.grid)
		}
	}
}

func (u *UI) Shutdown() {
	ui.Close()
}

func (u *UI) setupScreen() {

	termWidth, termHeight := ui.TerminalDimensions()

	u.header = widgets.NewParagraph()
	u.header.SetRect(0, 0, termWidth, 3)
	u.header.Text = "Overview of the Actor Nodes and Channels of the Event Processing Network"

	right := widgets.NewParagraph()
	right.Title = "Node Details"

	u.grid = ui.NewGrid()
	u.grid.SetRect(0, 3, termWidth, termHeight)

	u.grid.Set(
		ui.NewRow(1.0,
			ui.NewCol(1.0/2, u.nodes),
			ui.NewCol(1.0/2, right),
		),
	)

	ui.Clear()
	ui.Render(u.header, u.grid)
}
