package search

// Problem: Given a list of n elements, find the first repeated element.

// FindFirstDuplicates func
func FindFirstDuplicates(data []int) int {
	size := len(data)
	if size == 0 {
		return -1
	}
	h := make(map[int]int)
	for i := 0; i < size; i++ {
		h[data[i]]++
		if h[data[i]] > 1 {
			return i
		}
	}
	return -1
}
