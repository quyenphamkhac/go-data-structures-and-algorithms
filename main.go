package main

import (
	"fmt"
	"go-algorithms/search"
)

func main() {
	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("v1")
	search.TripletSum(data, 8)
	fmt.Println("v2")
	search.TripletSumV2(data, 8)
}
