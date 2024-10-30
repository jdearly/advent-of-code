package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func PartOne() {
	dat, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer dat.Close()

	fs := bufio.NewScanner(dat)
	fs.Split(bufio.ScanLines)

	var curr = 0
	var max = 0.0
	for fs.Scan() {
		if fs.Text() != "" {
			val, _ := strconv.Atoi(fs.Text())
			curr += val
		} else {
			max = math.Max(max, float64(curr))
			curr = 0
		}
	}
	fmt.Println(max)
}
