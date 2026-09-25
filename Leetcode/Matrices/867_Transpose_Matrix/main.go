package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(matrix[0])
	fmt.Println(matrix[1])
	fmt.Println(matrix[2])
	fmt.Println("======")
	transposed := transpose(matrix)
	fmt.Println(transposed[0])
	fmt.Println(transposed[1])
	fmt.Println(transposed[2])
}
func transpose(matrix [][]int) [][]int {
	rows := len(matrix)
	columns := len(matrix[0])
	transposedMatrix := make([][]int, columns)
	for i := 0; i < columns; i++ {
		transposedMatrix[i] = make([]int, rows)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			transposedMatrix[j][i] = matrix[i][j]
		}
	}
	return transposedMatrix
}
