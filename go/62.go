package main

func uniquePaths(m int, n int) int {
	grid := make([][]int, m)

	for idx := range grid {
		grid[idx] = make([]int, n)
	}

	grid[0][0] = 1

	for row := range m {
		for column := range n {
			if row > 0 && column > 0 {
				grid[row][column] = grid[row - 1][column] + grid[row][column - 1]
			} else if row > 0 {
				grid[row][column] = grid[row - 1][column]
			} else if column > 0 {
				grid[row][column] = grid[row][column - 1]
			}
		}
	}

	return grid[m - 1][n - 1]
}