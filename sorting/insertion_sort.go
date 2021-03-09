package sorting

// InsertionSort implement
func InsertionSort(arr []int, comp func(int, int) bool) {
	size := len(arr)
	var marker, unsortedIndex, temp int
	for marker = 1; marker < size; marker++ {
		temp = arr[marker]
		for unsortedIndex = marker; unsortedIndex > 0 && comp(arr[unsortedIndex-1], temp); unsortedIndex-- {
			arr[unsortedIndex] = arr[unsortedIndex-1]
		}
		arr[unsortedIndex] = temp
	}
}
