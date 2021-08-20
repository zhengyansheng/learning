package main

import (
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	g := new(errgroup.Group)

	g.Go(func() error {
		time.Sleep(time.Second * 3)
		fmt.Println("task 1")
		return nil
	})

	g.Go(func() error {
		time.Sleep(time.Second * 7)
		fmt.Println("task 2")
		return errors.New("execute task 2 error ")
	})

	g.Go(func() error {
		time.Sleep(time.Second * 4)
		fmt.Println("task 3")
		return nil
	})

	err := g.Wait()
	if err != nil {
		fmt.Printf("task exec response: %v\n", err)
		return
	}
	fmt.Println("End ...")
}
