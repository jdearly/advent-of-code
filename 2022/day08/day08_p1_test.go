package main

import (
	"testing"
)

func TestTreeVisible(t *testing.T) {
	t.Run("test top-left edge 3", func(t *testing.T) {
		grid := ParseGrid("test.txt")
		got := checkLeft(grid, 0, 0) && checkFromTop(grid, 0, 0) && !checkRight(grid, 0, 0) && !checkFromBottom(grid, 0, 0)
		want := true

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}

	})

	t.Run("test top edge 0", func(t *testing.T) {
		grid := ParseGrid("test.txt")
		got := !checkLeft(grid, 0, 1) && checkFromTop(grid, 0, 1) && !checkRight(grid, 0, 1) && !checkFromBottom(grid, 0, 1)
		want := true

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}

	})

	t.Run("test top edge 3", func(t *testing.T) {
		grid := ParseGrid("test.txt")
		got := !checkLeft(grid, 0, 4) && checkFromTop(grid, 0, 4) && checkRight(grid, 0, 4) && !checkFromBottom(grid, 0, 4)
		want := true

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}

	})

	t.Run("test bottom edge 9", func(t *testing.T) {
		grid := ParseGrid("test.txt")
		got := checkLeft(grid, 4, 3) && checkFromTop(grid, 4, 3) && checkRight(grid, 4, 3) && checkFromBottom(grid, 4, 3)
		want := true

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}

	})

	t.Run("test bottom 5", func(t *testing.T) {
		grid := ParseGrid("test.txt")
		got := checkLeft(grid, 3, 2) && !checkFromTop(grid, 3, 2) && !checkRight(grid, 3, 2) && checkFromBottom(grid, 3, 2)
		want := true

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}

	})

	t.Run("test top left 5", func(t *testing.T) {
		grid := ParseGrid("test.txt")
		got := checkLeft(grid, 1, 1) && checkFromTop(grid, 1, 1) && !checkRight(grid, 1, 1) && !checkFromBottom(grid, 1, 1)
		want := true

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}

	})

	t.Run("test top middle 5", func(t *testing.T) {
		grid := ParseGrid("test.txt")
		got := checkRight(grid, 1, 2) && checkFromTop(grid, 1, 2) && !checkLeft(grid, 1, 2) && !checkFromBottom(grid, 1, 2)
		want := true

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}

	})

	t.Run("test left middle 5", func(t *testing.T) {
		grid := ParseGrid("test.txt")
		got := checkRight(grid, 2, 1) && !checkFromTop(grid, 2, 1) && !checkLeft(grid, 2, 1) && !checkFromBottom(grid, 2, 1)
		want := true

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}

	})

	t.Run("test center 3", func(t *testing.T) {
		grid := ParseGrid("test.txt")
		got := !checkLeft(grid, 2, 2) && !checkFromTop(grid, 2, 2) && !checkRight(grid, 2, 2) && !checkFromBottom(grid, 2, 2)
		want := true

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}

	})

	t.Run("total count", func(t *testing.T) {
		got := PartOne()
		want := 21

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

}
