package main

import "fmt"

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))    // 2
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4})) // 2.5
}

func findMedianSortedArrays(nums1, nums2 []int) float64 {
	// Ensure binary search is performed on the smaller array
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	arrayA := nums1
	arrayB := nums2

	lengthA := len(arrayA)
	lengthB := len(arrayB)

	// Total size of the left half of the merged arrays
	totalLeftSize := (lengthA + lengthB + 1) / 2

	// Search boundaries for how many elements to take from arrayA into the left half
	leftPartitionMin := 0
	leftPartitionMax := lengthA

	// "Infinity" placeholders outside the valid range of numbers
	const NEG_INF = -(1_000_000 + 1)
	const POS_INF = 1_000_000 + 1

	for leftPartitionMin <= leftPartitionMax {
		// Number of elements taken from arrayA for the left half
		leftSizeA := (leftPartitionMin + leftPartitionMax) / 2
		// Remaining elements for the left half come from arrayB
		leftSizeB := totalLeftSize - leftSizeA

		// Max value in the left half from arrayA
		leftMaxA := NEG_INF
		if leftSizeA > 0 {
			leftMaxA = arrayA[leftSizeA-1]
		}
		// Min value in the right half from arrayA
		rightMinA := POS_INF
		if leftSizeA < lengthA {
			rightMinA = arrayA[leftSizeA]
		}

		// Max value in the left half from arrayB
		leftMaxB := NEG_INF
		if leftSizeB > 0 {
			leftMaxB = arrayB[leftSizeB-1]
		}
		// Min value in the right half from arrayB
		rightMinB := POS_INF
		if leftSizeB < lengthB {
			rightMinB = arrayB[leftSizeB]
		}

		// Check if we have a valid partition
		if leftMaxA <= rightMinB && leftMaxB <= rightMinA {
			if (lengthA+lengthB)%2 == 1 {
				// Odd total length → median is the largest in the left half
				if leftMaxA > leftMaxB {
					return float64(leftMaxA)
				}
				return float64(leftMaxB)
			}
			// Even total length → median is the average of max(left half) and min(right half)
			maxLeft := leftMaxA
			if leftMaxB > maxLeft {
				maxLeft = leftMaxB
			}
			minRight := rightMinA
			if rightMinB < minRight {
				minRight = rightMinB
			}
			return (float64(maxLeft) + float64(minRight)) / 2.0
		}

		// Adjust binary search range
		if leftMaxA > rightMinB {
			// Took too many elements from arrayA → move left
			leftPartitionMax = leftSizeA - 1
		} else {
			// Took too few elements from arrayA → move right
			leftPartitionMin = leftSizeA + 1
		}
	}

	// By problem constraints, execution should never reach here
	return 0
}
