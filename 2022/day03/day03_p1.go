package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func PartOne() {
	dat, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fs := bufio.NewScanner(dat)
	fs.Split(bufio.ScanLines)

	total := 0
	for fs.Scan() {
		str := fs.Text()
		first_half := str[:len(str)/2]
		second_half := str[len(str)/2:]

		intersect := map[rune]bool{}

		for _, v := range first_half {
			if strings.ContainsRune(second_half, v) {
				intersect[v] = true
			}
		}

		for k := range intersect {
			val, _ := strconv.Atoi(fmt.Sprintf("%d", k))
			if unicode.IsLower(k) {
				total += val - 96
			} else {
				total += val - 38
			}
		}
	}
	fmt.Println(total)
}
