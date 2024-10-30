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

	seen := map[point]bool{}
	head := point{0, nil, nil, 0, 0}
	curr := &head

	for i := 1; i < 10; i++ {
		next := point{i, curr, nil, 0, 0}
		curr.tail = &next
		curr = &next
		if i == 9 {
			seen[*curr] = true
		}
	}

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

			knot := head
			// this is all wrong..
			for knot.tail != nil {
				if distance(knot, *knot.tail) > 1 {
					prev_x := knot.x
					prev_y := knot.y
					switch direction {
					case "U":
						if knot.x != knot.tail.x {
							knot.tail.x = knot.x
						}
						prev_y, knot.tail.y = knot.tail.y, prev_y
					case "D":
						if knot.x != knot.tail.x {
							knot.tail.x = knot.x
						}
						prev_y, knot.tail.y = knot.tail.y, prev_y
					case "L":
						if knot.y != knot.tail.y {
							knot.tail.y = knot.y
						}
						prev_x, knot.tail.x = knot.tail.x, prev_x
					case "R":
						if knot.y != knot.tail.y {
							knot.tail.y = knot.y
						}
						prev_x, knot.tail.x = knot.tail.x, prev_x
					}
				}
				knot = *knot.tail
			}
			if knot.id == 9 {
				seen[knot] = true
			}
		}
	}
	fmt.Println(len(seen))
}
