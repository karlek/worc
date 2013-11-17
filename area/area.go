// Package area implements functions for moving objects over a 2d plane.
package area

import (
	"github.com/karlek/worc/coord"
	"github.com/karlek/worc/draw"
	"github.com/karlek/worc/object"
	"github.com/karlek/worc/screen"
	"github.com/nsf/termbox-go"
)

// Walkable are drawable and can answer the question if an object can be placed
// on top of it.
type Walkable interface {
	draw.Drawable
	IsWalkable() bool
}

// Area is a collection of terrain and objects placed on top of it.
type Area struct {
	Terrain       [][]Walkable
	Objects       map[coord.Coord]Stack
	Width, Height int
}

type Stack []Walkable

func (s Stack) push(d Walkable) {
	s = append(s, d)
}

func (s Stack) pop() Walkable {
	if len(s) == 0 {
		return nil
	}
	tmp := s[len(s)-1]
	s = s[:len(s)-1]
	return tmp
}

// New initalizes a new area.
func New(width, height int) *Area {
	a := Area{
		Terrain: make([][]Walkable, width),
		Objects: make(map[coord.Coord]Stack),
		Width:   width,
		Height:  height,
	}

	for x := 0; x < width; x++ {
		a.Terrain[x] = make([]Walkable, height)
	}
	return &a
}

/// Make DrawTerrain, DrawObjects, Draw
// Draw draws the terrain of the area to screen.AreaScreen.
func (a *Area) Draw() {
	for x := 0; x < screen.AreaScreen.Width; x++ {
		for y := 0; y < screen.AreaScreen.Height; y++ {
			// c := coord.Coord(x+screen.AreaScreen.XOffset, y+screen.AreaScreen.YOffset)
			// _, ok := a.Objects[c]
			// if ok {
			// 	o := a.Objects[c].pop()
			// 	if o == nil {
			// 		// draw terrain
			// 	}
			// }
			// terr := a.Terrain[x+screen.Screen.YOffset][y+screen.Screen.YOffset]
			// log.Fatal(len(a.Terrain), len(a.Terrain[x]))
			terr := a.Terrain[x][y]
			termbox.SetCell(x, y, terr.Graphic().Ch, terr.Graphic().Fg, terr.Graphic().Bg)
		}
	}
	termbox.Flush()
}

// MoveUp moves the object 1 tile upwards if possible.
func (a *Area) MoveUp(o *object.Object) error {
	return a.SetObjectXY(o, o.X, o.Y-1)
}

///
func (a *Area) MoveDown(o *object.Object) error {
	return a.SetObjectXY(o, o.X, o.Y+1)
}

///
func (a *Area) MoveRight(o *object.Object) error {
	return a.SetObjectXY(o, o.X+1, o.Y)
}

///
func (a *Area) MoveLeft(o *object.Object) error {
	return a.SetObjectXY(o, o.X-1, o.Y)
}

/// should return err on not exists?
///
func (a *Area) SetObjectXY(o *object.Object, x, y int) error {
	c := coord.Coord{x, y}
	p := coord.Plane{a.Width, a.Height}
	if !p.Exists(c) {
		// return errutil.NewNoPosf("(%d, %d) doesn't exist on that area.", x, y)
		return nil
	}

	// remove the object from the old position, add to the new position and
	// update both positions.
	if !a.Terrain[x][y].IsWalkable() {
		return MovementError{
			X: x,
			Y: y,
		}
	}
	// Update old position.
	a.Objects[coord.Coord{o.X, o.Y}].pop()
	draw.DrawXY(o.X, o.Y, a.Terrain[o.X][o.Y], screen.AreaScreen)

	// Update new position.
	o.X, o.Y = x, y
	a.Objects[c].push(o)
	draw.DrawXY(o.X, o.Y, o, screen.AreaScreen)

	return nil
}

// The error type if the unit can't move to the new coordinate
type MovementError struct {
	X int
	Y int
}

// Error message.
func (me MovementError) Error() string {
	return "couldn't move creature, path is blocked"
}
