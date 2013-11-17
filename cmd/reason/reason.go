package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/karlek/torc/area"
	"github.com/karlek/torc/coord"
	"github.com/karlek/torc/draw"
	"github.com/karlek/torc/menu"
	"github.com/karlek/torc/object"
	"github.com/karlek/torc/screen"
	"github.com/mewkiz/pkg/errutil"
	"github.com/nsf/termbox-go"
)

const (
	// Status messages
	pathIsBlockedStr = "Your path is blocked by %s"

	// Game screen size
	AreaScreenWidth  = 100
	AreaScreenHeight = 30

	// Status bar screen coordinates
	StatusX = 5
	StatusY = 31

	// Status bar size
	StatusWidth  = 100 - StatusX*2
	StatusHeight = 7
)

///
type Creature struct {
	O    object.Object
	Name string
}

// The Hero! This is the unit that the user will control.
var Hero = Creature{
	Name: "Hero",
	O: object.Object{
		G: termbox.Cell{
			Ch: '@',
			Fg: termbox.ColorWhite + termbox.AttrBold,
		},
		Walkable: false,
	},
}

///
type Environment struct {
	area.Walkable
	O        object.Object
	Name     string
	walkable bool
}

func (e Environment) IsWalkable() bool {
	return e.walkable
}

func (e Environment) Graphic() termbox.Cell {
	return e.O.Graphic()
}

var Wall = Environment{
	O: object.Object{G: termbox.Cell{
		Ch: '#',
		Fg: termbox.ColorWhite + termbox.AttrBold,
	}},
	Name:     "a wall",
	walkable: false,
}

var Soil = Environment{
	O: object.Object{G: termbox.Cell{
		Ch: '.',
		Fg: termbox.ColorYellow,
	}},
	Name:     "soil",
	walkable: true,
}

// Error wrapper
func main() {
	err := reason()
	if err != nil {
		log.Fatalln(err)
	}
}

func reason() (err error) {
	err = termbox.Init()
	if err != nil {
		return errutil.Err(err)
	}
	/// This will never run on 'q' (quit event)
	defer termbox.Close()

	// Init status menu
	menu.SetStatusSize(StatusWidth, StatusHeight)
	menu.SetStatusLoc(StatusX, StatusY)

	// Area is where all terrain and object data are stored.

	var ms = []area.Walkable{
		&Soil,
		&Soil,
		&Soil,
		&Soil,
		&Soil,
		&Soil,
		&Wall,
	}

	// AreaScreen is the active viewport of the area.
	// Used on big areas that are bigger then the screen.
	screen.AreaScreen.Width = 100
	screen.AreaScreen.Height = 30

	a := GenArea(ms, 100, 30)
	a.Draw()
	draw.DrawXY(0, 0, &Hero.O, screen.AreaScreen)

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Ch {
			case 'q':
				termbox.Close()
				os.Exit(0)
			}
			switch ev.Key {

			// Movement
			case termbox.KeyArrowUp:
				err = a.MoveUp(&Hero.O)
			case termbox.KeyArrowDown:
				err = a.MoveDown(&Hero.O)
			case termbox.KeyArrowLeft:
				err = a.MoveLeft(&Hero.O)
			case termbox.KeyArrowRight:
				err = a.MoveRight(&Hero.O)
			}

			switch err := err.(type) {
			case area.MovementError:
				menu.PrintStatus(fmt.Sprintf(pathIsBlockedStr, a.Terrain[err.X][err.Y].(*Environment).Name))
			}
		}
	}

	return nil
}

/// Debug function to generate terrain.
func GenArea(ms []area.Walkable, width, height int) area.Area {
	a := area.Area{
		Terrain: make([][]area.Walkable, width),
		Objects: make(map[coord.Coord]area.Stack),
		Width:   width,
		Height:  height,
	}

	for x := 0; x < width; x++ {
		a.Terrain[x] = make([]area.Walkable, height)
		for y := 0; y < height; y++ {
			a.Terrain[x][y] = ms[randInt(0, len(ms))]
		}
	}
	return a
}

/// Temp func
func randInt(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}
