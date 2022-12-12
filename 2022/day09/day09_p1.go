package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func PartOne(dat *os.File) {

	fs := bufio.NewScanner(dat)

	head := point{0, nil, nil, 0, 0}
	tail := point{1, &head, nil, 0, 0}
	head.tail = &tail

	seen := map[point]bool{}

	seen[*head.tail] = true

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
			if distance(head, *head.tail) >= 2 {
				if direction == "L" && head.tail.y != head.y {
					head.tail.y = head.y
					head.tail.x--
				} else if direction == "L" {
					head.tail.x--
				}

				if direction == "R" && head.tail.y != head.y {
					head.tail.y = head.y
					head.tail.x++
				} else if direction == "R" {
					head.tail.x++
				}

				if direction == "U" && head.tail.x != head.x {
					head.tail.x = head.x
					head.tail.y++
				} else if direction == "U" {
					head.tail.y++
				}

				if direction == "D" && head.tail.x != head.x {
					head.tail.x = head.x
					head.tail.y--
				} else if direction == "D" {
					head.tail.y--
				}
			}
			seen[*head.tail] = true
		}
	}

	fmt.Println(len(seen))
}
