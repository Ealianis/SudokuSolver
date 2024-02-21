package types

type Chain struct {
	dimension  int
	partitions []*Partition
}

func NewChain(dimension, rank int) *Chain {
	return &Chain{
		dimension:  dimension,
		partitions: make([]*Partition, rank),
	}
}
