package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func PartTwo() {
	rps_scores := map[string]int{"A": 1, "B": 2, "C": 3}
	lose_conditions := map[string]string{"A": "C", "B": "A", "C": "B"}
	win_conditions := map[string]string{"C": "A", "A": "B", "B": "C"}

	total_score := 0
	dat, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	fs := bufio.NewScanner(dat)
	fs.Split(bufio.ScanLines)

	for fs.Scan() {
		round := strings.Split(fs.Text(), " ")

		opponent_score := rps_scores[round[0]]
		outcome := round[1]

		if outcome == "Y" {
			total_score += opponent_score + 3
		} else if outcome == "X" {
			total_score += rps_scores[lose_conditions[round[0]]]
		} else if outcome == "Z" {
			total_score += rps_scores[win_conditions[round[0]]] + 6
		}
	}
	fmt.Println(total_score)
}
