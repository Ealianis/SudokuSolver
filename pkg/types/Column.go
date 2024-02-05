package types

type Column struct {
	cells    []*Cell
	valueMap map[int]*Cell
}

func NewColumn(rank int) *Column {
	return &Column{
		// The cell array size is based upon the Sudoku puzzle rank.
		cells:    make([]*Cell, rank),
		valueMap: make(map[int]*Cell),
	}
}

func (r *Column) HasValue(i int) bool {
	return r.valueMap[i] != nil
}

func (r *Column) SetCell(cell *Cell) {
	r.cells[cell.coordinate.xOrdinal] = cell
}

func (r *Column) GetCellByValue(value int) *Cell {
	if r.HasValue(value) {
		return r.valueMap[value]
	}
	return nil
}
