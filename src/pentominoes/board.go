// Package stringutil contains utility functions for working with strings.
package pentominoes

import "container/list"

const width = 12
const height = 6

//
// a board is a 12x6 array of bits.
// lower-right = (0, 0)
// upper-left = (11, 5)
//
type board struct {
	occupied_bits [6]uint16
	placedPieces *list.List
	pieceSet *pieceset
}

func NewBoard() *board {
	b := new(board)
	//b.occupied_bits = new([6]uint16)
	b.placedPieces = list.New()
	b.pieceSet = CreatePieceSet()

	return b;
}
