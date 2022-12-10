package main

func checkLeft(grid [][]int, i, j int) bool {
	for col := j - 1; col >= 0; col-- {
		if grid[i][col] >= grid[i][j] {
			return false
		}
	}
	return true
}

func checkRight(grid [][]int, i, j int) bool {
	for col := j + 1; col < len(grid); col++ {
		if grid[i][col] >= grid[i][j] {
			return false
		}
	}
	return true
}

func checkFromTop(grid [][]int, i, j int) bool {
	for row := i - 1; row >= 0; row-- {
		if grid[row][j] >= grid[i][j] {
			return false
		}
	}
	return true
}

func checkFromBottom(grid [][]int, i, j int) bool {
	for row := i + 1; row < len(grid); row++ {
		if grid[row][j] >= grid[i][j] {
			return false
		}
	}
	return true
}
