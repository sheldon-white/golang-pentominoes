// Package pentominoes is a package.
package pentominoes

import "fmt"

type Iterator struct {
	board          *Board
	pieceset       *PieceSet
	x, y           uint8  // current point where we're trying to place a piece
	orientationIdx int  // which orientation in the pieceSet
	pieceIdx       int  // what piece in the piece set we're placing
	done           bool // are we done?
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

// IterateThroughLevel recursively traverses part of the legal-position tree.
// It's the heart or the program.
func (i *Iterator) IterateThroughLevel() {
	for !i.Done() {

		var x uint8 = i.GetX()
		var y uint8 = i.GetY()
		piece := i.GetOrientation()
		if i.board.CanPlacePieceAtPoint(piece, x, y) {
			// clone the iterator, place the piece on the clone, recurse down to next level
			newIterator := i.Clone()
			newIterator.x = 0
			newIterator.y = 0
			newIterator.orientationIdx = 0
			newIterator.pieceIdx++
			newIterator.board.PlacePiece(piece, x, y)
			fmt.Printf("\nx:%d y:%d pIdx:%d oIdx:%d\n", int(x), int(y), i.pieceIdx, i.orientationIdx)
			fmt.Printf("\n" + i.board.Display() + "\n\n")

						// did we find a solution?
			if newIterator.pieceIdx == 12 {
				fmt.Printf(newIterator.board.Display())
			} else {
				newIterator.IterateThroughLevel()
			}
		}
		i.Advance()
	}
	// return up a node on the tree
}

// Advance moves the iteration point to the next spot on the Board.
func (i *Iterator) Advance() {
	//fmt.Printf("iterator: %+v\n", i)
	i.x++
	if i.x == BoardWidth {
		i.x = 0
		i.y++
		if i.y == BoardHeight {
			i.y = 0
			i.orientationIdx++
			if i.orientationIdx >= i.pieceset.GetOrientationCount(i.pieceIdx) {
				i.done = true // if we've exausted the last orientation, the iterator is done
			}
		}
	}
}


func (i *Iterator) Done() bool {
	return i.done
}

func (i *Iterator) GetX() uint8 {
	return i.x
}

func (i *Iterator) GetY() uint8 {
	return i.y
}

func (i *Iterator) GetOrientation() *Piece {
	return i.pieceset.GetOrientation(i.pieceIdx, i.orientationIdx)
}
