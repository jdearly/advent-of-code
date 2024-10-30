package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func PartOne() {
	dat, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fs := bufio.NewScanner(dat)
	fs.Split(bufio.ScanLines)

	m := map[int]int{}
	cycle := 1
	register := 1
	for fs.Scan() {

		line := fs.Text()
		command := line[:4]
		if command == "noop" {
			cycle += 1
			m[cycle] = register
		} else {
			value, _ := strconv.Atoi(line[5:])
			for i := 0; i < 2; i++ {
				cycle += 1
				m[cycle] = register
			}
			register = register + value
			m[cycle] = register
		}
	}
	result := 20*m[20] + 60*m[60] + 100*m[100] + 140*m[140] + 180*m[180] + 220*m[220]
	fmt.Println(result)
}
