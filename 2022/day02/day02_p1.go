package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func PartOne() {
	rps_scores := map[string]int{"A": 1, "B": 2, "C": 3, "X": 1, "Y": 2, "Z": 3}
	total_score := 0
	// Rock beats scissors
	// Paper beats rock
	// Scissors beats paper
	dat, err := os.Open("input.txt")
	if err != nil {
		panic("invalid input")
	}

	fs := bufio.NewScanner(dat)
	fs.Split(bufio.ScanLines)

	for fs.Scan() {
		round := strings.Split(fs.Text(), " ")

		opponent_score := rps_scores[round[0]]
		my_score := rps_scores[round[1]]

		if opponent_score == my_score {
			total_score += my_score + 3
		} else if round[0] == "A" && round[1] == "Z" {
			total_score += my_score
		} else if round[0] == "B" && round[1] == "X" {
			total_score += my_score
		} else if round[0] == "C" && round[1] == "Y" {
			total_score += my_score
		} else {
			total_score += my_score + 6
		}
	}
	fmt.Println(total_score)
}
