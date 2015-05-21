// Package stringutil contains utility functions for working with strings.
package pentominoes

type PieceSet struct {
	piecePermutations [12][]*Piece
}

func CreatePieceSet() *PieceSet {
	ps := new(PieceSet)
	// create all the permutations of all the pieces
	// each list entry is a list of piece orientations
	ps.createAllPermutations()

	return ps
}

func (ps PieceSet) GetOrientationCount(pieceIdx int) int {
	return len(ps.piecePermutations[pieceIdx])
}

func (ps *PieceSet) GetOrientation(pieceIdx int, orientationIdx int) *Piece {
	if pieceIdx < 12 {
		l := ps.piecePermutations[pieceIdx]
		if orientationIdx < len(l) {
			return l[orientationIdx]
		}
	}

	return nil
}

func (ps *PieceSet) createAllPermutations() {
	l0 := []*Piece{}
	// By adding only 2 orientations to this piece, we automatically eliminate solutions that are
	// merely mirror images of other solutions

	// ***
	//   **
	l0 = append(l0, NewPiece(0, 4, 2, 'S', []uint8{0x3, 0xe}, map[int]bool{6: true, 40: true}))

	// *
	// **
	//  *
	//  *
	l0 = append(l0, NewPiece(0, 2, 4, 'S', []uint8{0x1, 0x1, 0x3, 0x2}, map[int]bool{8: true, 20: true}))

	ps.piecePermutations[0] = l0

	////////////////////////////////////////////////////////////////////////////////////////////

	l1 := []*Piece{}

	// *****
	l1 = append(l1, NewPiece(1, 5, 1, 'A', []uint8{0x1f}, map[int]bool{10: true, 15: true, 40: true, 45: true}))

	// *
	// *
	// *
	// *
	// *
	l1 = append(l1, NewPiece(1, 1, 5, 'A', []uint8{0x1, 0x1, 0x1, 0x1, 0x1}, map[int]bool{1: true, 8: true, 51: true, 58: true}))

	ps.piecePermutations[1] = l1

	////////////////////////////////////////////////////////////////////////////////////////////

	l2 := []*Piece{}

	// ****
	// *
	l2 = append(l2, NewPiece(2, 4, 2, 'B', []uint8{0x8, 0xf}, map[int]bool{0: true}))

	// ****
	//    *
	l2 = append(l2, NewPiece(2, 4, 2, 'B', []uint8{0x1, 0xf}, map[int]bool{6: true}))

	// *
	// ****
	l2 = append(l2, NewPiece(2, 4, 2, 'B', []uint8{0xf, 0x8}, map[int]bool{40: true}))

	//    *
	// ****
	l2 = append(l2, NewPiece(2, 4, 2, 'B', []uint8{0xf, 0x1}, map[int]bool{56: true}))

	// **
	// *
	// *
	// *
	l2 = append(l2, NewPiece(2, 2, 4, 'B', []uint8{0x2, 0x2, 0x2, 0x3}, map[int]bool{0: true}))

	// **
	//  *
	//  *
	//  *
	l2 = append(l2, NewPiece(2, 2, 4, 'B', []uint8{0x1, 0x1, 0x1, 0x3}, map[int]bool{8: true}))

	// *
	// *
	// *
	// **
	l2 = append(l2, NewPiece(2, 2, 4, 'B', []uint8{0x3, 0x2, 0x2, 0x2}, map[int]bool{20: true}))

	//  *
	//  *
	//  *
	// **
	l2 = append(l2, NewPiece(2, 2, 4, 'B', []uint8{0x3, 0x1, 0x1, 0x1}, map[int]bool{58: true}))

	ps.piecePermutations[2] = l2

	////////////////////////////////////////////////////////////////////////////////////////////

	l3 := []*Piece{}

	// ****
	//  *
	l3 = append(l3, NewPiece(3, 4, 2, 'C', []uint8{0x4, 0xf}, map[int]bool{0: true, 6: true}))

	// ****
	//   *
	l3 = append(l3, NewPiece(3, 4, 2, 'C', []uint8{0x2, 0xf}, map[int]bool{0: true, 6: true}))

	//   *
	// ****
	l3 = append(l3, NewPiece(3, 4, 2, 'C', []uint8{0xf, 0x2}, map[int]bool{40: true, 56: true}))

	//  *
	// ****
	l3 = append(l3, NewPiece(3, 4, 2, 'C', []uint8{0xf, 0x4}, map[int]bool{40: true, 56: true}))

	// *
	// **
	// *
	// *
	l3 = append(l3, NewPiece(3, 2, 4, 'C', []uint8{0x2, 0x2, 0x3, 0x2}, map[int]bool{0: true, 20: true}))

	// *
	// *
	// **
	// *
	l3 = append(l3, NewPiece(3, 2, 4, 'C', []uint8{0x2, 0x3, 0x2, 0x2}, map[int]bool{0: true, 20: true}))

	//  *
	// **
	//  *
	//  *
	l3 = append(l3, NewPiece(3, 3, 4, 'C', []uint8{0x1, 0x1, 0x3, 0x1}, map[int]bool{8: true, 58: true}))

	//  *
	//  *
	// **
	//  *
	l3 = append(l3, NewPiece(3, 3, 4, 'C', []uint8{0x1, 0x3, 0x1, 0x1}, map[int]bool{8: true, 58: true}))

	ps.piecePermutations[3] = l3

	////////////////////////////////////////////////////////////////////////////////////////////

	l4 := []*Piece{}

	// ***
	// *
	// *
	l4 = append(l4, NewPiece(4, 3, 3, 'D', []uint8{0x4, 0x4, 0x7}, map[int]bool{0: true}))

	// ***
	//   *
	//   *
	l4 = append(l4, NewPiece(4, 3, 3, 'D', []uint8{0x1, 0x1, 0x7}, map[int]bool{7: true}))

	// *
	// *
	// ***
	l4 = append(l4, NewPiece(4, 3, 3, 'D', []uint8{0x7, 0x4, 0x4}, map[int]bool{30: true}))

	//   *
	//   *
	// ***
	l4 = append(l4, NewPiece(4, 3, 3, 'D', []uint8{0x7, 0x1, 0x1}, map[int]bool{57: true}))

	ps.piecePermutations[4] = l4

	////////////////////////////////////////////////////////////////////////////////////////////

	l5 := []*Piece{}

	// ***
	//  *
	//  *
	l5 = append(l5, NewPiece(5, 3, 3, 'E', []uint8{0x2, 0x2, 0x7}, map[int]bool{0: true, 7: true}))

	//  *
	//  *
	// ***
	l5 = append(l5, NewPiece(5, 3, 3, 'E', []uint8{0x7, 0x2, 0x2}, map[int]bool{30: true, 57: true}))

	// *
	// ***
	// *
	l5 = append(l5, NewPiece(5, 3, 3, 'E', []uint8{0x4, 0x7, 0x4}, map[int]bool{0: true, 30: true}))

	//   *
	// ***
	//   *
	l5 = append(l5, NewPiece(5, 3, 3, 'E', []uint8{0x1, 0x7, 0x1}, map[int]bool{7: true, 57: true}))

	ps.piecePermutations[5] = l5

	////////////////////////////////////////////////////////////////////////////////////////////

	l6 := []*Piece{}

	//  *
	// ***
	//  *
	l6 = append(l6, NewPiece(6, 3, 3, 'F', []uint8{0x2, 0x7, 0x2}, map[int]bool{0: true, 30: true, 7: true, 57: true}))

	ps.piecePermutations[6] = l6

	////////////////////////////////////////////////////////////////////////////////////////////

	l7 := []*Piece{}

	// ***
	// * *
	l7 = append(l7, NewPiece(7, 3, 2, 'G', []uint8{0x5, 0x7}, map[int]bool{
		0: true, 1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true}))

	// * *
	// ***
	l7 = append(l7, NewPiece(7, 3, 2, 'G', []uint8{0x7, 0x5}, map[int]bool{
		40: true, 41: true, 42: true, 43: true, 44: true, 45: true, 46: true, 47: true}))

	// FIXME
	// **
	// *
	// **
	l7 = append(l7, NewPiece(7, 2, 3, 'G', []uint8{0x3, 0x2, 0x3}, map[int]bool{
		0: true, 10: true, 20: true, 30: true}))

	// **
	//  *
	// **
	l7 = append(l7, NewPiece(7, 2, 3, 'G', []uint8{0x3, 0x1, 0x3}, map[int]bool{
		8: true, 18: true, 28: true, 38: true}))

	ps.piecePermutations[7] = l7

	////////////////////////////////////////////////////////////////////////////////////////////

	l8 := []*Piece{}

	// ***
	// **
	l8 = append(l8, NewPiece(8, 3, 2, 'H', []uint8{0x6, 0x7}, map[int]bool{0: true}))

	// ***
	//  **
	l8 = append(l8, NewPiece(8, 3, 2, 'H', []uint8{0x3, 0x7}, map[int]bool{7: true}))

	// **
	// ***
	l8 = append(l8, NewPiece(8, 3, 2, 'H', []uint8{0x7, 0x6}, map[int]bool{40: true}))

	//  **
	// ***
	l8 = append(l8, NewPiece(8, 3, 2, 'H', []uint8{0x7, 0x3}, map[int]bool{57: true}))

	// **
	// **
	// *
	l8 = append(l8, NewPiece(8, 2, 3, 'H', []uint8{0x2, 0x3, 0x3}, map[int]bool{0: true}))

	// **
	// **
	//  *
	l8 = append(l8, NewPiece(8, 2, 3, 'H', []uint8{0x1, 0x3, 0x3}, map[int]bool{8: true}))

	// *
	// **
	// **
	l8 = append(l8, NewPiece(8, 2, 3, 'H', []uint8{0x3, 0x3, 0x2}, map[int]bool{30: true}))

	//  *
	// **
	// **
	l8 = append(l8, NewPiece(8, 2, 3, 'H', []uint8{0x3, 0x3, 0x2}, map[int]bool{58: true}))

	ps.piecePermutations[8] = l8

	////////////////////////////////////////////////////////////////////////////////////////////

	l9 := []*Piece{}

	// *
	// ***
	//   *
	l9 = append(l9, NewPiece(9, 3, 3, 'J', []uint8{0x1, 0x7, 0x4}, map[int]bool{7: true, 30: true}))

	// **
	//  *
	//  **
	l9 = append(l9, NewPiece(9, 3, 3, 'J', []uint8{0x3, 0x2, 0x6}, map[int]bool{7: true, 30: true}))

	//    *
	//  ***
	//  *
	l9 = append(l9, NewPiece(9, 3, 3, 'J', []uint8{0x4, 0x7, 0x1}, map[int]bool{0: true, 57: true}))

	//   **
	//   *
	//  **
	l9 = append(l9, NewPiece(9, 3, 3, 'J', []uint8{0x6, 0x2, 0x3}, map[int]bool{0: true, 57: true}))

	ps.piecePermutations[9] = l9

	////////////////////////////////////////////////////////////////////////////////////////////

	l10 := []*Piece{}

	// *
	// **
	//  **
	l10 = append(l10, NewPiece(10, 3, 3, 'K', []uint8{0x3, 0x6, 0x4}, map[int]bool{7: true, 30: true}))

	//  **
	// **
	// *
	l10 = append(l10, NewPiece(10, 3, 3, 'K', []uint8{0x4, 0x6, 0x3}, map[int]bool{0: true, 57: true}))

	// **
	//  **
	//   *
	l10 = append(l10, NewPiece(10, 3, 3, 'K', []uint8{0x1, 0x3, 0x6}, map[int]bool{7: true, 30: true}))

	//   *
	//  **
	// **
	l10 = append(l10, NewPiece(10, 3, 3, 'K', []uint8{0x6, 0x3, 0x1}, map[int]bool{0: true, 57: true}))

	ps.piecePermutations[10] = l10

	////////////////////////////////////////////////////////////////////////////////////////////

	l11 := []*Piece{}

	// *
	// ***
	//  *
	l11 = append(l11, NewPiece(11, 3, 3, 'O', []uint8{0x2, 0x7, 0x4}, map[int]bool{0: true, 7: true, 30: true}))

	//   *
	// ***
	//  *
	l11 = append(l11, NewPiece(11, 3, 3, 'O', []uint8{0x2, 0x7, 0x1}, map[int]bool{0: true, 7: true, 57: true}))

	//  *
	// ***
	// *
	l11 = append(l11, NewPiece(11, 3, 3, 'O', []uint8{0x4, 0x7, 0x2}, map[int]bool{0: true, 30: true, 57: true}))

	//  *
	// ***
	//   *
	l11 = append(l11, NewPiece(11, 3, 3, 'O', []uint8{0x1, 0x7, 0x2}, map[int]bool{7: true, 30: true, 57: true}))

	//  **
	// **
	//  *
	l11 = append(l11, NewPiece(11, 3, 3, 'O', []uint8{0x2, 0x6, 0x3}, map[int]bool{0: true, 7: true, 57: true}))

	//  *
	// **
	//  **
	l11 = append(l11, NewPiece(11, 3, 3, 'O', []uint8{0x3, 0x6, 0x2}, map[int]bool{7: true, 30: true, 57: true}))

	// **
	//  **
	//  *
	l11 = append(l11, NewPiece(11, 3, 3, 'O', []uint8{0x2, 0x3, 0x6}, map[int]bool{0: true, 7: true, 30: true}))

	//  *
	//  **
	// **
	l11 = append(l11, NewPiece(11, 3, 3, 'O', []uint8{0x6, 0x3, 0x2}, map[int]bool{0: true, 30: true, 57: true}))

	ps.piecePermutations[11] = l11
}
