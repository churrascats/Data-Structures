package tree

import "slices"

func OrderSliceAndRemoveDuplicates(treeItems []int) []int {
	slices.Sort(treeItems)
	orderedAndNoDuplicatesSlice := slices.Compact(treeItems)

	return orderedAndNoDuplicatesSlice
}
