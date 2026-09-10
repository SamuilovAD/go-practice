package main

import "fmt"

func main() {
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	fmt.Printf("Number of islands: %d", numIslands(grid))
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	numberOfIslands := 0
	for row := 0; row < len(grid); row++ {
		for column := 0; column < len(grid[0]); column++ {
			if grid[row][column] == '1' {
				dfs(grid, row, column)
				numberOfIslands += 1
			}
		}
	}
	return numberOfIslands
}

func dfs(grid [][]byte, row int, column int) {
	if row < 0 ||
		column < 0 ||
		row >= len(grid) ||
		column >= len(grid[0]) ||
		grid[row][column] != '1' {
		return
	}
	grid[row][column] = '0'
	dfs(grid, row-1, column)
	dfs(grid, row+1, column)
	dfs(grid, row, column-1)
	dfs(grid, row, column+1)
}
