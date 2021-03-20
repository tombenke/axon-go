package ui

import (
	ui "github.com/gizak/termui/v3"
	"github.com/tombenke/axon-go-common/log"
	"github.com/tombenke/axon-go/axon-tui/epn"
	"sync"
)

type ColorSchema struct {
	Fg          int
	Bg          int
	BorderLabel int
	BorderLine  int
	Cursor      int
}

type ColorSchemas map[string]ColorSchema

var (
	colorSchemas = ColorSchemas{
		"default": ColorSchema{
			Fg:          7,
			Bg:          -1,
			BorderLabel: 7,
			BorderLine:  6,
			Cursor:      4,
		},
		"monokai": ColorSchema{
			Fg:          249,
			Bg:          -1,
			BorderLabel: 249,
			BorderLine:  239,
			Cursor:      197,
		},
		"solarized": ColorSchema{
			Fg:          250,
			Bg:          -1,
			BorderLabel: 250,
			BorderLine:  37,
			Cursor:      136,
		},
		"vice": ColorSchema{
			Fg:          231,
			Bg:          -1,
			BorderLabel: 123,
			BorderLine:  102,
			Cursor:      159,
		},
		"axon": ColorSchema{
			Fg:          255,
			Bg:          -1,
			BorderLabel: 37,
			BorderLine:  37,
			Cursor:      159,
		},
	}
	colorSchema ColorSchema
)

func init() {
	colorSchema = colorSchemas["default"]
}

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

	// Set the Theme
	colorSchema = colorSchemas["monokai"]
	ui.Theme.Default = ui.NewStyle(ui.Color(colorSchema.Fg), ui.Color(colorSchema.Bg))
	ui.Theme.Block.Title = ui.NewStyle(ui.Color(colorSchema.BorderLabel), ui.Color(colorSchema.Bg))
	ui.Theme.Block.Border = ui.NewStyle(ui.Color(colorSchema.BorderLine), ui.Color(colorSchema.Bg))

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

	// Make the main widget visible, and start processing the UI event
	u.main.Visible = true
	u.main.UseEvents = true

	ui.Clear()
	ui.Render(u.main)
}

// Shutdown shuts down the UI part of the application
func (u *UI) Shutdown() {
	ui.Close()
}
