package ui

type Controller interface {
	controller()
	UseEvents()
	Show()
	Hide()
}

type Control struct {
	Visible   bool
	UseEvents bool
}
