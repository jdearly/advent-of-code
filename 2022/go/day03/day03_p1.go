package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

type set struct {
	m map[rune]bool
}

func NewSet() *set {
	s := &set{}
	s.m = make(map[rune]bool)
	return s
}

func (s *set) Add(v rune) {
	s.m[v] = true
}

func Intersect(a, b *set) *set {
	intersect := NewSet()
	for k := range a.m {
		if b.Contains(k) {
			intersect.Add(k)
		}
	}
	return intersect
}

func (s *set) Contains(v rune) bool {
	if _, ok := s.m[v]; ok {
		return true
	}
	return false
}

func strToSet(str string) *set {
	s := NewSet()

	for _, v := range str {
		s.Add(v)
	}
	return s
}

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
		first := strToSet(str[:len(str)/2])
		second := strToSet(str[len(str)/2:])

		intersect := Intersect(first, second)

		for k := range intersect.m {
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
