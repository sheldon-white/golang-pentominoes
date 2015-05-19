// Package pentominoes is a package.
package pentominoes

//import "fmt"

type Iterator struct {
	board          *Board
	pieceset       *PieceSet
	x, y           int // current point where we're trying to place a piece
	orientationIdx int // which orientation in the pieceSet
	pieceIdx       int // what piece in the piece set we're placing
}

// NewIterator is a constructor
func NewIterator(board *Board, pieceset *PieceSet) *Iterator {
	iterator := new(Iterator)
	iterator.board = board
	iterator.pieceset = pieceset

	return iterator
}

// Clone is a constructor
func (i *Iterator) Clone() *Iterator {
	newI := new(Iterator)
	newI.x = i.x
	newI.y = i.y
	newI.pieceIdx = i.pieceIdx
	newI.orientationIdx = i.orientationIdx
	newI.board = i.board.Clone() // copy the board
	newI.pieceset = i.pieceset   // use the same pieceset

	return newI
}

// Advance moves the iteration point to the next spot on the Board.
func (i *Iterator) Advance() bool {
	i.x++
	if i.x == BoardWidth {
		i.x = 0
		i.y++
		if i.y == BoardHeight {
			i.y = 0
			i.orientationIdx++
			if i.orientationIdx >= i.pieceset.GetOrientationCount(i.pieceIdx) {
				return false // if we've exausted the last orientation, the iterator is done
			}
		}
	}

	return true // still more opportunities for traversing the board with this piece
}

func (i *Iterator) getX() int {
	return i.x
}

func (i *Iterator) getY() int {
	return i.y
}

func (i *Iterator) getOrientationIdx() int {
	return i.orientationIdx
}
