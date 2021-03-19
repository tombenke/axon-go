package ui

import (
	"fmt"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/tombenke/axon-go-common/msgs/orchestra"
	//"gopkg.in/yaml.v2"
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
	/*
		text, err := yaml.Marshal(&actor.Node)
		if err != nil {
			panic(err)
		}
	*/
	synchronization := "NO"
	if actor.Node.Synchronization {
		synchronization = "YES"
	}
	respText := fmt.Sprintf("  response time: [%.2f](fg:yellow,mod:bold) ms\n", float32(actor.ResponseTime/1000000))
	nameText := fmt.Sprintf("     actor name: [%s](fg:yellow,mod:bold)\n", actor.Node.Name)
	typeText := fmt.Sprintf("     actor type: [%s](fg:yellow,mod:bold)\n", actor.Node.Type)
	syncText := fmt.Sprintf("synchronization: [%s](fg:yellow,mod:bold)\n", synchronization)
	node.Text = respText + nameText + typeText + syncText
	ui.Render(node)
}
