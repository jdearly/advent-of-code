package main

type window struct {
	start int
	end   int
}

func (w *window) slide() {
	w.start++
	w.end++
}

func main() {
	PartOne()
	PartTwo()
}
