package pkg

// ----- Structure -----
// A Sudoku puzzle's structure is determined by its rank n, where n is >= 2.
// A rank of n composes the board such that:
//	- There are n^2 rows & columns that compose the entire board.
//	- There are a total of n^2 * n^2 cells.
//	- There are n^2 sub-boards, or "blocks".
//		- Each block is composed of n rows & columns.
//		- Each block has n^2 cells.

// ----- Rules -----
// Cell values are bound by [1,n^2]
// For each row & column, all cell values must be unique.
// For each block, all cells contained therein must have unique values.

type orientation string

const (
	horizontalOrientation orientation = "horizontal"
	verticalOrientation   orientation = "vertical"
)

type Board struct {
	blocks []*Block
}

// A chain is a collection on n blocks.
// It orientation determines its "focus", columns or rows.
type Chain struct {
	blocks      []*Block
	orientation orientation
}

type Block struct {
	id   int
	rows []*Row
}

type Column struct {
	id    int
	cells []*Cell
}

type Row struct {
	id    int
	cells []*Cell
}

type Cell struct {
	value int
	hint  []int
}

func (r *Row) containsValue(i int) bool {
	for _, c := range r.cells {
		if i == c.value {
			return true
		}
	}
	return false
}
