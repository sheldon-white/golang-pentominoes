// Package stringutil contains utility functions for working with strings.
package pentominoes

type position struct {
	x, y int
}

func NewPosition(x int, y int) *position {
	p := new(position)
	p.x = x
	p.y = y
	
	return p
}
