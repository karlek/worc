// Package object is something which can be placed on an area and can be drawn.
package object

import (
	"github.com/nsf/termbox-go"
)

// Object is something that is drawed on AreaScreen ontop of an area.
type Object struct {
	Xval      int          // X is x coordinate.
	Yval      int          // Y is y coordinate.
	G         termbox.Cell // G is graphics.
	Stackable bool         // Stackable is a boolean flag if objects can be placed on top of this object.
}

// Graphic is needed for draw.Drawable.
func (o Object) Graphic() termbox.Cell {
	return o.G
}

// IsStackable is needed for area.Stackable.
func (o Object) IsStackable() bool {
	return o.Stackable
}

// X returns the x value of the current coordinate.
func (o Object) X() int {
	return o.Xval
}

// Y returns the y value of the current coordinate.
func (o Object) Y() int {
	return o.Yval
}

// NewX sets a new x value for the coordinate.
func (o *Object) NewX(x int) {
	o.Xval = x
}

// NewY sets a new y value for the coordinate.
func (o *Object) NewY(y int) {
	o.Yval = y
}
