package main

import (
	"fmt"
)

var seen [][]bool

func main() {
	grid := [][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}

	fmt.Println(maxAreaOfIsland(grid))

}

func maxAreaOfIsland(grid [][]int) int {
	var max_area int = 0
	rows := len(grid)
	columns := len(grid[0])

	seen = make([][]bool, rows)
	for i := 0; i < rows; i++ {
		seen[i] = make([]bool, columns)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			_area := area(grid, i, j)
			if _area > max_area {
				max_area = _area
			}
		}
	}
	return max_area
}

func area(grid [][]int, sr int, sc int) int {
	if sr < 0 || sc < 0 || sr >= len(grid) || sc >= len(grid[sr]) || seen[sr][sc] == true || grid[sr][sc] == 0 {
		return 0
	}
	seen[sr][sc] = true
	return 1 + area(grid, sr+1, sc) + area(grid, sr-1, sc) +
		area(grid, sr, sc+1) + area(grid, sr, sc-1)
}
