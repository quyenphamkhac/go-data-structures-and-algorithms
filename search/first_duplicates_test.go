package search

import "testing"

func TestFindFirstDuplicates(t *testing.T) {
	data := []int{1, 2, 3, 2, 4}
	got := FindFirstDuplicates(data)
	if got != 3 {
		t.Errorf("Want 3; but got %d", got)
	}
}
