// Package area implements functions to draw and move moveable objects around
// in an area.
package area

import (
	"github.com/karlek/reason/name"
	"github.com/karlek/reason/terrain"

	"github.com/karlek/worc/coord"
	"github.com/karlek/worc/draw"
	"github.com/karlek/worc/model"

	"github.com/mewkiz/pkg/errutil"
)

// Area is a collection of terrain and objects placed on top of it.
type Area struct {
	// TODO(_): rename to Cell.
	Terrain       [][]*terrain.Terrain
	Items         map[coord.Coord]*Stack
	Objects       map[coord.Coord]DrawPather
	Monsters      map[coord.Coord]Mover
	Width, Height int
}

// New initalizes a new area.
func New(width, height int) *Area {
	a := Area{
		Terrain:  make([][]*terrain.Terrain, width),
		Items:    make(map[coord.Coord]*Stack),
		Objects:  make(map[coord.Coord]DrawPather),
		Monsters: make(map[coord.Coord]Mover),
		Width:    width,
		Height:   height,
	}

	for x := 0; x < width; x++ {
		a.Terrain[x] = make([]*terrain.Terrain, height)
	}
	return &a
}

// MoveUp moves a moveable object 1 tile upwards, if possible. Otherwise
// it returns the colliding object.
func (a *Area) MoveUp(m Mover) (*Collision, error) {
	return a.SetObjectXY(m, m.X(), m.Y()-1)
}

// MoveDown moves a moveable object 1 tile downwards, if possible. Otherwise
// it returns the colliding object.
func (a *Area) MoveDown(m Mover) (*Collision, error) {
	return a.SetObjectXY(m, m.X(), m.Y()+1)
}

// MoveRight moves a moveable object 1 tile rightwards, if possible. Otherwise
// it returns the colliding object.
func (a *Area) MoveRight(m Mover) (*Collision, error) {
	return a.SetObjectXY(m, m.X()+1, m.Y())
}

// MoveLeft moves a moveable object 1 tile leftwards, if possible. Otherwise
// it returns the colliding object.
func (a *Area) MoveLeft(m Mover) (*Collision, error) {
	return a.SetObjectXY(m, m.X()-1, m.Y())
}

func (a Area) ExistsXY(x, y int) bool {
	c := coord.Coord{x, y}
	p := coord.Plane{a.Width, a.Height, 0, 0}
	return p.Contains(c)
}

func (a Area) IsXYPathable(x, y int) bool {
	if !a.ExistsXY(x, y) {
		return false
	}

	// remove the object from the old position, add to the new position and
	// update both positions.
	if !a.Terrain[x][y].IsPathable() {
		return false
	}

	c := coord.Coord{x, y}
	// test if an non-Pather object is already on that location.
	if mob := a.Monsters[c]; mob != nil {
		if !mob.IsPathable() {
			return false
		}
	}
	return true
}

type Collision struct {
	S DrawIsPather
	X int
	Y int
}

// SetObjectXY sets an objects x and y value.
func (a *Area) SetObjectXY(m Mover, x, y int) (col *Collision, err error) {
	c := coord.Coord{x, y}
	p := coord.Plane{a.Width, a.Height, 0, 0}
	if !p.Contains(c) {
		return nil, errutil.Newf("out of bounds.")
	}

	// remove the object from the old position, add to the new position and
	// update both positions.
	if !a.Terrain[x][y].IsPathable() {
		return &Collision{a.Terrain[x][y], x, y}, nil
	}

	// test if an non-Pather object is already on that location.
	if mob := a.Monsters[c]; mob != nil {
		if !mob.IsPathable() {
			return &Collision{mob, x, y}, nil
		}
	}

	// Update old position.
	c = coord.Coord{m.X(), m.Y()}
	if a.Monsters[c] != nil {
		a.Monsters[c] = nil
	}

	// Update new position.
	m.SetX(x)
	m.SetY(y)

	c = coord.Coord{m.X(), m.Y()}
	a.Monsters[c] = m

	return nil, nil
}

// Mover asserts that the object can be moved.
type Mover interface {
	model.Modeler
	name.Namer
	Pather
	draw.Drawable
}
