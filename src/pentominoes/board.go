package pentominoes

//import "fmt"

const BoardWidth = 10
const BoardHeight = 6

//
// Board is a 10x6 array of bits.
// lower-right = (0, 0)
// upper-left = (9, 5)
//
type Board struct {
	occupiedBits [6]uint16
	placedPieces map[int]*placement
}

// NewBoard is a constructor
func NewBoard() *Board {
	b := new(Board)
	b.placedPieces = make(map[int]*placement)

	return b
}

// Clone is a constructor
func (b *Board) Clone() *Board {
	newB := new(Board)
	newB.occupiedBits = b.occupiedBits
	newB.placedPieces = make(map[int]*placement)

	for k, v := range b.placedPieces {
		newB.placedPieces[k] = v
	}

	return newB
}

// PlacePiece places a piece at a point.
func (b *Board) PlacePiece(p *Piece, x uint8, y uint8) {
	b.placedPieces[p.index] = NewPlacement(p, x, y)
	b.paintPieceOnBoard(p, x, y)
}

func (b *Board) paintPieceOnBoard(p *Piece, x uint8, y uint8) {
	var row uint8
	shift := uint8(x)

	for row = 0; row < p.height; row++ {
		boardY := uint8(y) + row
		pieceBits := uint16(p.bits[row]) << shift
		b.occupiedBits[boardY] = b.occupiedBits[boardY] | pieceBits
	}
}

// CanPlacePieceAtPoint checks that a piece can be placed at a point.
func (b *Board) CanPlacePieceAtPoint(p *Piece, x uint8, y uint8) bool {
	var row uint8
	shift := uint8(x)
	cellIdx := int(y)*BoardWidth + int(x)
	if p.PositionForbidden(cellIdx) {
		return false
	}

	if x+p.width > BoardWidth {
		return false
	}
	if y+p.height > BoardHeight {
		return false
	}

	for row = 0; row < p.height; row++ {
		boardY := uint8(y) + row
		pieceBits := uint16(p.bits[row]) << shift
		if b.occupiedBits[boardY]&pieceBits != 0 {
			return false
		}
	}
	return true
}


// Display returns a string representation of the board.
func (b *Board) Display() string {
	output := ""
	offset := uint8(BoardWidth*BoardHeight - 1)
	var cells [BoardWidth * BoardHeight]byte
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
				if piece.bits[r]&(1<<uint8(c)) != 0 {
					cells[offset-((BoardWidth*(pieceY+r))+pieceX+c)] = piece.symbol
				}
			}
		}
	}
	for y := 0; y < BoardHeight; y++ {
		offset := BoardWidth * y
		cellRow := string(cells[offset : offset+BoardWidth])
		output += cellRow
		output += "\n"
	}

	return output
}

////////////////////////////////////////////////////////////////////////////////////////////////////////

func (b *Board) Fillable() bool {
	var checkedBits [6]uint16
	for y := 0; y < BoardHeight; y++ {
		var mask uint16 = 1
		for x := 0; x < BoardWidth; x++ {
			// Is the bit empty and unvisited?
			if (b.occupiedBits[y] & mask == 0) && (checkedBits[y] & mask == 0) {
				// check the holesize then...
				size := b.holeSize(x, y, &checkedBits)
				//fmt.Printf("hole size = %d\n", size)
				if size%5 != 0 {
					return false
				}
			}
			mask <<= 1
		}
	}
	return true
}

func (b *Board) holeSize(x int, y int, checkedBits *[6]uint16) int {
	count := 1
	var mask uint16

	// mark this bit as checked
	checkedBits[y] |= (1 << uint(x))

	// if we're not on the left edge, check left
	if x < BoardWidth - 1 {
		mask = 1 << uint(x+1)
		if (b.occupiedBits[y] & mask == 0) && (checkedBits[y] & mask == 0) {
			count += b.holeSize(x+1, y, checkedBits)
		}
	}
	// if we're not on the right edge, check right
	if x > 0 {
		mask = 1 << uint(x-1)
		if (b.occupiedBits[y] & mask == 0) && (checkedBits[y] & mask == 0) {
			count += b.holeSize(x-1, y, checkedBits)
		}
	}
	// if we're not on the top edge, check up
	if y < BoardHeight - 1 {
		mask = 1 << uint(x)
		if (b.occupiedBits[y+1] & mask == 0) && (checkedBits[y+1] & mask == 0) {
			count += b.holeSize(x, y+1, checkedBits)
		}
	}
	// if we're not on the bottom edge, check down
	if y > 0 {
		mask = 1 << uint(x)
		if (b.occupiedBits[y-1] & mask == 0) && (checkedBits[y-1] & mask == 0) {
			count += b.holeSize(x, y-1, checkedBits)
		}
	}

	return count
}
