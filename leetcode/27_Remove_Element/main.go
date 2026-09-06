package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 3, 4, 2, 5, 7}
	newLen := removeElement(nums, 2)
	fmt.Printf("%v", nums[:newLen])
}

func removeElement(nums []int, val int) int {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}
