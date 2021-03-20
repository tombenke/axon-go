package ui

import (
	"fmt"
	"image"
	"log"
	"strings"

	. "github.com/gizak/termui/v3"
	ui "github.com/gizak/termui/v3"
)

// Table Represents a table base widget that also acts as a select list
type Table struct {
	*Block

	// Header the header of the table
	Header []string

	// Rows A two-dimensional array of data rows/cells
	Rows [][]string

	// ColWidths hold the width values of the columns of the table
	ColWidths []int

	// ColGap definee the gap among the columns given in spaces
	ColGap int

	// PadLeft define the padding before the columns
	PadLeft int

	// ShowCursor use different color for the current position of selection
	ShowCursor bool

	// CursorColor define the color of the selected row
	CursorColor Color

	// ShowLocation shows the index of the selected row in the visible range of rows
	ShowLocation bool

	// UniqueCol define the column used to uniquely identify each table row
	UniqueCol int

	// SelectedItem used to keep the cursor on the correct item if the data changes
	SelectedItem string

	// SelectedRow holds the index to the currently selected row
	SelectedRow int

	// TopRow used to indicate where in the table we are scrolled at
	TopRow int

	// ColResizer is a helper function that dinamically sets the width of columns in case of resize
	ColResizer func()
}

// NewTable returns a new Table instance
func NewTable() *Table {
	return &Table{
		Block:       NewBlock(),
		SelectedRow: 0,
		TopRow:      0,
		UniqueCol:   0,
		ColResizer:  func() {},
		CursorColor: ui.Color(colorSchema.Cursor),
	}
}

// Draw draws the content of the widget into the `Buffer` object provided as an argument
func (self *Table) Draw(buf *Buffer) {
	self.Block.Draw(buf)

	if self.ShowLocation {
		self.drawLocation(buf)
	}

	self.ColResizer()

	// finds exact column starting position
	colXPos := []int{}
	cur := 1 + self.PadLeft
	for _, w := range self.ColWidths {
		colXPos = append(colXPos, cur)
		cur += w
		cur += self.ColGap
	}

	// prints header
	for i, h := range self.Header {
		width := self.ColWidths[i]
		if width == 0 {
			continue
		}
		// don't render column if it doesn't fit in widget
		if width > (self.Inner.Dx()-colXPos[i])+1 {
			continue
		}
		buf.SetString(
			h,
			NewStyle(Theme.Default.Fg, ColorClear, ModifierBold),
			image.Pt(self.Inner.Min.X+colXPos[i]-1, self.Inner.Min.Y),
		)
	}

	if self.TopRow < 0 {
		log.Printf("table widget TopRow value less than 0. TopRow: %v", self.TopRow)
		return
	}

	// prints each row
	for rowNum := self.TopRow; rowNum < self.TopRow+self.Inner.Dy()-1 && rowNum < len(self.Rows); rowNum++ {
		row := self.Rows[rowNum]
		y := (rowNum + 2) - self.TopRow

		// prints cursor
		style := NewStyle(Theme.Default.Fg)
		if self.ShowCursor {
			if (self.SelectedItem == "" && rowNum == self.SelectedRow) || (self.SelectedItem != "" && self.SelectedItem == row[self.UniqueCol]) {
				style.Fg = self.CursorColor
				style.Modifier = ModifierReverse
				for _, width := range self.ColWidths {
					if width == 0 {
						continue
					}
					buf.SetString(
						strings.Repeat(" ", self.Inner.Dx()),
						style,
						image.Pt(self.Inner.Min.X, self.Inner.Min.Y+y-1),
					)
				}
				self.SelectedItem = row[self.UniqueCol]
				self.SelectedRow = rowNum
			}
		}

		// prints each col of the row
		for i, width := range self.ColWidths {
			if width == 0 {
				continue
			}
			// don't render column if width is greater than distance to end of widget
			if width > (self.Inner.Dx()-colXPos[i])+1 {
				continue
			}
			r := TrimString(row[i], width)
			buf.SetString(
				r,
				style,
				image.Pt(self.Inner.Min.X+colXPos[i]-1, self.Inner.Min.Y+y-1),
			)
		}
	}
}

func (self *Table) drawLocation(buf *Buffer) {
	total := len(self.Rows)
	topRow := self.TopRow + 1
	bottomRow := self.TopRow + self.Inner.Dy() - 1
	if bottomRow > total {
		bottomRow = total
	}

	loc := fmt.Sprintf(" %d - %d of %d ", topRow, bottomRow, total)

	width := len(loc)
	buf.SetString(loc, self.TitleStyle, image.Pt(self.Max.X-width-2, self.Min.Y))
}

// Scrolling ///////////////////////////////////////////////////////////////////

// CalcPos is used to calculate the cursor position and the current view into the table.
func (self *Table) CalcPos() {
	self.SelectedItem = ""

	if self.SelectedRow < 0 {
		self.SelectedRow = 0
	}
	if self.SelectedRow < self.TopRow {
		self.TopRow = self.SelectedRow
	}

	if self.SelectedRow > len(self.Rows)-1 {
		self.SelectedRow = len(self.Rows) - 1
	}
	if self.SelectedRow > self.TopRow+(self.Inner.Dy()-2) {
		self.TopRow = self.SelectedRow - (self.Inner.Dy() - 2)
	}
}

func (self *Table) ScrollUp() {
	self.SelectedRow--
	self.CalcPos()
}

func (self *Table) ScrollDown() {
	self.SelectedRow++
	self.CalcPos()
}

func (self *Table) ScrollTop() {
	self.SelectedRow = 0
	self.CalcPos()
}

func (self *Table) ScrollBottom() {
	self.SelectedRow = len(self.Rows) - 1
	self.CalcPos()
}

func (self *Table) ScrollHalfPageUp() {
	self.SelectedRow = self.SelectedRow - (self.Inner.Dy()-2)/2
	self.CalcPos()
}

func (self *Table) ScrollHalfPageDown() {
	self.SelectedRow = self.SelectedRow + (self.Inner.Dy()-2)/2
	self.CalcPos()
}

func (self *Table) ScrollPageUp() {
	self.SelectedRow -= (self.Inner.Dy() - 2)
	self.CalcPos()
}

func (self *Table) ScrollPageDown() {
	self.SelectedRow += (self.Inner.Dy() - 2)
	self.CalcPos()
}

func (self *Table) HandleClick(x, y int) {
	x = x - self.Min.X
	y = y - self.Min.Y
	if (x > 0 && x <= self.Inner.Dx()) && (y > 0 && y <= self.Inner.Dy()) {
		self.SelectedRow = (self.TopRow + y) - 2
		self.CalcPos()
	}
}
