package components

import (
	"fmt"

	"github.com/gizak/termui"
)

// Input is the definition of and input box
type Input struct {
	Par                   *termui.Par
	Offset                int
	VisibleCursorPosition int
	CursorPosition        int
	CursorFgColor         termui.Attribute
	CursorBgColor         termui.Attribute
}

// CreateInput is the constructor of the Input struct
func CreateInput() *Input {
	input := &Input{
		Par:            termui.NewPar(""),
		Offset:         0,
		CursorPosition: 0,
		CursorBgColor:  termui.ColorBlack,
		CursorFgColor:  termui.ColorWhite,
	}

	input.Par.Height = 3

	return input
}

// Buffer implements interface termui.Bufferer
func (i *Input) Buffer() termui.Buffer {
	buf := i.Par.Buffer()

	// Overflow
	if len(i.Par.Text)+1 > i.Par.InnerBounds().Dx() {
		cs := termui.DefaultTxBuilder.Build(i.Par.Text, i.Par.TextFgColor, i.Par.TextBgColor)

		// i.Offset = len(i.Par.Text) + 1 - i.Par.InnerBounds().Dx() - len(i.Par.Text)
		left := i.CursorPosition - i.VisibleCursorPosition
		right := ((i.Par.InnerBounds().Dx() - 1) - i.VisibleCursorPosition) + i.CursorPosition
		fmt.Printf("left, right: %d,%d", left, right)

		for n, ch := range cs[left:right] {
			buf.Set(
				i.Par.InnerX()+n+1,
				i.Par.Block.InnerY(),
				ch,
			)
		}
	}

	// Set visible cursor
	char := buf.At(i.Par.InnerX()+i.VisibleCursorPosition, i.Par.Block.InnerY())
	buf.Set(
		i.Par.InnerX()+i.VisibleCursorPosition,
		i.Par.Block.InnerY(),
		termui.Cell{Ch: char.Ch, Fg: termui.ColorBlack, Bg: termui.ColorWhite},
	)

	return buf
}

// GetHeight implements interface termui.GridBufferer
func (i *Input) GetHeight() int {
	return i.Par.Block.GetHeight()
}

// SetWidth implements interface termui.GridBufferer
func (i *Input) SetWidth(w int) {
	i.Par.SetWidth(w)
}

// SetX implements interface termui.GridBufferer
func (i *Input) SetX(x int) {
	i.Par.SetX(x)
}

// SetY implements interface termui.GridBufferer
func (i *Input) SetY(y int) {
	i.Par.SetY(y)
}

// Insert will insert a given key at the place of the current CursorPosition
func (i *Input) Insert(key string) {
	// TODO: from what is visible not Text
	// change i.CursorPosition to i.VisibleCursorPosition
	i.Par.Text = i.Par.Text[0:i.CursorPosition] + key + i.Par.Text[i.CursorPosition:]
	i.MoveVisibleCursorRight()
}

// Remove will remove a character at the place of the current CursorPosition
func (i *Input) Remove() {
	if i.CursorPosition > 0 {
		// TODO: from what is visible not Text
		// change i.CursorPosition to i.VisibleCursorPosition
		i.Par.Text = i.Par.Text[0:i.CursorPosition-1] + i.Par.Text[i.CursorPosition:]
		i.MoveVisibleCursorLeft()
	}
}

// MoveCursorRight will increase the current CursorPosition with 1
func (i *Input) MoveCursorRight() {
	if i.CursorPosition < len(i.Par.Text) {
		i.CursorPosition++
	}
}

// MoveCursorLeft will decrease the current CursorPosition with 1
func (i *Input) MoveCursorLeft() {
	if i.CursorPosition > 0 {
		i.CursorPosition--
	}
}

func (i *Input) MoveVisibleCursorRight() {
	// TODO: if (visible) cursor is at the bounds of the textbox redraw it
	if i.VisibleCursorPosition < i.Par.InnerBounds().Dx()-1 && i.VisibleCursorPosition < len(i.Par.Text) {
		i.VisibleCursorPosition++
	}
	i.MoveCursorRight()
}

func (i *Input) MoveVisibleCursorLeft() {
	// TODO: if (visible) cursor is at the bounds of the textbox redraw it
	if i.VisibleCursorPosition > 0 {
		i.VisibleCursorPosition--
	}
	i.MoveCursorLeft()
}

// IsEmpty will return true when the input is empty
func (i *Input) IsEmpty() bool {
	if i.Par.Text == "" {
		return true
	}
	return false
}

// Clear will empty the input and move the cursor to the start position
func (i *Input) Clear() {
	i.Par.Text = ""
	i.CursorPosition = 0
	i.VisibleCursorPosition = 0
	i.Offset = 0
}

// Text returns the text currently in the input
func (i *Input) Text() string {
	return i.Par.Text
}
