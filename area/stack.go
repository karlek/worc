package area

import (
	"github.com/karlek/worc/draw"
)

// Pathable are objects which can be drawn and answer the question if another
// object can be placed on top of it in the stack.
type Pathable interface {
	draw.Drawable
	IsPathable() bool
}

// Stack is a pile of stuff which can be walked upon or not walked upon.
type Stack []Pathable

// Peek returns the top most object of the stack without removing it.
func (s *Stack) Peek() Pathable {
	if s == nil {
		return nil
	}
	if len(*s) == 0 {
		return nil
	}
	return (*s)[len(*s)-1]
}

// Push adds a value ontop of the stack.
func (s *Stack) Push(d Pathable) {
	*s = append(*s, d)
}

// PopSecond returns the second value from the top.
func (s *Stack) PopSecond() Pathable {
	if len(*s) < 2 {
		return nil
	}

	tmp := (*s)[len(*s)-2]
	(*s)[len(*s)-2] = (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return tmp
}

// Pop returns the latest value.
func (s *Stack) Pop() Pathable {
	if len(*s) == 0 {
		return nil
	}
	tmp := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return tmp
}
