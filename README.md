# SudokuSolver

## Structure

### Terms

> TODO @ealianis - Use prepositional logic notation to finish rules documentation.

Let `S` be a Sudoku Puzzle
Let `R` be the rank of `S`
Let `D` be the number of dimensions for `S`
Let `P` be a partition or "block" within `S`.
Let `A` be an array of cells across a specific dimension.

A Sudoku puzzle's structure is determined by its rank R, where R is >= 2.
A rank of 2 composes the board such that:

- There are R^D number of P.
- There are R^D
- There are n^2 sub-boards, or "blocks".
- Each block is composed of n rows & columns.
- Each block has n^2 cells.

## Rules

- The values of each cell are confined to boundary set of integers [1, n^2].
- Every cell within a board's row, column, and blocks should exclusively contain a single integer from the bounding set.

## Strategies

### Last remaining cell

For any given row, column, or block on the board, the value of last empty cell corresponds to the integer from the boundary set that is not present in any other cell within the same group.

### Orthogonal Shadow Casting

The value of any set cell prohibits all other cells in its orthogonal alignment, spanning across all rows and columns, from possessing the same value.

### Gridlock

The available values for a cell within a block is confined to the integers not yet used in any adjacent rows or columns extending across all blocks.
