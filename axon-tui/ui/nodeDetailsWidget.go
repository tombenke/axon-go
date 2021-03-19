package ui

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/tombenke/axon-go-common/msgs/orchestra"
	"gopkg.in/yaml.v2"
)

type NodeDetailsWidget struct {
	*widgets.Paragraph
	Control
}

func NewNodeDetailsWidget() *NodeDetailsWidget {
	nodeDetails := &NodeDetailsWidget{
		Paragraph: widgets.NewParagraph(),
	}

	nodeDetails.Paragraph.Title = "Node Details"

	return nodeDetails
}

// Draw draws the content of the widget into the `Buffer` object provided as an argument
func (n *NodeDetailsWidget) Draw(buf *ui.Buffer) {
	if n.Visible {
		n.Paragraph.Draw(buf)
	} else {
		rect := n.Paragraph.GetRect()
		emptyCell := ui.NewCell(' ')
		buf.Fill(emptyCell, rect)
	}
}

func (node *NodeDetailsWidget) update(actor orchestra.Actor) {
	text, err := yaml.Marshal(&actor.Node)
	if err != nil {
		panic(err)
	}
	node.Text = string(text)
	ui.Render(node)
}
