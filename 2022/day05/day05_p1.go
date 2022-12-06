package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type stack []string

type move struct {
	n    int
	from int
	to   int
}

func (s stack) Push(v string) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, string, error) {
	l := s.Size()
	if l == 0 {
		return nil, "", errors.New("empty stack")
	}
	res := s[l-1]
	s = s[:l-1]
	return s, res, nil
}

func (s stack) Size() int {
	return len(s)
}

func reverse(stk []stack) []stack {
	res := []stack{}
	for len(stk) != 0 {
		res = append(res, stk[len(stk)-1])
		stk = stk[:len(stk)-1]
	}
	return res
}

func StartingStacks(fs *bufio.Scanner) []stack {
	lines := [][]string{}
	for fs.Scan() {
		if fs.Text() == "" {
			break
		}
		line := fs.Text()
		new_line := []string{}
		for i := 1; i < len(line); i++ {
			new_line = append(new_line, string(line[i]))
			i = i + 3
		}
		lines = append(lines, new_line)
	}

	lines = lines[:len(lines)-1]
	stacks := make([]stack, 8)
	for i, l := range lines {
		for j := range l {
			stacks[i] = stacks[i].Push(l[j])
		}
	}

	stacks = reverse(stacks)
	starting_stacks := []stack{}
	temp := stack{}

	for i := 0; i < 9; i++ {
		for j := 0; j < 8; j++ {
			if stacks[j][i] != " " {
				temp = append(temp, stacks[j][i])
			}
		}
		starting_stacks = append(starting_stacks, temp)
		temp = nil
	}

	return starting_stacks
}

func Moves(fs *bufio.Scanner) []move {
	res := []move{}
	for fs.Scan() {
		str := fs.Text()
		str_slice := strings.Split(str, " ")
		num_moves, _ := strconv.Atoi(fmt.Sprintf("%s", str_slice[1]))
		from, _ := strconv.Atoi(fmt.Sprintf("%s", str_slice[3]))
		to, _ := strconv.Atoi(fmt.Sprintf("%s", str_slice[5]))
		res = append(res, move{num_moves, from, to})
	}
	return res
}

func PartOne() {
	ans := ""
	dat, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fs := bufio.NewScanner(dat)
	fs.Split(bufio.ScanLines)

	start := StartingStacks(fs)
	moves := Moves(fs)

	for _, move := range moves {
		for i := 0; i < move.n; i++ {
			s, val, err := start[move.from-1].Pop()
			start[move.from-1] = s
			if err != nil {
				log.Fatal(err)
			}
			start[move.from-1] = start[move.from-1][:len(start[move.from-1])]
			start[move.to-1] = start[move.to-1].Push(val)
		}
	}

	for _, s := range start {
		stk, val, _ := s.Pop()
		s = stk
		ans += val
	}

	fmt.Println(ans)
}
