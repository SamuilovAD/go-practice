package main

import (
	"fmt"
)

func main() {
	nums := []int{-4, -1, 0, 3, 10}
	sortedSquares(nums)
	fmt.Println(nums)
	nums = []int{-4, -1, 0, 3, 10}
	result := sortedSquaresTwoPointers(nums)
	fmt.Println(result)
}

// Time: O(n^2) Space: O(1)
func sortedSquares(nums []int) []int {
	for i, num := range nums {
		nums[i] = num * num
	}
	for i := 0; i < len(nums); i++ {
		swapped := false
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				swapped = true
			}
		}
		if swapped == false {
			break
		}
	}

	return nums
}

// Time: O(N) Space: O(N)
func sortedSquaresTwoPointers(nums []int) []int {
	result := make([]int, len(nums))
	leftIndex := 0
	rightIndex := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		maxVal := 0
		if abs(nums[leftIndex]) > abs(nums[rightIndex]) {
			maxVal = nums[leftIndex]
			leftIndex++
		} else {
			maxVal = nums[rightIndex]
			rightIndex--
		}
		result[i] = maxVal * maxVal
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
