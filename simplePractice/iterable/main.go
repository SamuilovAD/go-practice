package main

import (
	"fmt"
	"slices"
)

func main() {
	sum := sum([]int{1, 2, 3, 4, 5})
	fmt.Println("Sum: ", sum)
	max := slices.Max([]int{7, 2, 15, 4, 9})
	fmt.Println("Max: ", max)
	even := even([]int{1, 2, 155, 4, 2, 244, 11})
	fmt.Println("Even count: ", even)
	avg := average([]float64{1.1, 5.5, 2.2, 3.3, 1.2, 2.0})
	fmt.Println("Avarage: ", avg)
	found := slices.Index([]int{1, 2, 3, 4, 5}, 9)
	fmt.Println("Find element: ", found)
	reverse := []int{1, 3, 5, 7, 9, 2, 1}
	slices.Reverse(reverse)
	fmt.Println("Reverse: ", reverse)
	nums := []int{1, 5, 7, 2, 8, 1, 4}
	indexToBeDeleted := 3
	nums = slices.Delete(nums, indexToBeDeleted, indexToBeDeleted+1)
	fmt.Println("After delete: ", nums)
	nums = slices.Insert(nums, 4, 111)
	fmt.Println("After insert: ", nums)
	slice1 := []string{"a", "b", "c"}
	slice2 := []string{"d", "e", "f"}
	fmt.Println("Slices merge", append(slice1, slice2...))
	nums = []int{1, 5, 7, 2, 8, 1, 4}
	numsClone := slices.Clone(nums)
	numsCloneRef := &numsClone
	numsRef := &nums
	fmt.Println("Slices: ", nums, numsClone)
	fmt.Printf("Refs: %p, %p", numsRef, numsCloneRef)
}

func sum(items []int) int {
	result := 0
	for _, item := range items {
		result += item
	}
	return result
}

func even(items []int) int {
	result := 0
	for _, item := range items {
		if item%2 == 0 {
			result += 1
		}

	}
	return result
}

func average(items []float64) float64 {
	if len(items) == 0 {
		return 0
	}
	sum := 0.0
	for _, item := range items {
		sum += item
	}

	return sum / float64(len(items))
}
