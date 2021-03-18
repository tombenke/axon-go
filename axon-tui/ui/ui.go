package ui

import (
	ui "github.com/gizak/termui/v3"
	"github.com/tombenke/axon-go-common/log"
	"github.com/tombenke/axon-go/axon-tui/epn"
	"sync"
)

// UI represents the UI part of the application, including the widgets and the event handling logic
type UI struct {
	eventsHub *EventsHub
	main      *MainWidget
}

// New create the UI part of the application
func New(epnStatus *epn.Status) UI {
	// Initialize the UI
	if err := ui.Init(); err != nil {
		log.Logger.Fatalf("failed to initialize termui: %v", err)
	}

	// Start the polling and processing of UI events
	eventsHub := NewEventsHub()

	// Create the main app widget
	u := UI{
		eventsHub: eventsHub,
		main:      NewMainWidget(epnStatus, eventsHub),
	}

	return u
}

// Start create the main widget and start the UI part of the application
func (u *UI) Start(appWg *sync.WaitGroup) {
	u.eventsHub.Start()

	ui.Clear()
	ui.Render(u.main)
}

// Shutdown shuts down the UI part of the application
func (u *UI) Shutdown() {
	ui.Close()
}
