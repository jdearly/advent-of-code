package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func PartTwo(file string) {
	ans := 0

	dat, err := os.Open(file)

	defer dat.Close()

	if err != nil {
		log.Fatal(err)
	}

	fs := bufio.NewScanner(dat)
	fs.Split(bufio.ScanLines)

	for fs.Scan() {

		line := fs.Text()
		pairs := strings.Split(line, ",")
		first := convert(strings.Split(pairs[0], "-"))
		second := convert(strings.Split(pairs[1], "-"))

		if first[0] <= second[1] && second[0] <= first[1] {
			ans++
		}
	}

	fmt.Println(ans)
}
