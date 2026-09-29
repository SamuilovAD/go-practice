package main

import "fmt"

func main() {
	fmt.Println("First bad API version:", firstBadVersion(17))
}

// Time: O(log N) Space: O(1)
func firstBadVersion(n int) int {
	left, right := 1, n
	for left < right {
		pivot := left + (right-left)/2
		if isBadVersion(pivot) {
			right = pivot
		} else {
			left = pivot + 1
		}
	}
	return left
}

func isBadVersion(version int) bool {
	if version > 5 {
		return true
	}
	return false
}
