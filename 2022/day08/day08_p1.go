package main

func PartOne(file string) int {
	ans := 0
	grid := ParseGrid(file)

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
