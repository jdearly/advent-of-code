package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func PartOne() {
	dat, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer dat.Close()

	fs := bufio.NewScanner(dat)
	fs.Split(bufio.ScanLines)

	var total = 0
	var stack []directory

	stack = append(stack, directory{name: "/"})

	for fs.Scan() {
		line := fs.Text()

		if strings.HasPrefix(line, "dir") || line == "$ cd /" || line == "$ ls" {
			continue
		}

		fmt.Println(stack)
		if strings.HasPrefix(line, "$ cd") {
			if line[5:] == ".." {
				s := stack[len(stack)-1].size
				if s <= 100_000 {
					total += s
				}
				stack = stack[:len(stack)-1]
				stack[len(stack)-1].size += s
			} else {
				dir := directory{}
				dir.name = line[5:]
				stack = append(stack, dir)
			}
		} else {
			size := strings.Split(line, " ")
			s, _ := strconv.Atoi(size[0])
			stack[len(stack)-1].size += s
		}
	}
	fmt.Println(total)
}
