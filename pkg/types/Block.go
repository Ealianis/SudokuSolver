package types

import (
	"fmt"
)

type Block struct {
	// An array of chains which the block is apart of. The number of chains
	// is limited to dimension of the puzzle, e.g,. 2d, 3d, ...
	Chains     []*Chain
	// The collection of arrays that span the dimensions of the block, e.g., rows, columns, layers.
	CellArrays []*CellArray
	// A map for specific values within the block, an array of CellArrays is used as
	// a cell's ordinal within the array may differ between dimensions.
	ValueMap   map[int][]*CellArray
}

func NewBlock(dimensions, rank int) *Block {
	block := &Block{
		Chains: make([]*Chain, dimension),
		CellArrays: make([]*CellArray, rank*dimensions),
		ValueMap:   make(map[int][]*CellArray),
	}

	// The cell arrays will have an even number of x & y axis arrays.
	for i := 0; i <= (rank*dimensions)-1; i++ {
		axis := YAxis
		if i > rank-1 {
			axis = XAxis
		}

		block.CellArrays = append(block.CellArrays, &CellArray{axis: axis})
	}

	return block
}

func (b *Block) HasValue(value int) (bool, []*CellCoordinates) {
	return (len(b.ValueMap[value]) > 0), b.ValueMap[value]
}

func (b *Block) CreateCell(xOrdinal, yOrdinal, value int) (*Cell, error) {
	// todo - refactor for n-dimensional puzzle
	if b.ValueMap[value] != nil {
		c := b.ValueMap[value][0].
		return b.ValueMap[value] , fmt.Errorf("a cell already exists ")
	}

	cell := &Cell{
		value: value,
	}

	xCellArrays := b.getCellArraysByAxis(XAxis)
	yCellArrays := b.getCellArraysByAxis(YAxis)

	// todo - this is a bit confusing, rename or refactor?
	// xCellArray = array or rows
	// yOrdinal = which row
	// xordinal = which column
	xCellArrays[yOrdinal-1].SetCell(xOrdinal, cell)
	yCellArrays[xOrdinal-1].SetCell(yOrdinal, cell)

	cc := &CellCoordinates{}
	b.ValueMap[value] = append(b.ValueMap[value], cc)

	return cell
}

func (b *Block) getCellArraysByAxis(axis Axis) []*CellArray {
	var result []*CellArray
	for _, ca := range b.CellArrays {
		if ca.axis == axis {
			result = append(result, ca)
		}
	}
	return result
}
