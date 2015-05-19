// Package stringutil contains utility functions for working with strings.
package pentominoes

//import "fmt"

type Piece struct {
	index int
	width, height uint8
	symbol byte
	bits []uint8
	forbiddenPositions map[int]bool
}

func NewPiece(index int, width uint8, height uint8, symbol byte, bits []uint8, forbiddenPositions map[int]bool) *Piece {
	p := new(Piece)
	p.index = index
	p.width = width
	p.height = height
	p.symbol = symbol
	p.bits = bits
	p.forbiddenPositions = forbiddenPositions

	return p
}

func (p *Piece) PositionForbidden(idx int) bool {
	_, forbidden := p.forbiddenPositions[idx]

	return forbidden
}

func (p *Piece) Display() string {
	output := ""
	var r, row, col uint8
	row = p.height - 1
	for r = 0; r < p.height; r++ {
		rowBits := p.bits[row]
		row--
		//fmt.Printf("row: %d bits: %x\n", row, rowBits)

	  var mask uint8
		mask = (1 << (p.width - 1))
	  for col = 0; col < p.width; col++ {
			//fmt.Printf("col: %d mask: %d\n", col, mask)
			if (rowBits & mask != 0) {
				output += string(p.symbol)
			} else {
				output += "."
			}
			mask >>= 1
		}
		output += "\n"
	}
	return output
}
