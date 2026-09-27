package main

import "fmt"

func main() {
	slice := []int{5, 4, 1, 2, 6, 8, 3, 3}
	bubbleSort(slice)
	fmt.Printf("Bubble sort: %v", slice)
}

// Time: O(n²), Space: (1)
func bubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		swapped := false
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				swapped = true
				nums[j+1], nums[j] = nums[j], nums[j+1]
			}
		}
		if !swapped {
			break
		}
	}
}
