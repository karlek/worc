package area

import (
	"github.com/karlek/worc/coord"
	"github.com/karlek/worc/draw"
	"github.com/karlek/worc/screen"

	"github.com/nsf/termbox-go"
)

func (a Area) Draw(x, y, cameraX, cameraY int, scr screen.Screen) {
	c := coord.Coord{x, y}
	if m := a.Monsters[c]; m != nil {
		draw.DrawXY(x-cameraX, y-cameraY, m, scr)
		return
	}
	if w := a.Objects[c]; w != nil {
		draw.DrawXY(x-cameraX, y-cameraY, w, scr)
		return
	}
	if i := a.Items[c].Peek(); i != nil {
		draw.DrawXY(x-cameraX, y-cameraY, i, scr)
		return
	}
	draw.DrawXY(x-cameraX, y-cameraY, a.Terrain[x][y], scr)
}

func (a Area) DrawExplored(scr screen.Screen, cameraX, cameraY int) {
	for x := 0; x < a.Width; x++ {
		for y := 0; y < a.Height; y++ {
			if !a.Terrain[x][y].IsExplored() {
				continue
			}
			a.DrawMemory(x-cameraX, y-cameraY, cameraX, cameraY, scr)
		}
	}
}

func (a Area) DrawMemory(x, y int, cameraX, cameraY int, scr screen.Screen) {
	c := coord.Coord{x + scr.XOffset, y + scr.YOffset}
	p := coord.Plane{scr.Width + scr.XOffset, scr.Height + scr.YOffset, scr.XOffset, scr.YOffset}
	if !p.Contains(c) {
		return
	}
	termbox.SetCell(x+scr.XOffset, y+scr.YOffset, a.Terrain[x+cameraX][y+cameraY].Graphic().Ch, termbox.ColorBlack+termbox.AttrBold, termbox.ColorDefault)
}
