package main

import "fmt"

func main() {
	nums := []int{2, 15, 11, 7}
	target := 9
	sum := twoSum(nums, target)
	fmt.Println(sum)
}
func twoSum(nums []int, target int) []int {
	var result []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			tempSum := nums[i] + nums[j]
			if tempSum == target {
				result = append(result, i, j)
			}
		}
	}
	return result
}
