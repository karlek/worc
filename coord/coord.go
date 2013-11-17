/* === [ Coord ] ============================================================ */

///
package coord

///
type Coord struct {
	X, Y int
}

///
type Plane struct {
	Width, Height int
}

///
func (p Plane) Exists(c Coord) bool {
	return (c.X < p.Width && c.X >= 0) && (c.Y < p.Height && c.Y >= 0)
}
