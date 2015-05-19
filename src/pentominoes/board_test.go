package pentominoes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleBoardConstructor() {
	b := NewBoard()
	fmt.Printf("Board: %+v\n", b)
}

func TestBoardClone(t *testing.T) {
	b := NewBoard()
	p1 := NewPiece(1, 4, 2, 'S', []uint8{0x3, 0xe}, map[int]bool{8: true, 48: true})
	p2 := NewPiece(2, 3, 3, 'Q', []uint8{0x1, 0x7, 0x4}, map[int]bool{8: true, 48: true})
	b.PlacePiece(p1, 0, 0)
	b.PlacePiece(p2, 3, 3)
	o1 := fmt.Sprintf("Board: %+v\n", b.occupiedBits)
	b2 := b.Clone()
	o2 := fmt.Sprintf("Board: %+v\n", b2.occupiedBits)
	assert.Equal(t, o1, o2, "boards should match")
}

func ExampleBoardDisplay() {
	b := NewBoard()
	p1 := NewPiece(1, 4, 2, 'S', []uint8{0x3, 0xe}, map[int]bool{8: true, 48: true})
	p2 := NewPiece(2, 3, 3, 'Q', []uint8{0x1, 0x7, 0x4}, map[int]bool{8: true, 48: true})
	b.PlacePiece(p1, 0, 0)
	b.PlacePiece(p2, 3, 3)
	fmt.Printf(b.Display())

	// Output:
	// ....Q.....
	// ....QQQ...
	// ......Q...
	// ..........
	// ......SSS.
	// ........SS
}

func TestCanPlacePieceAtPoint(t *testing.T) {
	b := NewBoard()
	p1 := NewPiece(1, 4, 2, 'S', []uint8{0x3, 0xe}, map[int]bool{8: true, 48: true})
	p2 := NewPiece(2, 3, 3, 'Q', []uint8{0x1, 0x7, 0x4}, map[int]bool{8: true, 48: true})

	assert.True(t, b.CanPlacePieceAtPoint(p1, 0, 0), "should be able to place the piece")
	b.PlacePiece(p1, 0, 0)
	assert.False(t, b.CanPlacePieceAtPoint(p2, 1, 1), "shouldn't be able to place the piece")
}
