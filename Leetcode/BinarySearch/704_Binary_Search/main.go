package main

import "fmt"

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	fmt.Println(search(nums, 9))
	fmt.Println(search(nums, 14))
}
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	var binarySearch func(
		nums []int,
		target int,
		left int,
		right int,
	) int
	binarySearch = func(
		nums []int,
		target int,
		left int,
		right int,
	) int {
		if left > right {
			return -1
		}
		pivot := left + (right-left)/2
		if target == nums[pivot] {
			return pivot
		}
		if target > nums[pivot] {
			return binarySearch(nums, target, pivot+1, right)
		}
		return binarySearch(nums, target, left, pivot-1)
	}

	return binarySearch(nums, target, 0, len(nums)-1)
}
