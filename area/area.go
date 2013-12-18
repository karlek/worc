// Package area implements functions to draw and move moveable objects around
// in an area.
package area

import (
	"github.com/karlek/worc/coord"
)

// Area is a collection of terrain and objects placed on top of it.
type Area struct {
	Terrain       [][]Tile
	Items         map[coord.Coord]*Stack
	Objects       map[coord.Coord]Stackable
	Monsters      map[coord.Coord]Moveable
	Width, Height int
}

type Tile interface {
	Stackable
	IsBlockingLineOfSight() bool
	IsExplored() bool
	SetExplored(bool)
}

// New initalizes a new area.
func New(width, height int) *Area {
	a := Area{
		Terrain:  make([][]Tile, width),
		Items:    make(map[coord.Coord]*Stack),
		Objects:  make(map[coord.Coord]Stackable),
		Monsters: make(map[coord.Coord]Moveable),
		Width:    width,
		Height:   height,
	}

	for x := 0; x < width; x++ {
		a.Terrain[x] = make([]Tile, height)
	}
	return &a
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

func (a Area) ExistsXY(x, y int) bool {
	c := coord.Coord{x, y}
	p := coord.Plane{a.Width, a.Height, 0, 0}
	return p.Contains(c)
}

func (a Area) IsXYStackable(x, y int) bool {
	if !a.ExistsXY(x, y) {
		return false
	}

	// remove the object from the old position, add to the new position and
	// update both positions.
	if !a.Terrain[x][y].IsStackable() {
		return false
	}

	c := coord.Coord{x, y}
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
	p := coord.Plane{a.Width, a.Height, 0, 0}
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

	// Update new position.
	m.NewX(x)
	m.NewY(y)

	c = coord.Coord{m.X(), m.Y()}
	a.Monsters[c] = m

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
