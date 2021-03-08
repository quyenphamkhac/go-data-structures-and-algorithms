package search

import (
	"sort"
)

// Problems: Given two integer Lists. You have to check whether they are permutation of each other.

// CheckPermutation use approach 1: sort and check
func CheckPermutation(data1 []int, data2 []int) bool {
	size1 := len(data1)
	size2 := len(data2)
	if size1 != size2 {
		return false
	}

	sort.Ints(data1)
	sort.Ints(data2)
	for i := 0; i < size1; i++ {
		if data1[i] != data2[i] {
			return false
		}
	}
	return true
}
