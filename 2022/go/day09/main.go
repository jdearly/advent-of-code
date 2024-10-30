package main

import (
	"log"
	"math"
	"os"
)

type point struct {
	id     int
	parent *point
	tail   *point
	x      int
	y      int
}

func distance(head, tail point) float64 {
	d := math.Max(math.Abs(float64(head.x-tail.x)), math.Abs(float64(head.y-tail.y)))
	return d
}

func main() {
	dat, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer dat.Close()

	//PartOne(dat)
	PartTwo(dat)
}
