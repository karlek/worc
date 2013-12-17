// Package draw implements functions to draw to the screen.
package draw

import (
	"github.com/karlek/worc/coord"
	"github.com/karlek/worc/screen"
	"github.com/nsf/termbox-go"
)

// Drawable asserts that an object can be drawn on screen.
type Drawable interface {
	Graphic() termbox.Cell
}

// DrawXY draws a drawable object to the screen.
func DrawXY(x, y int, d Drawable, scr screen.Screen) {
	// Check if the coordinate exists on the plane.
	// Since the screen isn't always located at (0, 0) we have to take
	// the offsets into account.
	c := coord.Coord{x + scr.XOffset, y + scr.YOffset}
	p := coord.Plane{scr.Width, scr.Height}
	if !p.Contains(c) {
		return
	}
	termbox.SetCell(x, y, d.Graphic().Ch, d.Graphic().Fg, d.Graphic().Bg)
	termbox.Flush()
}
