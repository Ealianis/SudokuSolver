package types

type Block struct {
	columns  []*Column
	rows     []*Row
	valueMap map[int]*CellCoordinate
}

func NewBlock(rank int) *Block {
	return &Block{
		// The column & row array size is based upon the Sudoku puzzle rank.
		columns:  make([]*Column, rank),
		rows:     make([]*Row, rank),
		valueMap: make(map[int]*CellCoordinate),
	}
}
