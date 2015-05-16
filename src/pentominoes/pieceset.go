// Package stringutil contains utility functions for working with strings.
package pentominoes

import "container/list"

type pieceset struct {
	piecePermutations [12]*list.List
}

func CreatePieceSet() *pieceset {
	ps := new(pieceset)
	// create all the permutations of all the pieces
	// each list entry is a list of piece orientations
	ps.createAllPermutations()

	return ps;
}

func (ps pieceset) createAllPermutations() {
	l0 := list.New()
	ps.piecePermutations[0] = l0

	// ***
	//   **
	l0.PushFront(NewPiece(0, 4, 2, 'S', []uint8{0x3, 0xe}, map[int]bool {6: true, 40: true}))

	//  ***
	// **
	l0.PushFront(NewPiece(0, 4, 2, 'S', []uint8{0xc, 0x7}, map[int]bool {0: true, 56: true}))

	// **
	//  ***
	l0.PushFront(NewPiece(0, 4, 2, 'S', []uint8{0x7, 0xc}, map[int]bool {6: true, 40: true}))

	//   **
	// ***
	l0.PushFront(NewPiece(0, 4, 2, 'S', []uint8{0xe, 0x3}, map[int]bool {0: true, 56: true}))

	// *
	// **
  //  *
  //  *
	l0.PushFront(NewPiece(0, 2, 4, 'S', []uint8{0x1, 0x1, 0x3, 0x2}, map[int]bool {8: true, 20: true}))

	//  *
	//  *
  // **
  // *
	l0.PushFront(NewPiece(0, 2, 4, 'S', []uint8{0x2, 0x3, 0x1, 0x1}, map[int]bool {0: true, 58: true}))

	//  *
	// **
  // *
  // *
	l0.PushFront(NewPiece(0, 2, 4, 'S', []uint8{0x2, 0x2, 0x3, 0x1}, map[int]bool {0: true, 58: true}))

	// *
	// *
  // **
  //  *
	l0.PushFront(NewPiece(0, 2, 4, 'S', []uint8{0x1, 0x3, 0x2, 0x2}, map[int]bool {8: true, 20: true}))

	////////////////////////////////////////////////////////////////////////////////////////////

	l1 := list.New()
	ps.piecePermutations[1] = l1

	// *****
	l1.PushFront(NewPiece(1, 5, 1, 'A', []uint8{0x1f}, map[int]bool { }))

	// *
	// *
	// *
	// *
	// *
	l1.PushFront(NewPiece(1, 1, 5, 'A', []uint8{0x1, 0x1, 0x1, 0x1, 0x1}, map[int]bool { }))

	////////////////////////////////////////////////////////////////////////////////////////////

	l2 := list.New()
	ps.piecePermutations[2] = l2

	// ****
	// *
	l2.PushFront(NewPiece(2, 4, 2, 'B', []uint8{0x10, 0xf}, map[int]bool {0: true,}))

	// ****
	//    *
	l2.PushFront(NewPiece(2, 4, 2, 'B', []uint8{0x1, 0xf}, map[int]bool {6: true,}))

	// *
	// ****
	l2.PushFront(NewPiece(2, 4, 2, 'B', []uint8{0xf, 0x10}, map[int]bool {40: true,}))

	//    *
	// ****
	l2.PushFront(NewPiece(2, 4, 2, 'B', []uint8{0xf, 0x1}, map[int]bool {56: true,}))

	// **
	// *
	// *
	// *
	l2.PushFront(NewPiece(2, 2, 4, 'B', []uint8{0x2, 0x2, 0x2, 0x3}, map[int]bool {0: true,}))

	// **
	//  *
	//  *
	//  *
	l2.PushFront(NewPiece(2, 2, 4, 'B', []uint8{0x1, 0x1, 0x1, 0x3}, map[int]bool {8: true,}))

	// *
	// *
	// *
	// **
	l2.PushFront(NewPiece(2, 2, 4, 'B', []uint8{0x3, 0x2, 0x2, 0x2}, map[int]bool {20: true,}))

	//  *
	//  *
	//  *
	// **
	l2.PushFront(NewPiece(2, 2, 4, 'B', []uint8{0x3, 0x1, 0x1, 0x1}, map[int]bool {58: true,}))

	////////////////////////////////////////////////////////////////////////////////////////////

	l3 := list.New()
	ps.piecePermutations[3] = l3

	// ****
	//  *
	l3.PushFront(NewPiece(3, 4, 2, 'C', []uint8{0x4, 0xf}, map[int]bool {0: true, 6: true}))

	// ****
	//   *
	l3.PushFront(NewPiece(3, 4, 2, 'C', []uint8{0x2, 0xf}, map[int]bool {0: true, 6: true}))

	//   *
	// ****
	l3.PushFront(NewPiece(3, 4, 2, 'C', []uint8{0xf, 0x2}, map[int]bool {40: true, 56: true}))

	//  *
	// ****
	l3.PushFront(NewPiece(3, 4, 2, 'C', []uint8{0xf, 0x4}, map[int]bool {40: true, 56: true}))

	// *
	// **
  // *
  // *
	l3.PushFront(NewPiece(3, 2, 4, 'C', []uint8{0x2, 0x2, 0x3, 0x2}, map[int]bool {0: true, 20: true}))

	// *
	// *
  // **
  // *
	l3.PushFront(NewPiece(3, 2, 4, 'C', []uint8{0x2, 0x3, 0x2, 0x2}, map[int]bool {0: true, 20: true}))

	//  *
	// **
  //  *
  //  *
	l3.PushFront(NewPiece(3, 3, 4, 'C', []uint8{0x1, 0x1, 0x3, 0x1}, map[int]bool {8: true, 58: true}))

	//  *
	//  *
  // **
  //  *
	l3.PushFront(NewPiece(3, 3, 4, 'C', []uint8{0x1, 0x3, 0x1, 0x1}, map[int]bool {8: true, 58: true}))

	////////////////////////////////////////////////////////////////////////////////////////////

	l4 := list.New()
	ps.piecePermutations[4] = l4

	// ***
	// *
	// *
	l4.PushFront(NewPiece(4, 3, 3, 'D', []uint8{0x4, 0x4, 0x7}, map[int]bool {0: true}))

	// ***
	//   *
	//   *
	l4.PushFront(NewPiece(4, 3, 3, 'D', []uint8{0x1, 0x1, 0x7}, map[int]bool {7: true}))

	// *
	// *
	// ***
	l4.PushFront(NewPiece(4, 3, 3, 'D', []uint8{0x7, 0x4, 0x4}, map[int]bool {30: true}))

	//   *
	//   *
	// ***
	l4.PushFront(NewPiece(4, 3, 3, 'D', []uint8{0x7, 0x1, 0x1}, map[int]bool {57: true}))

	////////////////////////////////////////////////////////////////////////////////////////////

	l5 := list.New()
	ps.piecePermutations[5] = l5

	// ***
	//  *
	//  *
	l5.PushFront(NewPiece(5, 3, 3, 'E', []uint8{0x2, 0x2, 0x7}, map[int]bool {0: true, 7: true}))

	//  *
	//  *
	// ***
	l5.PushFront(NewPiece(5, 3, 3, 'E', []uint8{0x7, 0x2, 0x2}, map[int]bool {30: true, 57: true}))

	// *
	// ***
	// *
	l5.PushFront(NewPiece(5, 3, 3, 'E', []uint8{0x4, 0x7, 0x4}, map[int]bool {0: true, 30: true}))

	//   *
	// ***
	//   *
	l5.PushFront(NewPiece(5, 3, 3, 'E', []uint8{0x1, 0x7, 0x1}, map[int]bool {7: true, 57: true}))

	////////////////////////////////////////////////////////////////////////////////////////////

	l6 := list.New()
	ps.piecePermutations[6] = l6

	//  *
	// ***
	//  *
	l6.PushFront(NewPiece(6, 3, 3, 'F', []uint8{0x2, 0x7, 0x2}, map[int]bool {0: true, 30: true, 7: true, 57: true}))

	////////////////////////////////////////////////////////////////////////////////////////////

	l7 := list.New()
	ps.piecePermutations[7] = l7

	// ***
	// * *
	l7.PushFront(NewPiece(7, 3, 2, 'G', []uint8{0x5, 0x7}, map[int]bool {
		0: true, 1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true}))

	// * *
	// ***
	l7.PushFront(NewPiece(7, 3, 2, 'G', []uint8{0x7, 0x5}, map[int]bool {
		40: true, 41: true, 42: true, 43: true, 44: true, 45: true, 46: true, 47: true}))

// FIXME
	// **
	// *
  // **
	l7.PushFront(NewPiece(7, 2, 3, 'G', []uint8{0x3, 0x2, 0x3}, map[int]bool {
		0: true, 10: true, 20: true, 30: true}))

	// **
	//  *
  // **
	l7.PushFront(NewPiece(7, 2, 3, 'G', []uint8{0x3, 0x1, 0x3}, map[int]bool {
		8: true, 18: true, 28: true, 38: true}))

		////////////////////////////////////////////////////////////////////////////////////////////

	l8 := list.New()
	ps.piecePermutations[8] = l8

	// ***
	// **
	l8.PushFront(NewPiece(8, 3, 2, 'H', []uint8{0x6, 0x7}, map[int]bool {0: true}))

	// ***
	//  **
	l8.PushFront(NewPiece(8, 3, 2, 'H', []uint8{0x3, 0x7}, map[int]bool {7: true}))

	// **
	// ***
	l8.PushFront(NewPiece(8, 3, 2, 'H', []uint8{0x7, 0x6}, map[int]bool {40: true}))

	//  **
	// ***
	l8.PushFront(NewPiece(8, 3, 2, 'H', []uint8{0x7, 0x3}, map[int]bool {57: true}))

	// **
	// **
	// *
	l8.PushFront(NewPiece(8, 2, 3, 'H', []uint8{0x2, 0x3, 0x3}, map[int]bool {0: true}))

	// **
	// **
	//  *
	l8.PushFront(NewPiece(8, 2, 3, 'H', []uint8{0x1, 0x3, 0x3}, map[int]bool {8: true}))

	// *
	// **
	// **
	l8.PushFront(NewPiece(8, 2, 3, 'H', []uint8{0x3, 0x3, 0x2}, map[int]bool {30: true}))

	//  *
	// **
	// **
	l8.PushFront(NewPiece(8, 2, 3, 'H', []uint8{0x3, 0x3, 0x2}, map[int]bool {58: true}))

	////////////////////////////////////////////////////////////////////////////////////////////

	l9 := list.New()
	ps.piecePermutations[9] = l9

	// *
	// ***
	//   *
	l9.PushFront(NewPiece(9, 3, 3, 'O', []uint8{0x1, 0x7, 0x4}, map[int]bool {7: true, 30: true}))

	// **
	//  *
	//  **
	l9.PushFront(NewPiece(9, 3, 3, 'O', []uint8{0x3, 0x2, 0x6}, map[int]bool {7: true, 30: true}))

	//    *
	//  ***
	//  *
	l9.PushFront(NewPiece(9, 3, 3, 'O', []uint8{0x4, 0x7, 0x1}, map[int]bool {0: true, 57: true}))

	//   **
	//   *
	//  **
	l9.PushFront(NewPiece(9, 3, 3, 'O', []uint8{0x6, 0x2, 0x3}, map[int]bool {0: true, 57: true}))

	////////////////////////////////////////////////////////////////////////////////////////////

	l10 := list.New()
	ps.piecePermutations[10] = l10

	// *
	// **
	//  **
	l10.PushFront(NewPiece(10, 3, 3, 'O', []uint8{0x3, 0x6, 0x4}, map[int]bool {7: true, 30: true}))

	//  **
	// **
	// *
	l10.PushFront(NewPiece(10, 3, 3, 'O', []uint8{0x4, 0x6, 0x3}, map[int]bool {0: true, 57: true}))

	// **
	//  **
	//   *
	l10.PushFront(NewPiece(10, 3, 3, 'O', []uint8{0x1, 0x3, 0x6}, map[int]bool {7: true, 30: true}))

	//   *
	//  **
	// **
	l10.PushFront(NewPiece(10, 3, 3, 'O', []uint8{0x6, 0x3, 0x1}, map[int]bool {0: true, 57: true}))

	////////////////////////////////////////////////////////////////////////////////////////////

	l11 := list.New()
	ps.piecePermutations[11] = l11

	// *
	// ***
	//  *
	l11.PushFront(NewPiece(11, 3, 3, 'O', []uint8{0x2, 0x7, 0x4}, map[int]bool {0: true, 7: true, 30: true}))

	//   *
	// ***
	//  *
	l11.PushFront(NewPiece(11, 3, 3, 'O', []uint8{0x2, 0x7, 0x1}, map[int]bool {0: true, 7: true, 57: true}))

	//  *
	// ***
	// *
	l11.PushFront(NewPiece(11, 3, 3, 'O', []uint8{0x4, 0x7, 0x2}, map[int]bool {0: true, 30: true, 57: true}))

	//  *
	// ***
	//   *
	l11.PushFront(NewPiece(11, 3, 3, 'O', []uint8{0x1, 0x7, 0x2}, map[int]bool {7: true, 30: true, 57: true}))

	//  **
	// **
	//  *
	l11.PushFront(NewPiece(11, 3, 3, 'O', []uint8{0x2, 0x6, 0x3}, map[int]bool {0: true, 7: true, 57: true}))

	//  *
	// **
	//  **
	l11.PushFront(NewPiece(11, 3, 3, 'O', []uint8{0x3, 0x6, 0x2}, map[int]bool {7: true, 30: true, 57: true}))

	// **
	//  **
	//  *
	l11.PushFront(NewPiece(11, 3, 3, 'O', []uint8{0x2, 0x3, 0x6}, map[int]bool {0: true, 7: true, 30: true}))

	//  *
	//  **
	// **
	l11.PushFront(NewPiece(11, 3, 3, 'O', []uint8{0x6, 0x3, 0x2}, map[int]bool {0: true, 30: true, 57: true}))
}
