package area

import (
	"github.com/karlek/worc/draw"
)

// Stackable are objects which can be drawn and answer the question if another
// object can be placed on top of it in the stack.
type Stackable interface {
	draw.Drawable
	IsStackable() bool
}

// Stack is a pile of stuff which can be walked upon or not walked upon.
type Stack []Stackable

// Peek returns the top most object of the stack without removing it.
func (s *Stack) Peek() Stackable {
	if s == nil {
		return nil
	}
	if len(*s) == 0 {
		return nil
	}
	return (*s)[len(*s)-1]
}

// Push adds a value ontop of the stack.
func (s *Stack) Push(d Stackable) {
	*s = append(*s, d)
}

// PopSecond returns the second value from the top.
func (s *Stack) PopSecond() Stackable {
	if len(*s) < 2 {
		return nil
	}

	tmp := (*s)[len(*s)-2]
	(*s)[len(*s)-2] = (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return tmp
}

// Pop returns the latest value.
func (s *Stack) Pop() Stackable {
	if len(*s) == 0 {
		return nil
	}
	tmp := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return tmp
}
