package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	fmt.Println(sortSliceByParitySort([]int{3, 1, 2, 4}))
	fmt.Println(sortArrayByParityTwoPass([]int{7, 5, 6, 3, 1, 2, 4}))
}

// Time: O(N log N) Space: O(N)
func sortSliceByParitySort(nums []int) []int {
	slices.SortFunc(nums, func(a int, b int) int {
		return cmp.Compare(a%2, b%2)
	})
	return nums
}

// Time: O(N) Space: O(N)
func sortArrayByParityTwoPass(nums []int) []int {
	numsCopy := make([]int, len(nums))
	t := 0

	for _, elem := range nums {
		if elem%2 == 0 {
			numsCopy[t] = elem
			t++
		}
	}

	for _, elem := range nums {
		if elem%2 == 1 {
			numsCopy[t] = elem
			t++
		}
	}

	return numsCopy
}
