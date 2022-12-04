package main

import (
	"log"
	"strconv"
)

func convert(pairs []string) []int {
	pair_int := make([]int, len(pairs))
	for i, v := range pairs {
		val, err := strconv.Atoi(v)
		pair_int[i] = val
		if err != nil {
			log.Fatal(err)
		}
	}
	return pair_int
}

func main() {
	file := "input.txt"
	PartOne(file)
	PartTwo(file)
}
