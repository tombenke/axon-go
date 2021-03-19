package ui

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type HeaderWidget struct {
	*widgets.Paragraph
	Control
	Height int
}

func NewHeaderWidget(width, height int) *HeaderWidget {
	header := &HeaderWidget{
		Paragraph: widgets.NewParagraph(),
		Height:    height,
	}

	header.SetRect(0, 0, width, header.Height)
	header.Text = "Overview of the Actor Nodes and Channels of the Event Processing Network"

	return header
}

// Draw draws the content of the widget into the `Buffer` object provided as an argument
func (w *HeaderWidget) Draw(buf *ui.Buffer) {
	if w.Visible {
		w.Paragraph.Draw(buf)
	}
}
