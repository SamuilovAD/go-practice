package main

import "fmt"

func main() {
	area := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Printf("Max area for this set: %v, is %v", area, maxAreaTwoPointers(area))
}

//Time: O(n), Space: O(1)
func maxAreaTwoPointers(heightList []int) int {
	maxArea := 0
	leftIndex := 0
	rightIndex := len(heightList) - 1
	for leftIndex < rightIndex {
		width := rightIndex - leftIndex
		currentMaxArea := width * min(heightList[leftIndex], heightList[rightIndex])
		maxArea = max(currentMaxArea, maxArea)
		if heightList[leftIndex] < heightList[rightIndex] {
			leftIndex++
		} else {
			rightIndex--
		}
	}
	return maxArea
}
