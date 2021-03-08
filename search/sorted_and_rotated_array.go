package search

// RotatedSortedArray func
func RotatedSortedArray(nums []int, target int) int {
	size := len(nums)
	low := 0
	high := size - 1
	if size == 0 {
		return -1
	}
	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > nums[low] {
			if nums[low] <= target && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			if nums[mid] < target && target < nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}
