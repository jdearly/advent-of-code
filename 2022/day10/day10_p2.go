package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func contains(s []int, val int) bool {
	for _, v := range s {
		if v == val {
			return true
		}
	}
	return false
}

func drawPixel(sprite_window []int, row int) {
	if contains(sprite_window, row) {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}
}

func incrementCrt(crt_row int) int {
	crt_row++
	if crt_row > 40 {
		fmt.Println()
		crt_row = 1
	}
	return crt_row
}

func PartTwo() {
	dat, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fs := bufio.NewScanner(dat)
	fs.Split(bufio.ScanLines)

	crt_row := 1
	register := 1
	sprite_window := []int{register, register + 1, register + 2}

	for fs.Scan() {

		line := fs.Text()
		command := line[:4]
		if command == "noop" {
			drawPixel(sprite_window, crt_row)
			crt_row = incrementCrt(crt_row)
		} else {
			value, _ := strconv.Atoi(line[5:])
			for i := 0; i < 2; i++ {
				drawPixel(sprite_window, crt_row)
				crt_row = incrementCrt(crt_row)
			}
			register = register + value
			sprite_window = []int{register, register + 1, register + 2}
		}
	}
}
