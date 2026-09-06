package main

import "fmt"

func main() {
	sortedArr := []int{1, 5, 8, 12, 27, 41, 28, 34}
	fmt.Println(binarySearch(sortedArr, 12))
}

func binarySearch(inputSlice []int, targetValue int) int {
	leftIndex := 0
	rightIndex := len(inputSlice) - 1
	for leftIndex <= rightIndex {
		middleIndex := leftIndex + (rightIndex-leftIndex)/2
		if inputSlice[middleIndex] == targetValue {
			return middleIndex
		}
		if inputSlice[middleIndex] < targetValue {
			leftIndex = middleIndex + 1
		} else {
			rightIndex = middleIndex - 1
		}
	}

	return -1
}
