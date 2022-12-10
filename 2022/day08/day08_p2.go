package main

import "math"

func PartTwo(file string) float64 {
	scenic_score := 0.0
	grid := ParseGrid(file)

	// For each tree, walk in from edge to check visibility from each
	// direction.
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			_, left_score := checkLeft(grid, i, j)
			_, right_score := checkRight(grid, i, j)
			_, top_score := checkFromTop(grid, i, j)
			_, bottom_score := checkFromBottom(grid, i, j)
			scenic_score = math.Max(scenic_score, float64(left_score*right_score*top_score*bottom_score))
		}
	}

	return scenic_score
}
