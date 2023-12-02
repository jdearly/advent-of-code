package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func PartOne() {
	dat, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer dat.Close()

	fs := bufio.NewScanner(dat)
	fs.Split(bufio.ScanLines)
	ans := 0
	for fs.Scan() {
		text := fs.Text()
		var first rune
		var last rune
		for _, c := range text {
			if unicode.IsDigit(c) {
				first = c
				break
			}
		}
		for i := len(text) - 1; i >= 0; i-- {
			if unicode.IsDigit(rune(text[i])) {
				last = rune(text[i])
				break
			}
		}
		numberAsString := fmt.Sprintf("%c%c", first, last)
		num, _ := strconv.Atoi(numberAsString)
		ans += num
	}
	fmt.Println(ans)
}
