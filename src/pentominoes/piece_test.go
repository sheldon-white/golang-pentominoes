package pentominoes

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
	 )

func ExamplePieceConstructor() {
	p := NewPiece(1, 4, 2, 'S', []uint8{03, 016}, map[int]bool{8: true, 48: true})
	fmt.Printf("Piece: %d %d %d %d\n", p.index, p.width, p.height, p.symbol)
	// Output:
	// Piece: 1 4 2 83
}

func ExamplePieceDisplay() {
	p := NewPiece(1, 4, 2, 'S', []uint8{03, 016}, map[int]bool{8: true, 48: true})
	output := p.Display()
	fmt.Printf(output)
	// Output:
	// SSS.
	// ..SS
}

func TestForbiddenPositions(t *testing.T) {
	p := NewPiece(1, 4, 2, 'S', []uint8{03, 016}, map[int]bool{8: true, 48: true})
	assert.True(t, p.PositionForbidden(8), "piece isn't placeable there")
}
