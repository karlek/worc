///
package object

import (
	"github.com/nsf/termbox-go"
)

// Object is something that is drawed on AreaScreen ontop of an area.
type Object struct {
	X        int          ///
	Y        int          ///
	G        termbox.Cell ///
	Walkable bool         ///
}

///
func (o *Object) Graphic() termbox.Cell {
	return o.G
}

///
func (o *Object) IsWalkable() bool {
	return o.Walkable
}
