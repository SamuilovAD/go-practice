package main

import (
	"fmt"
	"sort"
)

func main() {
	slice := []int{4, -2, 2, -2, 3, 6, 8, 3, 3}
	bucketSort(slice)
	fmt.Printf("Bucket sort: %v \n", slice)
	slice = []int{4, -2, 2, -2, 3, 6, 8, 3, 3}
	countingSort(slice)
	fmt.Printf("Counting sort: %v \n", slice)
	slice = []int{5, 4, 1, 2, 6, 8, 3, 3}
	selectionSort(slice)
	fmt.Printf("Selection sort: %v \n", slice)
	slice = []int{5, 4, 1, 2, 6, 8, 3, 3}
	insertionSort(slice)
	fmt.Printf("Insertion sort: %v \n", slice)
	slice = []int{5, 4, 1, 2, 6, 8, 3, 3}
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

// Time: O(n²), Space: (1)
func insertionSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		current := nums[i]
		j := i - 1
		for j >= 0 && nums[j] > current {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = current
	}
}

// Time: O(n²), Space: (1)
func selectionSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		minIndex := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
}

// Time: O(n+k), Space: (k) k = maxValue - minValue + 1
func countingSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	minValue := nums[0]
	maxValue := nums[0]
	for _, val := range nums {
		if minValue > val {
			minValue = val
		}
		if maxValue < val {
			maxValue = val
		}
	}
	counts := make([]int, maxValue-minValue+1)
	for _, val := range nums {
		counts[val-minValue]++
	}
	index := 0
	for value, count := range counts {
		for count > 0 {
			nums[index] = value + minValue
			index++
			count--
		}
	}
}

func bucketSort(nums []int) {
	if len(nums) < 2 {
		return
	}
	minVal, maxVal := nums[0], nums[0]
	for _, num := range nums[1:] {
		if num < minVal {
			minVal = num
		}
		if num > maxVal {
			maxVal = num
		}
	}
	sizeOfBucket := 10
	numberOfBuckets := (maxVal-minVal)/sizeOfBucket + 1
	buckets := make([][]int, numberOfBuckets)
	for _, num := range nums {
		bucketIdx := (num - minVal) / sizeOfBucket
		buckets[bucketIdx] = append(buckets[bucketIdx], num)
	}
	pos := 0
	for _, bucket := range buckets {
		sort.Ints(bucket)
		pos += copy(nums[pos:], bucket)
	}
}
