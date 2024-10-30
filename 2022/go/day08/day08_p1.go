package main

func PartOne(file string) int {
	ans := 0
	grid := ParseGrid(file)

	// For each tree, walk in from edge to check visibility from each
	// direction.
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			left, _ := checkLeft(grid, i, j)
			right, _ := checkRight(grid, i, j)
			top, _ := checkFromTop(grid, i, j)
			bottom, _ := checkFromBottom(grid, i, j)
			if left || right || top || bottom {
				ans++
			}
		}
	}
	return ans
}
