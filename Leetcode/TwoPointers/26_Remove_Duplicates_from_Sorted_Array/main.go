package main

import "fmt"

func main() {
	nums := []int{1, 1, 2, 2, 2, 2, 3, 3, 3, 4, 4, 4}
	resultLen := removeDuplicates(nums)
	fmt.Println(nums[0:resultLen])
}

// Time: O(n) Space: O(1)
func removeDuplicates(nums []int) int {
	insertIndex := 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] != nums[i] {
			nums[insertIndex] = nums[i]
			insertIndex++
		}
	}

	return insertIndex
}
