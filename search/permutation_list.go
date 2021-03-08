package search

import (
	"sort"
)

// Problems: Given two integer Lists. You have to check whether they are permutation of each other.

// CheckPermutation solve problme use approach 1: sort and check
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

// CheckPermutationV2 solve problem use approach 2: hash-table
// 1. Create a Hash-Table for all the elements of the first list.
// 2. Traverse the other list from beginning to the end and search for each element in the Hash-Table.
// 3. If all the elements are found in, the Hash-Table return true, otherwise return false.
func CheckPermutationV2(data1 []int, data2 []int) bool {
	size1 := len(data1)
	size2 := len(data2)
	if size1 != size2 {
		return false
	}

	h := make(map[int]int)

	for i := 0; i < size1; i++ {
		h[data1[i]]++
		h[data2[i]]--
	}
	for i := 0; i < size1; i++ {
		if h[data1[i]] != 0 {
			return false
		}
	}
	return true
}
