///
package menu

import (
	"github.com/nsf/termbox-go"
)

var (
	statusMesg  []string
	statusWidth int

	statusX int
	statusY int
)

// Set number of status messages to be shown
func SetStatusSize(width, height int) {
	statusWidth = width
	statusMesg = make([]string, height)
}

// Set number of status messages to be shown
func SetStatusLoc(x, y int) {
	statusX = x
	statusY = y
}

// Prints to string to screen taking x coordinate, y coordinate,
// foreground color (attributes) and background color (attributes)
func print(str string, x, y, width int, fg termbox.Attribute, bg termbox.Attribute) {
	// Clears the line from old characters
	for i := len(str); i < width; i++ {
		termbox.SetCell(x+i, y, ' ', termbox.ColorDefault, termbox.ColorDefault)
	}

	for charOffset, char := range str {
		termbox.SetCell(x+charOffset, y, char, fg, bg)
	}
}

/// Breaks on very long strings
// Writes string to status buffer
func PrintStatus(str string) {
	statusLen := len(statusMesg)

	var strs []string
	if len(str) > statusWidth {
		// We reverse the order because []statusMesg is FILO
		strs = append(strs, str[statusWidth:])
		strs = append(strs, str[:statusWidth])

	} else {
		strs = append(strs, str)
	}

	for _, str := range strs {
		// Insert the new string first
		statusMesg = append(statusMesg[:0], append([]string{str}, statusMesg[0:]...)...)

		// Delete the last string
		statusMesg = append(statusMesg[:statusLen], statusMesg[statusLen+1:]...)
	}

	for y, status := range statusMesg {
		print(status, statusX, statusY+y, statusWidth, termbox.ColorWhite, termbox.ColorDefault)
	}
	termbox.Flush()
}
