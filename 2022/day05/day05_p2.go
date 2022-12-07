package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func PartTwo() {
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
		temp := stack{}
		for i := 0; i < move.n; i++ {
			s, val, err := start[move.from].Pop()
			start[move.from] = s
			if err != nil {
				log.Fatal(err)
			}
			temp = temp.Push(val)
		}
		for temp != nil {
			t, val, _ := temp.Pop()
			temp = t
			start[move.to] = start[move.to].Push(val)
		}
		temp = nil
		fmt.Println(start[move.to])
	}

	for _, s := range start {
		stk, val, _ := s.Pop()
		s = stk
		ans += val
	}

	fmt.Println(ans)
}
