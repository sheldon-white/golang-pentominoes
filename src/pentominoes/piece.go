// Package stringutil contains utility functions for working with strings.
package pentominoes

import "fmt"

type piece struct {
	width, height int
	symbol rune
	bits []uint8
	forbiddenPositions map[int]bool
}

func NewPiece(width int, height int, symbol rune, bits []uint8, forbiddenPositions map[int]bool) *piece {
	p := new(piece)
	p.width = width
	p.height = height
	p.symbol = symbol
	p.bits = bits
	p.forbiddenPositions = forbiddenPositions

	return p
}

func (p piece) Display() string {
	return fmt.Sprintf("%b", p)
}
