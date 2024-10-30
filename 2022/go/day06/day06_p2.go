package main

import (
	"fmt"
	"log"
	"os"
)

func PartTwo() {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input_str := string(dat)
	w := window{0, 14}
	for w.end != len(input_str) {
		m := make(map[rune]bool)
		sub := input_str[w.start:w.end]
		for _, v := range sub {
			m[v] = true
		}
		if len(m) == 14 {
			fmt.Println(w.end)
			break
		}
		w.slide()
	}
}
