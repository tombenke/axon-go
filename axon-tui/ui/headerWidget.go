package ui

import (
	"github.com/gizak/termui/v3/widgets"
)

type HeaderWidget struct {
	*widgets.Paragraph
	visible bool
	inFocus bool
	Height  int
}

func NewHeaderWidget(width, height int) *HeaderWidget {
	header := &HeaderWidget{
		Paragraph: widgets.NewParagraph(),
		visible:   true,
		inFocus:   false,
		Height:    height,
	}

	header.Paragraph.Title = "Node Details"
	header.SetRect(0, 0, width, header.Height)
	header.Text = "Overview of the Actor Nodes and Channels of the Event Processing Network"

	return header
}
