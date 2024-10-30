package main

func checkLeft(grid [][]int, i, j int) (bool, int) {
	c := 0
	for col := j - 1; col >= 0; col-- {
		c++
		if grid[i][col] >= grid[i][j] {
			return false, c
		}
	}
	return true, c
}

func checkRight(grid [][]int, i, j int) (bool, int) {
	c := 0
	for col := j + 1; col < len(grid); col++ {
		c++
		if grid[i][col] >= grid[i][j] {
			return false, c
		}
	}
	return true, c
}

func checkFromTop(grid [][]int, i, j int) (bool, int) {
	c := 0
	for row := i - 1; row >= 0; row-- {
		c++
		if grid[row][j] >= grid[i][j] {
			return false, c
		}
	}
	return true, c
}

func checkFromBottom(grid [][]int, i, j int) (bool, int) {
	c := 0
	for row := i + 1; row < len(grid); row++ {
		c++
		if grid[row][j] >= grid[i][j] {
			return false, c
		}
	}
	return true, c
}
