package main

import "pentominoes"

func main() {
	pieceSet := pentominoes.CreatePieceSet()
	board := pentominoes.NewBoard()
	iterator := pentominoes.NewIterator(board, pieceSet)
  iterator.IterateThroughLevel()
}
