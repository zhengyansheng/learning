package main

type PopProcessFunc func(x, y int) int

func ProcessFunc(x, y int) int {
	return x + y
}

func logic(process PopProcessFunc) {
	process(1, 2)
}

func main() {
	logic(PopProcessFunc(ProcessFunc))
}
