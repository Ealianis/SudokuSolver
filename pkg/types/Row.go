package types

type Row struct {
	cells      []*Cell
	ordinalMap map[int]*Cell
	valueMap   map[int]*Cell
}

func NewRow(rank int) *Row {
	return &Row{
		// The cell array size is based upon the Sudoku puzzle rank.
		cells:      make([]*Cell, rank),
		ordinalMap: make(map[int]*Cell),
		valueMap:   make(map[int]*Cell),
	}
}

func (r *Row) HasValue(i int) bool {
	return r.valueMap[i] != nil
}

func (r *Row) SetCell(ordinal int, cell *Cell) {
	r.cells = append(r.cells, cell)
	r.valueMap[cell.value] = cell
	r.ordinalMap[ordinal] = cell
}

func (r *Row) GetCellByValue(value int) *Cell {
	return r.valueMap[value]
}

func (r *Row) GetCellByOrdinal(ordinal int) *Cell {
	return r.ordinalMap[ordinal]
}
