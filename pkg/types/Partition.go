package types

type Partition struct {
	Chains     []*Chain
	CellArrays []*CellArray
	ValueMap   map[int][]*CellArray
}

func NewPartition(dimensions, rank int) *Partition {
	partition := &Partition{
		Chains:     make([]*Chain, dimensions),
		CellArrays: make([]*CellArray, rank*dimensions),
		ValueMap:   make(map[int][]*CellArray),
	}
	for i := 0; i <= dimensions-1; i++ {
		for ii := 0; ii <= rank-1; ii++ {
			ca := NewCellArray(i, rank)
			partition.CellArrays = append(partition.CellArrays, ca)
		}
	}

	return partition
}
