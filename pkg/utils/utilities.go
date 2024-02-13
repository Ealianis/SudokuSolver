package utils

import (
	"fmt"

	"github.com/ealianis/SudokuSolver/pkg/types"
)

func PivotAxis(axis types.Axis) types.Axis {
	switch axis {
	case types.YAxis:
		return types.XAxis
	case types.XAxis:
		return types.YAxis
	}
	panic(fmt.Sprintf("unknown axis %s, pivoted", axis))
}
