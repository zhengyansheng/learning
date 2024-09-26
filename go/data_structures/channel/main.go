package main

import (
	"log"
)

func main() {
	ch := producer()

	log.Println(<-ch)

}

func producer() chan int {
	ch := make(chan int)

	go func() {
		ch <- 1
	}()

	return ch

}

/*
go build -gcflags="-m -m -l" .\main.go
*/
