package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

func distance(head, tail point) float64 {
	d := math.Max(math.Abs(float64(head.x-tail.x)), math.Abs(float64(head.y-tail.y)))
	return d
}

func PartOne(dat *os.File) {

	fs := bufio.NewScanner(dat)

	head := point{0, 0}
	tail := point{0, 0}

	seen := map[point]bool{}

	seen[tail] = true

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
			if distance(head, tail) >= 2 {
				if direction == "L" && tail.y != head.y {
					tail.y = head.y
					tail.x--
				} else if direction == "L" {
					tail.x--
				}

				if direction == "R" && tail.y != head.y {
					tail.y = head.y
					tail.x++
				} else if direction == "R" {
					tail.x++
				}

				if direction == "U" && tail.x != head.x {
					tail.x = head.x
					tail.y++
				} else if direction == "U" {
					tail.y++
				}

				if direction == "D" && tail.x != head.x {
					tail.x = head.x
					tail.y--
				} else if direction == "D" {
					tail.y--
				}
			}
			seen[tail] = true
		}
	}

	fmt.Println(len(seen))
}
