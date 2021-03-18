package ui

import (
	ui "github.com/gizak/termui/v3"
	"github.com/tombenke/axon-go-common/log"
)

// EventsHub represents an observer/dispatcher hub of the UI keyboard and mouse events
// This hub listens for the UI events, then forwards them to its subscribers, e.g. the widgets of the application
type EventsHub struct {
	eventsCh    <-chan ui.Event
	subscribers []chan ui.Event
}

// NewEventsHub create a new EventsHub instance
func NewEventsHub() *EventsHub {
	return &EventsHub{
		eventsCh:    ui.PollEvents(),
		subscribers: make([]chan ui.Event, 0),
	}
}

// Start starts the observing/forwarding of the UI events
func (e *EventsHub) Start() {
	go func() {
		for event := range e.eventsCh {
			log.Logger.Debugf("EventsHub received event: %v", event)
			for _, subscriber := range e.subscribers {
				log.Logger.Debugf("EventsHub sends status to subscriber: %v > %v", event, subscriber)
				subscriber <- event
			}
		}
	}()

	log.Logger.Debugf("EventsHub is started")
}

// Shutdown shuts down the EventsHub instance
func (e *EventsHub) Shutdown() {
	log.Logger.Debugf("EventsHub is shutting down")
}

// Subscribe returns a channel through which the EventsHub will send UI events to the subscriber
func (e *EventsHub) Subscribe() chan ui.Event {
	subscriber := make(chan ui.Event)
	(*e).subscribers = append(e.subscribers, subscriber)

	log.Logger.Debugf("EventsHub subscribers: %v", e.subscribers)
	return subscriber
}
