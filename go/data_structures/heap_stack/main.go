package main

import "fmt"

type Person struct {
	Name string
}

func main() {

	say := "hello world"
	_ = say

	var name = "zhengyansheng"
	p := Person{Name: name}
	_ = p

	Hello()
}

func Hello() {
	msg := "hi nihao"
	fmt.Println(msg)
}
