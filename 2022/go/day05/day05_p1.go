package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func PartOne() {
	ans := ""
	dat, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer dat.Close()

	fs := bufio.NewScanner(dat)
	fs.Split(bufio.ScanLines)

	start := StartingStacks(fs)
	moves := Moves(fs)

	for _, move := range moves {
		for i := 0; i < move.n; i++ {
			s, val, err := start[move.from].Pop()
			start[move.from] = s
			if err != nil {
				log.Fatal(err)
			}
			start[move.to] = start[move.to].Push(val)
		}
	}

	for _, s := range start {
		stk, val, _ := s.Pop()
		s = stk
		ans += val
	}

	fmt.Println(ans)
}
