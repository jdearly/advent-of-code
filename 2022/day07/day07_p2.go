package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func processDir(stack *[]directory, directory_sizes *[]int) {
	s := (*stack)[len((*stack))-1].size
	*directory_sizes = append((*directory_sizes), s)
	*stack = (*stack)[:len((*stack))-1]
	if len((*stack)) > 0 {
		(*stack)[len((*stack))-1].size += s
	}

}

func PartTwo() {
	dat, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer dat.Close()

	fs := bufio.NewScanner(dat)
	fs.Split(bufio.ScanLines)

	var stack []directory
	total_space := 70_000_000
	space_needed := 30_000_000
	directory_sizes := []int{}
	stack = append(stack, directory{name: "/"})

	for fs.Scan() {
		line := fs.Text()

		if strings.HasPrefix(line, "dir") || line == "$ cd /" || line == "$ ls" {
			continue
		}

		if strings.HasPrefix(line, "$ cd") {
			if line[5:] == ".." {
				var s *[]directory = &stack
				var ds *[]int = &directory_sizes
				processDir(s, ds)
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

	for len(stack) != 0 {
		var s *[]directory = &stack
		var ds *[]int = &directory_sizes
		processDir(s, ds)
	}

	sort.Ints(directory_sizes)
	min_dir_size := space_needed - (total_space - directory_sizes[len(directory_sizes)-1])

	for _, v := range directory_sizes {
		if v >= min_dir_size {
			fmt.Println(v)
			break
		}
	}
}
