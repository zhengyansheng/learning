package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer func() {
		fmt.Println("Main goroutine end")
		cancel()
	}()

	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			fmt.Println("child 2 goroutine")
		case <-time.After(time.Second * 5):
			fmt.Println("child timeout")
		}
	}(ctx)

	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			fmt.Println("child 1 goroutine")
		case <-time.After(time.Second * 5):
			fmt.Println("child timeout")
		}
	}(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("Main: ", ctx.Err())
		return
	}
}

func main2() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer func() {
		fmt.Println("Main goroutine end")
		cancel()
	}()

	select {
	case <-ctx.Done():
		fmt.Println("normal")
	case <-time.After(time.Second * 1):
		fmt.Println("timeout")
	}
}
