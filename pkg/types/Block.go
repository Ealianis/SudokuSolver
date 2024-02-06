package types

type CellCoordinates struct {
	column *Column
	row    *Row
}

type Block struct {
	columns  []*Column
	rows     []*Row
	valueMap map[int]*CellCoordinates
}

func NewBlock(rank int) *Block {
	block := &Block{
		// The column & row array size is based upon the Sudoku puzzle rank.
		columns:  make([]*Column, rank),
		rows:     make([]*Row, rank),
		valueMap: make(map[int]*CellCoordinates),
	}

	for i := 0; i <= rank-1; i++ {
		block.columns = append(block.columns, &Column{})
		block.rows = append(block.rows, &Row{})
	}

	return block
}

func (b *Block) HasValue(value int) bool {
	return b.valueMap[value] != nil
}

func (b *Block) SetCell(columnOrdinal, columnCellOrdinal, rowOrdinal, rowCellOrdinal, value int) *Cell {
	cell := &Cell{
		value: value,
	}

	col := b.columns[columnOrdinal-1]
	col.SetCell(columnCellOrdinal, cell)
	row := b.rows[rowOrdinal-1]
	row.SetCell(rowCellOrdinal, cell)

	cc := &CellCoordinates{
		column: col,
		row:    row,
	}
	b.valueMap[value] = cc

	return cell
}
