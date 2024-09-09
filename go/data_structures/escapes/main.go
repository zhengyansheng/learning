package main

func main() {
	add(1, 2)
}

func add(x, y int) *int {
	res := 0
	res = x + y
	return &res
}
