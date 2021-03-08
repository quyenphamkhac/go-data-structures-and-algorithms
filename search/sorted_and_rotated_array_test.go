package search

import "testing"

func TestSortedAndRotatedArraySearching(t *testing.T) {
	data := []int{4, 5, 6, 7, 0, 1, 2}
	got := RotatedSortedArray(data, 0)
	if got != 4 {
		t.Errorf("Want 4; but got %d", got)
	}
}
