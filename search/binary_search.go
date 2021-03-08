package search

// BinarySearch function
func BinarySearch(data []int, key int) bool {
	size := len(data)
	low := 0
	high := size - 1
	mid := 0
	for low <= high {
		mid = (low + high) / 2
		if data[mid] == key {
			return true
		} else if data[mid] < key {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}

// BinarySearchV2 func
func BinarySearchV2(data []int, low int, high int, key int) bool {
	mid := (low + high) / 2
	if data[mid] == key {
		return true
	} else if data[mid] < key {
		return BinarySearchV2(data, mid+1, high, key)
	} else {
		return BinarySearchV2(data, low, mid-1, key)
	}
}
