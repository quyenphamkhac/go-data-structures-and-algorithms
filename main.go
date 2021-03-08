package main

import (
	"fmt"
	"go-algorithms/search"
)

func main() {
	data3 := []int{1, 2, 2, 2, 3, 3}
	fmt.Println(data3)
	j := search.RemoveDuplicates(data3)
	fmt.Println(data3[:j], j)
}
