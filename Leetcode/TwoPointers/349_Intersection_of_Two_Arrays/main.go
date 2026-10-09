package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}

	fmt.Println(intersection(nums1, nums2)) // [2]
}

// intersection returns unique elements present in both slices.
// Average time: O(n + m), where n = len(nums1), m = len(nums2).
// Auxiliary space: O(u), where u is the number of unique values in nums1.
// Result space: O(r), where r is the number of unique common values.
func intersection(nums1 []int, nums2 []int) []int {
	result := []int{}
	seen := make(map[int]bool)

	for _, num := range nums1 {
		seen[num] = true
	}

	for _, num := range nums2 {
		if seen[num] {
			result = append(result, num)
			seen[num] = false
		}
	}

	return result
}
