package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

func PartOne() {
	ans := 0
	dat, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	fs := bufio.NewScanner(dat)
	fs.Split(bufio.ScanLines)

	for fs.Scan() {

		line := fs.Text()
		pairs := strings.Split(line, ",")
		first := convert(strings.Split(pairs[0], "-"))
		second := convert(strings.Split(pairs[1], "-"))

		if first[0] >= second[0] && first[1] <= second[1] ||
			second[0] >= first[0] && second[1] <= first[1] {
			ans++
		}
	}

	fmt.Println(ans)
}
