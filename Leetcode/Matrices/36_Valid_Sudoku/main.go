package main

import "fmt"

func main() {
	board := [][]byte{
		[]byte("........."),
		[]byte("........."),
		[]byte("........."),
		[]byte("...7....."),
		[]byte("........."),
		[]byte("........."),
		[]byte("........."),
		[]byte("........."),
		[]byte("........."),
	}

	fmt.Println(isValidSudoku(board)) // true

	// Both cells, (3, 3) and (4, 4), belong to box 4.
	board[4][4] = '7'
	fmt.Println(isValidSudoku(board)) // false
}

// Time: O(n^2) Space: O(n^2)
func isValidSudoku(board [][]byte) bool {
	const size = 9
	rows := make([]map[byte]bool, size)
	cols := make([]map[byte]bool, size)
	boxes := make([]map[byte]bool, size)
	for i := 0; i < size; i++ {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		boxes[i] = make(map[byte]bool)
	}
	for r := 0; r < size; r++ {
		for c := 0; c < size; c++ {
			val := board[r][c]
			if val == '.' {
				continue
			}
			idx := (r/3)*3 + c/3
			if rows[r][val] || cols[c][val] || boxes[idx][val] {
				return false
			}
			rows[r][val] = true
			cols[c][val] = true
			boxes[idx][val] = true
		}
	}

	return true
}
