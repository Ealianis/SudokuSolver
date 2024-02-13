# SudokuSolver

## Structure

A Sudoku puzzle's structure is determined by its rank n, where n is >= 2.
A rank of n composes the board such that:

- There are n^2 rows & columns that compose the entire board.
- There are a total of n^2 * n^2 cells.
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
