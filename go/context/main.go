package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	timeout := time.Second * 10
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println(1)
			time.Sleep(time.Second)
		}
	}

}
