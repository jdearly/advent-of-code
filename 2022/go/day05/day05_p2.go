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
	p2_moves := Moves(fs)

	temp := []string{}
	for _, move := range p2_moves {
		for i := 0; i < move.n; i++ {
			stk, val, err := start[move.from].Pop()
			start[move.from] = stk
			if err != nil {
				log.Fatal(err)
			}
			temp = append(temp, val)
		}

		for len(temp) > 0 {
			v := temp[len(temp)-1]
			temp = temp[:len(temp)-1]
			start[move.to] = start[move.to].Push(v)
		}
		temp = nil
	}

	for _, s := range start {
		stk, val, _ := s.Pop()
		s = stk
		ans += val
	}

	fmt.Println(ans)
}
