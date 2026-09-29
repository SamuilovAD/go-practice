package main

import "fmt"

func main() {
	nums := []int{-4, -1, 0, 3, 10}
	sortedSquares(nums)
	fmt.Println(nums)
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
