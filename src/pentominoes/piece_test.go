package pentominoes

import (
	"fmt"
)

func ExamplePieceConstructor() {
	bits := []uint8{01, 01, 07}
	forbiddenPositions := map[int]bool {
		0: true,
		1: true,
	}
	p1 := NewPiece(3, 3, 'A', bits, forbiddenPositions)
	fmt.Printf("Piece: %+v\n", p1)
	// Output:
  // Piece: &{width:3 height:3 symbol:65 bits:[1 1 7] forbiddenPositions:map[0:true 1:true]}
}
