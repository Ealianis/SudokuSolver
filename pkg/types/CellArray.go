package types

// A CellArray is an array of Cells which spans some dimension, e.g. x,y,z.
// Often referred to as rows, columns, and layers (for 3rd); the usage
// of the term CellArray provides an abstraction to such constructs
type CellArray struct {
	dimension  int
	cells      []*Cell
	ordinalMap map[int]*Cell
	valueMap   map[int]*Cell
}

func NewCellArray(dimension, rank int) *CellArray {
	return &CellArray{
		dimension:  dimension,
		cells:      make([]*Cell, rank),
		ordinalMap: make(map[int]*Cell),
		valueMap:   make(map[int]*Cell),
	}
}

func (ca *CellArray) HasValue(i int) bool {
	return ca.valueMap[i] != nil
}

func (ca *CellArray) SetCell(ordinal, value int) *Cell {
	c := &Cell{value: value}
	ca.cells = append(ca.cells, c)
	ca.valueMap[c.value] = c
	ca.ordinalMap[ordinal] = c
	return c
}

func (ca *CellArray) GetCellByValue(value int) *Cell {
	return ca.valueMap[value]
}

func (ca *CellArray) GetCellByOrdinal(ordinal int) *Cell {
	return ca.ordinalMap[ordinal]
}
