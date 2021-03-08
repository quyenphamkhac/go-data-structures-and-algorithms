package search

import (
	"fmt"
	"sort"
)

// Problem: Given a list of n elements, write an algorithm to find three elements in a list whose sum is a given value.

// TripletSum - Approach 1: Brute force
func TripletSum(data []int, target int) {
	size := len(data)
	if size < 3 {
		fmt.Println("InvalidInput")
	}
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			for k := j + 1; k < size; k++ {
				sum := data[i] + data[j] + data[k]
				if sum == target {
					fmt.Printf("Triplet values: [%d, %d, %d]\n", data[i], data[j], data[k])
				}
			}
		}
	}
}

// TripletSumV2 - Approach 1: Brute force
func TripletSumV2(data []int, target int) {
	size := len(data)
	if size < 3 {
		fmt.Println("InvalidInput")
	}
	sort.Ints(data)
	for i := 0; i < size; i++ {
		left := i + 1
		right := size - 1
		for left < right {
			if data[i]+data[left]+data[right] == target {
				fmt.Printf("Triplet values: [%d, %d, %d]\n", data[i], data[left], data[right])
				left++
			} else if data[left]+data[right] < target-data[i] {
				left++
			} else {
				right--
			}
		}
	}
}
