package types

type CellArray struct {
	dimension  int
	cells      []*Cell
	ordinalMap map[int]*Cell
	valueMap   map[int]*Cell
}

func NewCellArray(rank, dimension int) *CellArray {
	return &CellArray{
		dimension:  dimension,
		cells:      make([]*Cell, rank),
		ordinalMap: make(map[int]*Cell),
		valueMap:   make(map[int]*Cell),
	}
}

func (c *CellArray) HasValue(i int) bool {
	return c.valueMap[i] != nil
}

func (c *CellArray) SetCell(ordinal int, cell *Cell) {
	c.cells = append(c.cells, cell)
	c.valueMap[cell.value] = cell
	c.ordinalMap[ordinal] = cell
}

func (c *CellArray) GetCellByValue(value int) *Cell {
	return c.valueMap[value]
}

func (c *CellArray) GetCellByOrdinal(ordinal int) *Cell {
	return c.ordinalMap[ordinal]
}
