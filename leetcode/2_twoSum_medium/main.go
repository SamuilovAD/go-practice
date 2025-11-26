package main

//Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
//You may assume that each input would have exactly one solution, and you may not use the same element twice.
//You can return the answer in any order.

func main() {
	twoSum([]int{2, 7, 11, 15}, 9)
}

func twoSum(nums []int, target int) []int {
	cache := make(map[int]int)
	for i, num := range nums {
		diffNum := target - num
		if j, ok := cache[diffNum]; ok {
			return []int{i, j}
		}
		cache[num] = i
	}
	return nil
}
