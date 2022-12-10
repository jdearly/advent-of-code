package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ParseGrid(file string) [][]int {
	dat, err := os.Open(file)

	if err != nil {
		log.Fatal(err)
	}

	fs := bufio.NewScanner(dat)
	fs.Split(bufio.ScanLines)

	var grid = [][]int{}

	for fs.Scan() {
		line := fs.Text()
		row := []int{}
		for i := 0; i < len(line); i++ {
			val, _ := strconv.Atoi(string(line[i]))
			row = append(row, val)
		}

		grid = append(grid, row)
	}

	return grid
}
