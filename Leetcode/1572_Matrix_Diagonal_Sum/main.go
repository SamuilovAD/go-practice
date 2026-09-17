package main

import "fmt"

func main() {
	mat := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Print(diagonalSum(mat))
}
func diagonalSum(mat [][]int) int {
	n := len(mat)
	resultSum := 0
	for i := 0; i < n; i++ {
		resultSum += mat[i][i] + mat[n-1-i][i]
	}
	if n%2 != 0 {
		resultSum -= mat[n/2][n/2]
	}
	return resultSum
}
