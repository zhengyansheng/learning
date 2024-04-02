package main

import (
	"context"
	"fmt"
	"time"
)

func RunWithDeadline() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	defer cancel()

	go worker(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("exec timeout ")
	case <-time.After(3 * time.Second):
		fmt.Println("exec complete")
	}

	//<-time.NewTimer(1 * time.Second).C
}

func RunWithTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	go worker(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("exec timeout ")
	case <-time.After(7 * time.Second):
		fmt.Println("exec complete")
		cancel()
	}

	<-time.NewTimer(1 * time.Second).C
}

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker done")
			return
		default:
			fmt.Println("working")
			time.Sleep(1 * time.Second)
		}
	}

}
