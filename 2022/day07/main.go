package main

type directory struct {
	size     int
	name     string
	parent   *directory
	children []directory
	files    []file
}

type file struct {
	size int32
	name string
}

func main() {
	PartOne()
	PartTwo()
}
