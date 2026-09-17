package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge3Pointers(nums1, 3, nums2, 3)
	fmt.Print(nums1)
}

func mergeSimple(nums1 []int, m int, nums2 []int, n int) {
	for i := 0; i < n; i++ {
		nums1[m+i] = nums2[i]
	}
	sort.Ints(nums1)

}

func merge3Pointers(nums1 []int, m int, nums2 []int, n int) {
	nums1Copy := append([]int{}, nums1[:m]...)
	p1, p2 := 0, 0
	for p := 0; p < m+n; p++ {
		if p2 >= n || (p1 < m && nums1Copy[p1] < nums2[p2]) {
			nums1[p] = nums1Copy[p1]
			p1++
		} else {
			nums1[p] = nums2[p2]
			p2++
		}
	}
}
