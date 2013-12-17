// Package status implements functions to print status messages.
package status

import (
	"github.com/nsf/termbox-go"
)

var (
	statusMesg []string
	width      int

	x int
	y int
)

// SetSize determines how many status messages will be shown.
func SetSize(w, h int) {
	width = w
	statusMesg = make([]string, h)
}

// SetLoc determines where the status messages will be outputted.
func SetLoc(xVal, yVal int) {
	x = xVal
	y = yVal
}

// Prints to string to screen taking x coordinate, y coordinate,
// foreground color (attributes) and background color (attributes)
func print(str string, x, y, width int, fg termbox.Attribute, bg termbox.Attribute) {
	// Clears the line from old characters.
	for i := len(str); i < width; i++ {
		termbox.SetCell(x+i, y, ' ', termbox.ColorDefault, termbox.ColorDefault)
	}

	for charOffset, char := range str {
		termbox.SetCell(x+charOffset, y, char, fg, bg)
	}
}

// Print writes a string to the status buffer.
func Print(str string) {
	/// Print Breaks on very long strings.
	statusLen := len(statusMesg)

	var strs []string
	if len(str) > width {
		// We reverse the order because []statusMesg is FILO
		strs = append(strs, str[width:])
		strs = append(strs, str[:width])

	} else {
		strs = append(strs, str)
	}

	for _, str := range strs {
		// Insert the new string first
		statusMesg = append(statusMesg[:0], append([]string{str}, statusMesg[0:]...)...)

		// Delete the last string
		statusMesg = append(statusMesg[:statusLen], statusMesg[statusLen+1:]...)
	}

	for offset, status := range statusMesg {
		print(status, x, y+len(statusMesg)-1-offset, width, termbox.ColorWhite, termbox.ColorDefault)
	}
	termbox.Flush()
}

func Update() {
	for offset, status := range statusMesg {
		print(status, x, y+len(statusMesg)-1-offset, width, termbox.ColorWhite, termbox.ColorDefault)
	}
	termbox.Flush()
}
