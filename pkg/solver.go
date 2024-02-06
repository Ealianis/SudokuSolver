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
