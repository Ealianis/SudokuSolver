package types

type Chain struct {
	blocks 		[]*Block
	dimension	int
	ordinal		int
	valueMap	[]*
}

func NewChain(rank, dimension, ordinal int) *Chain {
	return &Chain{
		dimension:   dimension,
		ordinal: ordinal,
		blocks: make([]*Block, rank),
	}
}

func (c *Chain) HasValue(value int) (bool, []*CellCoordinates) {
	var results []*CellCoordinates
	for _, b := range c.blocks {
		if b.valueMap[value] != nil {
			results = append(results, b.valueMap[value]...)
		}
	}
	return (len(results) > 0), results
}
