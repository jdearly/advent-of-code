package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func PartTwo(dat *os.File) {

	fs := bufio.NewScanner(dat)

	head := point{0, nil, nil, 0, 0}
	curr := &head

	for i := 1; i <= 9; i++ {
		next := point{i, curr, nil, 0, 0}
		curr.tail = &next
		curr = &next
	}

	seen := map[point]bool{}

	for fs.Scan() {
		line := fs.Text()

		move := strings.Split(line, " ")
		direction := move[0]
		amount, _ := strconv.Atoi(move[1])

		for i := 0; i < amount; i++ {
			switch direction {
			case "U":
				head.y++
			case "D":
				head.y--
			case "L":
				head.x--
			case "R":
				head.x++
			}

			curr := head
			for i := 0; i < 9; i++ {
				if distance(curr, *curr.tail) > 1 {
					if direction == "L" && curr.tail.y != curr.y {
						curr.tail.y = curr.y
						curr.tail.x--
					} else if direction == "L" {
						curr.tail.x--
					}

					if direction == "R" && curr.tail.y != curr.y {
						curr.tail.y = curr.y
						curr.tail.x++
					} else if direction == "R" {
						curr.tail.x++
					}

					if direction == "U" && curr.tail.x != curr.x {
						curr.tail.x = curr.x
						curr.tail.y++
					} else if direction == "U" {
						curr.tail.y++
					}

					if direction == "D" && curr.tail.x != curr.x {
						curr.tail.x = curr.x
						curr.tail.y--
					} else if direction == "D" {
						curr.tail.y--
					}
				}
				curr = *curr.tail
			}
			seen[curr] = true
			//		fmt.Println(len(seen))
		}
	}
	fmt.Println(len(seen))
}
