package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(3)
	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 3)
		fmt.Println("task 1")
	}()

	go func() error {
		defer wg.Done()
		time.Sleep(time.Second * 7)
		fmt.Println("task 2")
		return errors.New("execute task 2 error ")

	}()

	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 4)
		fmt.Println("task 3")
	}()

	wg.Wait()
	fmt.Println("End ...")
}
