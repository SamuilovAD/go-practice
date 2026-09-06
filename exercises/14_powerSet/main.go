package main

import "fmt"

func main() {
	fmt.Println(subset([]int{1, 2, 3}))
}

func subset(nums []int) [][]int {
	var result [][]int
	var path []int
	var backtrack func(start int)
	backtrack = func(start int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		result = append(result, tmp)
		for i := start; i < len(nums); i++ {
			path = append(path, nums[i]) // choose
			backtrack(i + 1)             // explore
			path = path[:len(path)-1]    // undo
		}
	}
	backtrack(0)
	return result
}
