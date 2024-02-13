package types

type Board struct {
	Blocks []*Block
}

// A block is n-dimensional, and the arrays cardinality in all dimensions are equal.
// Given we use n-dimensional array of ints. We can determine the block we are in
// by tracking each dimensional array's ordinal.

func LoadBoard(dimensions, rank int) {
	var testPuzzle [9][9]int = [9][9]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	totalBlocks := rank * dimensions
	blocks := make([]*Block, totalBlocks)

	for i := 0; i <= totalBlocks; i++ {
		blocks = append(blocks, &Block{})
	}

	for oY, r := range testPuzzle {
		for oX, c := range r {

		}
	}
}
