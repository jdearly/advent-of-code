package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func PartTwo() {
	dat, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer dat.Close()

	fs := bufio.NewScanner(dat)
	fs.Split(bufio.ScanLines)

	var curr = 0
	var calories []int
	for fs.Scan() {
		if fs.Text() != "" {
			val, _ := strconv.Atoi(fs.Text())
			curr += val
		} else {
			calories = append(calories, curr)
			curr = 0
		}
	}

	sort.Slice(calories, func(i, j int) bool {
		return calories[i] < calories[j]
	})

	var ans = 0
	for i := 1; i <= 3; i++ {
		ans += calories[len(calories)-i]
	}
	fmt.Println(ans)
}
