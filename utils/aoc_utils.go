package aoc_utils

import "slices"

func AbsInt(x int) int {
	if x < 0 {
		x = -x
	}

	return x
}

func IsOrdered(arr []int) bool {
	clone := slices.Clone(arr)
	if slices.IsSorted(clone) == false {
		slices.Reverse(clone)
	}

	return slices.IsSorted(clone)
}
