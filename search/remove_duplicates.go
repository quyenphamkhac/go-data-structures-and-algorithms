package search

import "sort"

// Problems: Remove duplicates in an integer list

// RemoveDuplicates using sorting
func RemoveDuplicates(data []int) int {
	j := 0
	size := len(data)
	if size == 0 {
		return 0
	}

	sort.Ints(data)
	for i := 0; i < size; i++ {
		if data[i] != data[j] {
			j++
			data[j] = data[i]
		}
	}
	return j + 1
}
