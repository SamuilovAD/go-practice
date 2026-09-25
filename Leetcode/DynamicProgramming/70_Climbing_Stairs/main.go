package main

import "fmt"

func main() {
	fmt.Println(climbStairsBruteForce(5))
	fmt.Println(climbStairsDP(8))
}

// Time O(2^n) Memory O(n)
func climbStairsBruteForce(n int) int {
	var climbStairsRecursion func(i, n int) int
	climbStairsRecursion = func(i, n int) int {
		if i > n {
			return 0
		}
		if i == n {
			return 1
		}
		return climbStairsRecursion(i+1, n) + climbStairsRecursion(i+2, n)
	}

	return climbStairsRecursion(0, n)
}

// Time O(n) Memory O(n)
func climbStairsDP(n int) int {
	if n <= 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for step := 3; step <= n; step++ {
		dp[step] = dp[step-1] + dp[step-2]
	}
	return dp[n]
}
