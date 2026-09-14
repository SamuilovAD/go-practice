package main

func main() {
	numArray := Constructor([]int{-2, 0, 3, -5, 2, -1})
	numArray.SumRangeBruteForce(0, 2)
}

type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	return NumArray{
		nums: nums,
	}
}

// Time: O(N) Space:(1)
func (this *NumArray) SumRangeBruteForce(left int, right int) int {
	sum := 0
	for i := left; i <= right; i++ {
		sum += this.nums[i]
	}
	return sum
}

type NumArrayPrefixSum struct {
	prefixSum []int
}

func ConstructorPrefixSum(nums []int) NumArrayPrefixSum {
	prefixSum := make([]int, len(nums)+1)
	for i, num := range nums {
		prefixSum[i+1] = prefixSum[i] + num
	}

	return NumArrayPrefixSum{
		prefixSum: prefixSum,
	}
}

// Time:read O(1), write O(N), Space: O(N)
func (n *NumArrayPrefixSum) SumRange(left int, right int) int {
	return n.prefixSum[right+1] - n.prefixSum[left]
}
