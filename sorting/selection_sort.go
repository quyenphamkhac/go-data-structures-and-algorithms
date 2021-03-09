package sorting

// SelectionSort implement
func SelectionSort(arr []int) {
	size := len(arr)
	var i, j, max int
	for i = 0; i < size-1; i++ {
		max = 0
		for j = 1; j < size-i-1; j++ {
			if arr[j] > arr[max] {
				max = j
			}
		}
		arr[size-i-1], arr[max] = arr[max], arr[size-i-1]
	}
}

// SelectionSort2 implement
func SelectionSort2(arr []int, comp func(int, int) bool) {
	size := len(arr)
	var i, j, s int
	for i = 0; i < size-1; i++ {
		s = i
		for j = i; j < size; j++ {
			if comp(arr[s], arr[j]) {
				s = j
			}
		}
		arr[i], arr[s] = arr[s], arr[i]
	}
}
