package main

import (
	"context"
	"fmt"
	"time"
)

func RunWithValue() {
	// front -> backend(app1, app2, app3, app4)
	ctx := context.WithValue(context.Background(), "traceId", "32897297492742")
	go App1(ctx)
	time.Sleep(1 * time.Second)
	go App2(ctx)

	select {}
}

func App1(ctx context.Context) {
	fmt.Println(ctx.Value("traceId"))
}

func App2(ctx context.Context) {
	for {
		fmt.Println(ctx.Value("traceId"))
		time.Sleep(1 * time.Second)
	}
}
