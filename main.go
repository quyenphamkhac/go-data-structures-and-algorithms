package main

import (
	"fmt"
	"go-algorithms/search"
)

func main() {
	data1 := []int{1, 2, 3, 4}
	data2 := []int{4, 3, 2, 1}
	fmt.Println("Check permutation: ", search.CheckPermutation(data1, data2))
}
