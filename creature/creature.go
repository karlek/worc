package creature

import "github.com/nsf/termbox-go"

// Creatures are objects that can move on an area
type Creature struct {
	Graphic termbox.Cell
	Name    string
	X       int
	Y       int
	Stats   map[string]int
}
