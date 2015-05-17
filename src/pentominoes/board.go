// Package stringutil contains utility functions for working with strings.
package pentominoes

//import "fmt"

const width = 10
const height = 6

//
// a board is a 10x6 array of bits.
// lower-right = (0, 0)
// upper-left = (9, 5)
//
type board struct {
	occupiedBits [6]uint16
	placedPieces map[int]*placement
	pieceSet *pieceset
}

func NewBoard() *board {
	b := new(board)
	b.pieceSet = CreatePieceSet()
	b.placedPieces = make(map[int]*placement)

	return b;
}

func (b *board) PlacePiece(p *piece, x uint8, y uint8) {
	b.placedPieces[p.index] = NewPlacement(p, x, y)
	b.paintPieceOnBoard(p, x, y)
}

func (b *board) paintPieceOnBoard(p *piece, x uint8, y uint8) {
	var row uint8
	shift := uint8(x)

	for row = 0; row < p.height; row++ {
		boardY := uint8(y) + row
		pieceBits := uint16(p.bits[row]) << shift
		b.occupiedBits[boardY] = b.occupiedBits[boardY] | pieceBits
	}
}

func (b *board) CanPlacePieceAtPoint(p *piece, x uint8, y uint8) bool {
	var row uint8
	shift := uint8(x)

	for row = 0; row < p.height; row++ {
		boardY := uint8(y) + row
		pieceBits := uint16(p.bits[row]) << shift
		if (b.occupiedBits[boardY] & pieceBits != 0) {
			return false
		}
	}

	return true
}

func (b *board) Display() string {
	output := ""
	offset := uint8(width * height - 1)
	var cells [width * height]byte
	for i := range cells {
	    cells[i] = '.'
	}

	for _, pl := range b.placedPieces {
		piece := pl.p
		pieceX := pl.x
		pieceY := pl.y
		var r, c uint8

		for r = 0; r < piece.height; r++ {
			for c = 0; c < piece.width; c++ {
				if (piece.bits[r] & (1 << uint8(c)) != 0) {
					cells[offset - ((width * (pieceY + r)) + pieceX + c)] = piece.symbol
				}
			}
		}
	}
	for y := 0; y < height; y++ {
		offset := width * y
		cellRow := string(cells[offset : offset + width])
		output += cellRow
		output += "\n"
	}

	return output
}
