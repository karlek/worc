// Package model is something which can be placed on an area and can be drawn.
package model

import (
	"github.com/nsf/termbox-go"
)

type Modelable interface {
	X() int
	Y() int
	NewX(int)
	NewY(int)
	SetGraphics(termbox.Cell)
	SetPathable(bool)
	IsPathable() bool
}

// Model is something that is drawed on AreaScreen ontop of an area.
type Model struct {
	Modelable
	Xval     int          // X is x coordinate.
	Yval     int          // Y is y coordinate.
	G        termbox.Cell // G is graphics.
	Pathable bool         // Pathable is a boolean flag if other objects can be placed on top of this model.
}

// Graphic is needed for draw.Drawable.
func (m Model) Graphic() termbox.Cell {
	return m.G
}

// IsPathable is needed for area.Pathable.
func (m Model) IsPathable() bool {
	return m.Pathable
}

// X returns the x value of the current coordinate.
func (m Model) X() int {
	return m.Xval
}

// Y returns the y value of the current coordinate.
func (m Model) Y() int {
	return m.Yval
}

// NewX sets a new x value for the coordinate.
func (m *Model) NewX(x int) {
	m.Xval = x
}

// NewY sets a new y value for the coordinate.
func (m *Model) NewY(y int) {
	m.Yval = y
}

// SetGraphics sets graphics for the model.
func (m *Model) SetGraphics(g termbox.Cell) {
	m.G = g
}

// SetPathable sets pathability for the model.
func (m *Model) SetPathable(pathable bool) {
	m.Pathable = pathable
}
