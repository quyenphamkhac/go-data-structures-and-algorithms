package main

import (
	"fmt"
	"go-algorithms/sorting"
)

func main() {
	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 9, 8}
	sorting.BubbleSort2(data, more)
	fmt.Println(data)
}

func less(v1 int, v2 int) bool {
	return v1 < v2
}

func more(v1 int, v2 int) bool {
	return v1 > v2
}
