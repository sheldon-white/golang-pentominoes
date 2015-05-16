package pentominoes

import "fmt"

func ExamplePieceConstructor() {
	p := NewPiece(1, 4, 2, 'S', []uint8{03, 016}, map[int]bool {8: true, 48: true})
	fmt.Printf("Piece: %+v\n", p)
	// Output:
  // Piece: &{index:1 width:4 height:2 symbol:83 bits:[3 14] forbiddenPositions:map[8:true 48:true]}
}


func ExamplePieceDisplay() {
	p := NewPiece(1, 4, 2, 'S', []uint8{03, 016}, map[int]bool {8: true, 48: true})
	output := p.Display()
	fmt.Printf(output)
	// Output:
  // SSS.
	// ..SS
}
