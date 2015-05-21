package pentominoes

import "fmt"

func ExampleOrientationCountDisplay() {
	ps := CreatePieceSet()
	for i := 0; i < 12; i++ {
		fmt.Printf("%d:%d ", i, ps.GetOrientationCount(i))
	}

	// Output:
	// 0:2 1:2 2:8 3:8 4:4 5:4 6:1 7:4 8:8 9:4 10:4 11:8
}
