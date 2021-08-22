package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		// 方式1
		for ch := range callTask() {
			fmt.Println(ch)
		}

	*/

	// 方式2 超时控制
	select {
	case v := <-callTask():
		fmt.Printf("call task response: %v\n", v)
	case <-time.After(time.Second * 3):
		fmt.Println("Timeout")
		//default:
		//	fmt.Println("not do something")
	}
}

func execTask() string {
	time.Sleep(time.Second * 2)
	fmt.Println("task exec end")
	return "Done"
}

func callTask() chan string {
	ch := make(chan string, 1)

	go func() {
		ch <- execTask()
		fmt.Println("call Task end")
		close(ch) // 关闭channel
	}()

	return ch
}
