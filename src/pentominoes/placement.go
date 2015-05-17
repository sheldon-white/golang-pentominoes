// Package stringutil contains utility functions for working with strings.
package pentominoes

type placement struct {
	p *piece
	x, y uint8
}

func NewPlacement(p *piece, x uint8, y uint8) *placement {
	pl := new(placement)
	pl.p = p
	pl.x = x
	pl.y = y

	return pl
}
