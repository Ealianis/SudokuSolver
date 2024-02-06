package types

type Column struct {
	cells      []*Cell
	ordinalMap map[int]*Cell
	valueMap   map[int]*Cell
}

func NewColumn(rank int) *Column {
	return &Column{
		// The cell array size is based upon the Sudoku puzzle rank.
		cells:      make([]*Cell, rank),
		ordinalMap: make(map[int]*Cell),
		valueMap:   make(map[int]*Cell),
	}
}

func (c *Column) HasValue(i int) bool {
	return c.valueMap[i] != nil
}

func (c *Column) SetCell(ordinal int, cell *Cell) {
	c.cells = append(r.cells, cell)
	c.valueMap[cell.value] = cell
	c.ordinalMap[ordinal] = cell
}

func (c *Column) GetCellByValue(value int) *Cell {
	return c.valueMap[value]
}

func (c *Column) GetCellByOrdinal(ordinal int) *Cell {
	return c.ordinalMap[ordinal]
}
