package ui

import (
	"github.com/gizak/termui/v3/widgets"
)

type NodeDetailsWidget struct {
	*widgets.Paragraph
	visible bool
	inFocus bool
}

func NewNodeDetailsWidget() *NodeDetailsWidget {
	nodeDetails := &NodeDetailsWidget{
		Paragraph: widgets.NewParagraph(),
		visible:   true,
		inFocus:   false,
	}

	nodeDetails.Paragraph.Title = "Node Details"

	return nodeDetails
}
