///
package draw

import (
	"github.com/karlek/torc/coord"
	"github.com/karlek/torc/screen"
	"github.com/nsf/termbox-go"
)

///
type Drawable interface {
	Graphic() termbox.Cell
}

/// return error if x,y coordinate doesn't exist?
///
func DrawXY(x, y int, d Drawable, scr screen.Screen) {
	// Check if the coordinate exists on the plane.
	// Since the screen isn't always located at (0, 0) we have to take
	// the offsets into account.
	c := coord.Coord{x + scr.XOffset, y + scr.YOffset}
	p := coord.Plane{scr.Width, scr.Height}
	if !p.Exists(c) {
		return
	}
	termbox.SetCell(x, y, d.Graphic().Ch, d.Graphic().Fg, d.Graphic().Bg)
	termbox.Flush()
}
