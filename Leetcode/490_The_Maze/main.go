package main

import "fmt"

func main() {
	maze := [][]int{
		{0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0},
		{1, 1, 0, 1, 1},
		{0, 0, 0, 0, 0},
	}
	start := []int{0, 4}
	destination := []int{4, 4}

	fmt.Println(hasPathDfs(maze, start, destination)) // true
}

func hasPathDfs(maze [][]int, start []int, destination []int) bool {
	rows, cols := len(maze), len(maze[0])
	visited := make([][]bool, rows)
	for r := range visited {
		visited[r] = make([]bool, cols)
	}
	directions := [][2]int{
		{0, -1}, // Left
		{1, 0},  // Down
		{0, 1},  // Right
		{-1, 0}, // Up
	}
	var dfs func(r, c int) bool
	dfs = func(r, c int) bool {
		if visited[r][c] {
			return false
		}
		if r == destination[0] && c == destination[1] {
			return true
		}
		visited[r][c] = true
		for _, direction := range directions {
			nextR, nextC := r, c
			for {
				nr := nextR + direction[0]
				nc := nextC + direction[1]
				if nr < 0 || nr >= rows || nc < 0 || nc >= cols || maze[nr][nc] == 1 {
					break
				}
				nextR = nr
				nextC = nc
			}
			if dfs(nextR, nextC) {
				return true
			}
		}
		return false
	}
	return dfs(start[0], start[1])
}
