package search

import "testing"

func TestCheckPermutation(t *testing.T) {
	data1 := []int{1, 2, 3, 4}
	data2 := []int{3, 4, 2, 1}
	got := CheckPermutation(data1, data2)
	if got != true {
		t.Errorf("Want True; but got %t", got)
	}
}

func TestCheckPermutationV2(t *testing.T) {
	data1 := []int{1, 2, 3, 4}
	data2 := []int{3, 4, 2, 1}
	got := CheckPermutationV2(data1, data2)
	if got != true {
		t.Errorf("Want True; but got %t", got)
	}
}
