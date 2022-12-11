package main

import (
	"log"
	"os"
)

func main() {
	dat, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer dat.Close()

	PartOne(dat)
	//partTwo(dat)
}
