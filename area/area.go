// Package area implements functions to draw and move moveable objects around
// in an area.
package area

import (
	"github.com/karlek/worc/coord"
	"github.com/karlek/worc/draw"
	"github.com/karlek/worc/screen"
)

// Area is a collection of terrain and objects placed on top of it.
type Area struct {
	Terrain       [][]Stackable
	Items         map[coord.Coord]*Stack
	Objects       map[coord.Coord]Stackable
	Monsters      map[coord.Coord]Moveable
	Width, Height int
	Screen        screen.Screen
}

// New initalizes a new area.
func New(width, height int, scr screen.Screen) *Area {
	a := Area{
		Terrain:  make([][]Stackable, width),
		Items:    make(map[coord.Coord]*Stack),
		Objects:  make(map[coord.Coord]Stackable),
		Monsters: make(map[coord.Coord]Moveable),
		Width:    width,
		Height:   height,
		Screen:   scr,
	}

	for x := 0; x < width; x++ {
		a.Terrain[x] = make([]Stackable, height)
	}
	return &a
}

// DrawTerrain draws the terrain of the area to screen.
func (a *Area) DrawTerrain() {
	for x := 0; x < a.Screen.Width; x++ {
		for y := 0; y < a.Screen.Height; y++ {
			draw.DrawXY(x, y, a.Terrain[x][y], a.Screen)
		}
	}
}

// DrawObjects draws all objects in an area to the screen.
func (a Area) DrawObjects() {
	for c, s := range a.Objects {
		if s == nil {
			continue
		}
		draw.DrawXY(c.X, c.Y, s, a.Screen)
	}
}

// DrawItems draws all items in an area to the screen.
func (a Area) DrawItems() {
	for c, s := range a.Items {
		w := s.Peek()
		if w == nil {
			continue
		}
		draw.DrawXY(c.X, c.Y, w, a.Screen)
	}
}

// DrawMonsters draws all monsters in an area to the screen.
func (a Area) DrawMonsters() {
	for c, m := range a.Monsters {
		if m == nil {
			continue
		}
		draw.DrawXY(c.X, c.Y, m, a.Screen)
	}
}

func (a Area) ReDraw(x, y int) {
	c := coord.Coord{x, y}
	p := coord.Plane{a.Width, a.Height}
	if !p.Contains(c) {
		return
	}

	if m := a.Monsters[c]; m != nil {
		draw.DrawXY(x, y, m, a.Screen)
		return
	}
	if w := a.Objects[c]; w != nil {
		draw.DrawXY(x, y, w, a.Screen)
		return
	}
	if i := a.Items[c].Peek(); i != nil {
		draw.DrawXY(x, y, i, a.Screen)
		return
	}
	draw.DrawXY(x, y, a.Terrain[x][y], a.Screen)
}

// Draw is a convenience function which draws both terrain and objects to the
// screen, in that order.
func (a Area) Draw() {
	a.DrawTerrain()
	a.DrawObjects()
	a.DrawItems()
	a.DrawMonsters()
}

// MoveUp moves a moveable object 1 tile upwards, if possible. Otherwise
// it returns the colliding object.
func (a *Area) MoveUp(m Moveable) *Collision {
	return a.SetObjectXY(m, m.X(), m.Y()-1)
}

// MoveDown moves a moveable object 1 tile downwards, if possible. Otherwise
// it returns the colliding object.
func (a *Area) MoveDown(m Moveable) *Collision {
	return a.SetObjectXY(m, m.X(), m.Y()+1)
}

// MoveRight moves a moveable object 1 tile rightwards, if possible. Otherwise
// it returns the colliding object.
func (a *Area) MoveRight(m Moveable) *Collision {
	return a.SetObjectXY(m, m.X()+1, m.Y())
}

// MoveLeft moves a moveable object 1 tile leftwards, if possible. Otherwise
// it returns the colliding object.
func (a *Area) MoveLeft(m Moveable) *Collision {
	return a.SetObjectXY(m, m.X()-1, m.Y())
}

type Collision struct {
	S Stackable
	X int
	Y int
}

func (a Area) IsXYStackable(x, y int) bool {
	c := coord.Coord{x, y}
	p := coord.Plane{a.Width, a.Height}
	if !p.Contains(c) {
		return false
	}

	// remove the object from the old position, add to the new position and
	// update both positions.
	if !a.Terrain[x][y].IsStackable() {
		return false
	}

	// test if an non-stackable object is already on that location.
	if mob := a.Monsters[c]; mob != nil {
		if !mob.IsStackable() {
			return false
		}
	}
	return true
}

// SetObjectXY sets an objects x and y value.
func (a *Area) SetObjectXY(m Moveable, x, y int) *Collision {
	c := coord.Coord{x, y}
	p := coord.Plane{a.Width, a.Height}
	if !p.Contains(c) {
		return nil
	}

	// remove the object from the old position, add to the new position and
	// update both positions.
	if !a.Terrain[x][y].IsStackable() {
		return &Collision{a.Terrain[x][y], x, y}
	}

	// test if an non-stackable object is already on that location.
	if mob := a.Monsters[c]; mob != nil {
		if !mob.IsStackable() {
			return &Collision{mob, x, y}
		}
	}

	// Update old position.
	c = coord.Coord{m.X(), m.Y()}
	if a.Monsters[c] != nil {
		a.Monsters[c] = nil
	}

	// Object beneath the current object.
	a.ReDraw(m.X(), m.Y())

	// Update new position.
	m.NewX(x)
	m.NewY(y)

	c = coord.Coord{m.X(), m.Y()}
	a.Monsters[c] = m

	// Redraw need coordinate.
	a.ReDraw(m.X(), m.Y())

	return nil
}

// Moveable asserts that the object can be moved.
type Moveable interface {
	X() int
	Y() int
	NewX(int)
	NewY(int)
	Stackable
}
