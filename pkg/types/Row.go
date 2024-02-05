package types

type Row struct {
	cells    []*Cell
	valueMap map[int]*Cell
}

func NewRow(rank int) *Row {
	return &Row{
		// The array size is based upon the Sudoku puzzle rank.
		cells:    make([]*Cell, rank),
		valueMap: make(map[int]*Cell),
	}
}

func (r *Row) HasValue(i int) bool {
	return r.valueMap[i] != nil
}

func (r *Row) SetCell(cell *Cell) {
	r.cells[cell.coordinate.xOrdinal] = cell
}

func (r *Row) GetCellByValue(value int) *Cell {
	if r.HasValue(value) {
		return r.valueMap[value]
	}
	return nil
}
