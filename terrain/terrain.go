package terrain

import "math/rand"
import "time"

import "github.com/karlek/worc/creature"
import "github.com/karlek/worc/menu"
import "github.com/nsf/termbox-go"

const scrollNum = 3

// Stationary objects on an area
type Terrain struct {
	Graphic  termbox.Cell
	Name     string
	Moveable bool
}

type Coord struct {
	X int
	Y int
}

// two-dimenstional Terrain slice which makes an area
type Area struct {
	Terrain [][]Terrain
	Objects map[Coord]interface{}
	Width   int
	Height  int
}

/// Ugly code
func GenArea(terrainObjs []Terrain, areaWidth, areaHeight int) Area {
	a := Area{
		Terrain: make([][]Terrain, areaHeight),
		Objects: make(map[Coord]interface{}),
		Width:   areaWidth,
		Height:  areaHeight,
	}

	for y := 0; y < areaHeight; y++ {
		a.Terrain[y] = make([]Terrain, areaWidth)
		for x := 0; x < areaWidth; x++ {
			a.Terrain[y][x] = terrainObjs[randInt(0, len(terrainObjs))]
		}
	}
	return a
}

/// Debug function to create mobs
func (a *Area) SpawnMobs(mob creature.Creature) {
	for i := 0; i < 10; i++ {
		tmpCoord := Coord{X: randInt(0, a.Width), Y: randInt(0, a.Height)}
		if _, found := a.Objects[tmpCoord]; !found && a.Terrain[tmpCoord.Y][tmpCoord.X].Moveable {
			mob.X = tmpCoord.X
			mob.Y = tmpCoord.Y
			a.Objects[tmpCoord] = mob
		}
	}
}

// Draws an area to terminal
func (a Area) DrawArea(as menu.AreaScreen) {
	for y := 0; y < as.Height; y++ {
		for x := 0; x < as.Width; x++ {
			terr := a.Terrain[y+as.YOffset][x+as.XOffset]
			termbox.SetCell(x, y, terr.Graphic.Ch, terr.Graphic.Fg, terr.Graphic.Bg)
		}
	}
	termbox.Flush()
}

// Draws all objects to terminal
func (a Area) DrawObjects(as menu.AreaScreen) {
	for y := 0; y < as.Height; y++ {
		for x := 0; x < as.Width; x++ {
			coord := Coord{X: x + as.XOffset, Y: y + as.YOffset}
			object, found := a.Objects[coord]
			if found {
				switch obj := object.(type) {
				case (creature.Creature):
					termbox.SetCell(coord.X-as.XOffset, coord.Y-as.YOffset, obj.Graphic.Ch, obj.Graphic.Fg, obj.Graphic.Bg)
				}
			}
		}
	}
	termbox.Flush()
}

// Convenience function to draw to terminal
func (a Area) Draw(as menu.AreaScreen) {
	a.DrawArea(as)
	a.DrawObjects(as)
}

/// Make me depracted
func (a Area) UpdateCoord(x, y int, as menu.AreaScreen) {
	terr := a.Terrain[y][x]
	termbox.SetCell(x, y, terr.Graphic.Ch, terr.Graphic.Fg, terr.Graphic.Bg)
	termbox.Flush()
}

// The error type if the unit can't move to the new coordinate
type MovementError struct {
	X int
	Y int
}

// Error message
func (me MovementError) Error() string {
	return "couldn't move creature, path is blocked"
}

// 168, 53 -> 168, 52
// Move creature
func (a *Area) Move(c *creature.Creature, x, y int, as *menu.AreaScreen) (err error) {

	// Prevent from moving outside the area to the top
	if y < 0 {
		y = 0
	}

	// Prevent from moving outside the area to the bottom
	if y == a.Height {
		y = a.Height - 1
	}

	// Prevent from moving outside the area to the left
	if x < 0 {
		x = 0
	}

	// Prevent from moving outside the area to the right
	if x == a.Width {
		x = a.Width - 1
	}

	_, found := a.Objects[Coord{X: x, Y: y}]
	if !found && a.Terrain[y][x].Moveable {
		delete(a.Objects, Coord{X: c.X, Y: c.Y})

		c.X = x
		c.Y = y
		a.Objects[Coord{X: x, Y: y}] = *c
	} else if !a.Terrain[y][x].Moveable {
		return MovementError{
			X: x,
			Y: y,
		}
	}

	if c.Name == "Hero" {

		// Scroll right
		if c.X == (as.Width-scrollNum)+as.XOffset && x < (a.Width-scrollNum) {
			as.XOffset += 1
		}

		// Scroll left
		if c.X < (as.XOffset+scrollNum) && x > scrollNum {
			as.XOffset -= 1
		}

		// Scroll down
		if c.Y == (as.Height-scrollNum)+as.YOffset && y < (a.Height-scrollNum) {
			as.YOffset += 1
		}

		// Scroll up
		if c.Y < (as.YOffset+scrollNum) && y+as.YOffset > scrollNum {
			as.YOffset -= 1
		}
	}

	a.Draw(*as)
	return nil
}

/// Temp func
func randInt(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}
