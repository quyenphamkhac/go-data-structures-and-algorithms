package main

import (
	"fmt"
	"go-algorithms/sorting"
)

func main() {
	data := []int{9, 1, 8, 2, 7, 3, 6, 4, 5}
	sorting.SelectionSort2(data, less)
	fmt.Println(data)
}

func less(v1 int, v2 int) bool {
	return v1 < v2
}

func more(v1 int, v2 int) bool {
	return v1 > v2
}
