package main

import "fmt"

func main() {
	fmt.Println(selectionSort([]int{1, 5, 17, 19, 11, 2}))
}

func selectionSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		minIdx := nums[i]
		for j := i + 1; j < len(nums); i++ {
			if nums[j] < nums[minIdx] {
				minIdx = j
			}
		}
		nums[i], nums[minIdx] = nums[minIdx], nums[i]
	}

	return nums
}
