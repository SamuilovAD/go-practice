package main

import "fmt"

func main() {
	fmt.Printf("Pivot index %d", pivotIndex([]int{1, 7, 3, 6, 5, 6}))
}

func pivotIndex(nums []int) int {
	totalSum := 0
	leftSum := 0

	for _, num := range nums {
		totalSum += num
	}

	for i, num := range nums {
		if leftSum == totalSum-leftSum-num {
			return i
		}
		leftSum += num
	}

	return -1
}
