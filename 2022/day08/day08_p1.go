package main

func checkLeft(grid [][]int, i, j int) bool {
	var ans bool
	for col := 0; col < len(grid); col++ {
		if col == j {
			ans = true
			break
		}

		if grid[i][col] >= grid[i][j] {
			return false
		}
	}
	return ans
}

func checkRight(grid [][]int, i, j int) bool {
	var ans bool
	for col := len(grid) - 1; col >= 0; col-- {
		if col == j {
			ans = true
			break
		}

		if grid[i][col] >= grid[i][j] {
			return false
		}
	}
	return ans
}

func checkFromTop(grid [][]int, i, j int) bool {
	var ans bool
	for row := 0; row < len(grid); row++ {
		if row == i {
			ans = true
			break
		}
		if grid[row][j] >= grid[i][j] {
			return false
		}
	}
	return ans
}

func checkFromBottom(grid [][]int, i, j int) bool {
	var ans bool
	for row := len(grid) - 1; row >= 0; row-- {
		if row == i {
			ans = true
			break
		}
		if grid[row][j] >= grid[i][j] {
			return false
		}
	}
	return ans
}

func PartOne() int {
	ans := 0
	grid := ParseGrid("input.txt")

	// For each tree, walk in from edge to check visibility from each
	// direction.
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			if checkLeft(grid, i, j) || checkRight(grid, i, j) || checkFromTop(grid, i, j) || checkFromBottom(grid, i, j) {
				ans++
			}
		}
	}
	return ans
}
