package possition

import (
	"sort"
)

const (
	defaultPosition = -1
)

// Int is a wrapper to slice of integers, implements sort.Interface
type Int []int

func (r Int) Len() int           { return len(r) }
func (r Int) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r Int) Less(i, j int) bool { return r[i] < r[j] }

// Position returns the position of x in array or defaultPosition if x is not in array.
// It requires a sorted array as input, otherwise return defaultPosition
func Position(array []int, x int) int {
	var sorted Int
	sorted = array

	if !sort.IsSorted(sorted) {
		return defaultPosition
	}

	index := sort.SearchInts(array, x)
	// sort.SearchInts returns the len(array) if x is not in array, but defaultPosition is required
	if index == len(array) {
		return defaultPosition
	}

	return index
}

