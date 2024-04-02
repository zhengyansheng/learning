package main

import (
	"context"
	"fmt"
	"time"
)

func RunWithCancel() {
	/*
		1. 创建三个goroutine，每个goroutine都是一个死循环
		2. 其中一个goroutine退出，其他的goroutine也要退出
	*/
	var stopCh = make(chan struct{})
	// ctx 上下文context
	// cancel 取消函数
	ctx, cancel := context.WithCancel(context.Background()) // context.Background()
	//defer cancel()

	go worker1(ctx, stopCh)
	go worker2(ctx, stopCh)

	// 阻塞
	<-stopCh
	cancel()

	// 等待一段时间，以确保工作协程接收到取消信号并退出
	<-time.After(1 * time.Second)
}

func worker1(ctx context.Context, stopCh chan struct{}) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker1 done")
			return
		default:
			fmt.Println("worker1 working")
			time.Sleep(1 * time.Second)
		}
	}
}

func worker2(ctx context.Context, stopCh chan struct{}) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker2 done")
			return
		default:
			fmt.Println("worker2 working")
			time.Sleep(5 * time.Second)
			stopCh <- struct{}{}
		}
	}
}
