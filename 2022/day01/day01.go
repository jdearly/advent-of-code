package main

import (
	"fmt"
	"os"
)

func main() {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Print(string(dat))
	//TODO: the actual problem
}
