package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var counter int32

	go func() {
		for {
			fmt.Println("goroutine 1")
			<-time.NewTimer(time.Second).C
			atomic.AddInt32(&counter, 1)
		}
	}()

	go func() {
		// recover 在defer延迟执行，goroutine就结束了
		defer func() {
			if recover() != nil {
				fmt.Println("-----> recover, counter > 10")
			}
		}()
		for {
			if counter > 10 {
				panic("counter > 10")
			}
			fmt.Println("goroutine 2")
			<-time.NewTimer(time.Second * 2).C
			atomic.AddInt32(&counter, 1)
		}
	}()

	go func() {
		for {
			fmt.Println("goroutine 3")
			<-time.NewTimer(time.Second * 2).C
			atomic.AddInt32(&counter, 1)
		}
	}()

	for {
		fmt.Println("main goroutine")
		<-time.NewTimer(time.Second * 1).C
	}
}
