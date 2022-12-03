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

func PartTwo() {
	dat, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fs := bufio.NewScanner(dat)
	fs.Split(bufio.ScanLines)

	total := 0
	group := []string{}

	for fs.Scan() {
		str := fs.Text()

		group = append(group, str)

		if len(group) == 3 {
			intersect := map[rune]bool{}

			for _, v := range group[0] {
				if strings.ContainsRune(group[1], v) &&
					strings.ContainsRune(group[2], v) {
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
			group = nil
		}
	}
	fmt.Println(total)
}
