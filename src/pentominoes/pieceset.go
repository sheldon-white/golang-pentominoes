// Package stringutil contains utility functions for working with strings.
package pentominoes

import "container/list"

type pieceset struct {
	piecePermutations [12]*list.List
}

func CreatePieceSet() *pieceset {
	ps := new(pieceset)
	//ps.piecePermutations = new([12]list.List)
	// create all the permutations of all the pieces
	// each list entry is a list of piece orientations
	ps.createAllPermutations()

	return ps;
}

// FIXME flip to headecimal
func (ps pieceset) createAllPermutations() {
	l0 := list.New()
	ps.piecePermutations[0] = l0

	// ***
	//   **
	l0.PushFront(NewPiece(4, 2, 'S', []uint8{03, 016}, map[int]bool {8: true, 48: true}))

	//  ***
	// **
	l0.PushFront(NewPiece(4, 2, 'S', []uint8{014, 07}, map[int]bool {0: true, 56: true}))

	// **
	//  ***
	l0.PushFront(NewPiece(4, 2, 'S', []uint8{07, 014}, map[int]bool {8: true, 48: true}))

	//   **
	// ***
	l0.PushFront(NewPiece(4, 2, 'S', []uint8{016, 03}, map[int]bool {0: true, 56: true}))

	// *
	// **
  //  *
  //  *
	l0.PushFront(NewPiece(2, 4, 'S', []uint8{01, 01, 03, 02}, map[int]bool {10: true, 48: true}))

	//  *
	//  *
  // **
  // *
	l0.PushFront(NewPiece(2, 4, 'S', []uint8{02, 03, 01, 01}, map[int]bool {0: true, 58: true}))

	//  *
	// **
  // *
  // *
	l0.PushFront(NewPiece(2, 4, 'S', []uint8{02, 02, 03, 01}, map[int]bool {0: true, 58: true}))

	// *
	// *
  // **
  //  *
	l0.PushFront(NewPiece(2, 4, 'S', []uint8{01, 03, 02, 02}, map[int]bool {10: true, 48: true}))

	////////////////////////////////////////////////////////////////////////////////////////////

	l1 := list.New()
	ps.piecePermutations[1] = l1

	// *****
	l1.PushFront(NewPiece(5, 1, 'A', []uint8{037}, map[int]bool { }))

	// *
	// *
	// *
	// *
	// *
	l1.PushFront(NewPiece(1, 5, 'A', []uint8{01, 01, 01, 01, 01}, map[int]bool { }))

	////////////////////////////////////////////////////////////////////////////////////////////

	l2 := list.New()
	ps.piecePermutations[2] = l2

	// ****
	// *
	l2.PushFront(NewPiece(4, 2, 'B', []uint8{010, 017}, map[int]bool {0: true,}))

	// ****
	//    *
	l2.PushFront(NewPiece(4, 2, 'B', []uint8{01, 017}, map[int]bool {8: true,}))

	// *
	// ****
	l2.PushFront(NewPiece(4, 2, 'B', []uint8{017, 010}, map[int]bool {48: true,}))

	//    *
	// ****
	l2.PushFront(NewPiece(4, 2, 'B', []uint8{017, 01}, map[int]bool {56: true,}))

	// **
	// *
	// *
	// *
	l2.PushFront(NewPiece(2, 4, 'B', []uint8{02, 02, 02, 03}, map[int]bool {0: true,}))

	// **
	//  *
	//  *
	//  *
	l2.PushFront(NewPiece(2, 4, 'B', []uint8{01, 01, 01, 03}, map[int]bool {10: true,}))

	// *
	// *
	// *
	// **
	l2.PushFront(NewPiece(2, 4, 'B', []uint8{03, 02, 02, 02}, map[int]bool {48: true,}))

	//  *
	//  *
	//  *
	// **
	l2.PushFront(NewPiece(2, 4, 'B', []uint8{03, 01, 01, 01}, map[int]bool {58: true,}))

	////////////////////////////////////////////////////////////////////////////////////////////

	l3 := list.New()
	ps.piecePermutations[3] = l3

	// ****
	//  *
	l3.PushFront(NewPiece(4, 2, 'C', []uint8{04, 017}, map[int]bool {0: true, 8: true}))

	// ****
	//   *
	l3.PushFront(NewPiece(4, 2, 'C', []uint8{02, 017}, map[int]bool {0: true, 8: true}))

	//   *
	// ****
	l3.PushFront(NewPiece(4, 2, 'C', []uint8{017, 02}, map[int]bool {48: true, 56: true}))

	//  *
	// ****
	l3.PushFront(NewPiece(4, 2, 'C', []uint8{017, 04}, map[int]bool {48: true, 56: true}))

	// *
	// **
  // *
  // *
	l3.PushFront(NewPiece(2, 4, 'C', []uint8{02, 02, 03, 02}, map[int]bool {0: true, 48: true}))

	// *
	// *
  // **
  // *
	l3.PushFront(NewPiece(2, 4, 'C', []uint8{02, 03, 02, 02}, map[int]bool {0: true, 48: true}))

	//  *
	// **
  //  *
  //  *
	l3.PushFront(NewPiece(2, 4, 'C', []uint8{01, 01, 03, 01}, map[int]bool {10: true, 58: true}))

	//  *
	//  *
  // **
  //  *
	l3.PushFront(NewPiece(2, 4, 'C', []uint8{01, 03, 01, 01}, map[int]bool {10: true, 58: true}))

	////////////////////////////////////////////////////////////////////////////////////////////

	l4 := list.New()
	ps.piecePermutations[4] = l4

	// ***
	// *
	// *
	l4.PushFront(NewPiece(3, 3, 'D', []uint8{04, 04, 07}, map[int]bool {0: true}))

	// ***
	//   *
	//   *
	l4.PushFront(NewPiece(3, 3, 'D', []uint8{01, 01, 07}, map[int]bool {9: true}))

	// *
	// *
	// ***
	l4.PushFront(NewPiece(3, 3, 'D', []uint8{07, 04, 04}, map[int]bool {48: true}))

	//   *
	//   *
	// ***
	l4.PushFront(NewPiece(3, 3, 'D', []uint8{07, 01, 01}, map[int]bool {57: true}))

	////////////////////////////////////////////////////////////////////////////////////////////

	l5 := list.New()
	ps.piecePermutations[5] = l5

	// ***
	//  *
	//  *
	l5.PushFront(NewPiece(3, 3, 'E', []uint8{02, 02, 07}, map[int]bool {0: true, 9: true}))

	//  *
	//  *
	// ***
	l5.PushFront(NewPiece(3, 3, 'E', []uint8{07, 02, 02}, map[int]bool {48: true, 57: true}))

	// *
	// ***
	// *
	l5.PushFront(NewPiece(3, 3, 'E', []uint8{04, 07, 04}, map[int]bool {0: true, 48: true}))

	//   *
	// ***
	//   *
	l5.PushFront(NewPiece(3, 3, 'E', []uint8{01, 07, 01}, map[int]bool {9: true, 57: true}))

	////////////////////////////////////////////////////////////////////////////////////////////

	l6 := list.New()
	ps.piecePermutations[6] = l6

	//  *
	// ***
	//  *
	l6.PushFront(NewPiece(3, 3, 'F', []uint8{02, 07, 02}, map[int]bool {0: true, 48: true, 9: true, 57: true}))

	////////////////////////////////////////////////////////////////////////////////////////////

	l7 := list.New()
	ps.piecePermutations[7] = l7

	// ***
	// * *
	l7.PushFront(NewPiece(3, 2, 'G', []uint8{05, 07}, map[int]bool {
		0: true, 1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true}))

	// * *
	// ***
	l7.PushFront(NewPiece(3, 2, 'G', []uint8{07, 05}, map[int]bool {
		48: true, 49: true, 50: true, 51: true, 52: true, 53: true, 54: true, 55: true, 56: true, 57: true}))

	// **
	// *
  // **
	l7.PushFront(NewPiece(2, 3, 'G', []uint8{03, 02, 03}, map[int]bool {
		0: true, 12: true, 24: true, 36: true, 48: true}))

	// **
	//  *
  // **
	l7.PushFront(NewPiece(2, 3, 'G', []uint8{03, 01, 03}, map[int]bool {
		10: true, 22: true, 34: true, 46: true, 58: true}))

		////////////////////////////////////////////////////////////////////////////////////////////

	l8 := list.New()
	ps.piecePermutations[8] = l8

	// ***
	// **
	l8.PushFront(NewPiece(3, 2, 'H', []uint8{06, 07}, map[int]bool {0: true}))

	// ***
	//  **
	l8.PushFront(NewPiece(3, 2, 'H', []uint8{03, 07}, map[int]bool {9: true}))

	// **
	// ***
	l8.PushFront(NewPiece(3, 2, 'H', []uint8{07, 06}, map[int]bool {48: true}))

	//  **
	// ***
	l8.PushFront(NewPiece(3, 2, 'H', []uint8{07, 03}, map[int]bool {57: true}))

	// **
	// **
	// *
	l8.PushFront(NewPiece(2, 3, 'H', []uint8{02, 03, 03}, map[int]bool {0: true}))

	// **
	// **
	//  *
	l8.PushFront(NewPiece(2, 3, 'H', []uint8{01, 03, 03}, map[int]bool {10: true}))

	// *
	// **
	// **
	l8.PushFront(NewPiece(2, 3, 'H', []uint8{03, 03, 02}, map[int]bool {48: true}))

	//  *
	// **
	// **
	l8.PushFront(NewPiece(2, 3, 'H', []uint8{03, 03, 02}, map[int]bool {58: true}))

	////////////////////////////////////////////////////////////////////////////////////////////

	l9 := list.New()
	ps.piecePermutations[9] = l9

	// *
	// ***
	//   *
	l9.PushFront(NewPiece(3, 3, 'O', []uint8{01, 07, 04}, map[int]bool {11: true, 48: true}))

	// **
	//  *
	//  **
	l9.PushFront(NewPiece(3, 3, 'O', []uint8{03, 02, 05}, map[int]bool {11: true, 48: true}))

	//    *
	//  ***
	//  *
	l9.PushFront(NewPiece(3, 3, 'O', []uint8{04, 07, 01}, map[int]bool {0: true, 57: true}))

	//   **
	//   *
	//  **
	l9.PushFront(NewPiece(3, 3, 'O', []uint8{05, 04, 03}, map[int]bool {0: true, 57: true}))

	////////////////////////////////////////////////////////////////////////////////////////////

	l10 := list.New()
	ps.piecePermutations[10] = l10

	// *
	// **
	//  **
	l10.PushFront(NewPiece(3, 3, 'O', []uint8{03, 06, 04}, map[int]bool {9: true, 48: true}))

	//  **
	// **
	// *
	l10.PushFront(NewPiece(3, 3, 'O', []uint8{04, 06, 03}, map[int]bool {0: true, 57: true}))

	// **
	//  **
	//   *
	l10.PushFront(NewPiece(3, 3, 'O', []uint8{01, 03, 06}, map[int]bool {9: true, 48: true}))

	//   *
	//  **
	// **
	l10.PushFront(NewPiece(3, 3, 'O', []uint8{06, 03, 01}, map[int]bool {0: true, 57: true}))

	////////////////////////////////////////////////////////////////////////////////////////////

	l11 := list.New()
	ps.piecePermutations[11] = l11

	// *
	// ***
	//  *
	l11.PushFront(NewPiece(3, 3, 'O', []uint8{02, 07, 04}, map[int]bool {0: true, 9: true, 48: true}))

	//   *
	// ***
	//  *
	l11.PushFront(NewPiece(3, 3, 'O', []uint8{02, 07, 01}, map[int]bool {0: true, 9: true, 57: true}))

	//  *
	// ***
	// *
	l11.PushFront(NewPiece(3, 3, 'O', []uint8{04, 07, 02}, map[int]bool {0: true, 48: true, 57: true}))

	//  *
	// ***
	//   *
	l11.PushFront(NewPiece(3, 3, 'O', []uint8{01, 07, 02}, map[int]bool {9: true, 48: true, 57: true}))

	//  **
	// **
	//  *
	l11.PushFront(NewPiece(3, 3, 'O', []uint8{02, 06, 03}, map[int]bool {0: true, 9: true, 57: true}))

	//  *
	// **
	//  **
	l11.PushFront(NewPiece(3, 3, 'O', []uint8{03, 06, 02}, map[int]bool {9: true, 48: true, 57: true}))

	// **
	//  **
	//  *
	l11.PushFront(NewPiece(3, 3, 'O', []uint8{02, 03, 06}, map[int]bool {0: true, 9: true, 48: true}))

	//  *
	//  **
	// **
	l11.PushFront(NewPiece(3, 3, 'O', []uint8{06, 03, 02}, map[int]bool {0: true, 48: true, 57: true}))
}
