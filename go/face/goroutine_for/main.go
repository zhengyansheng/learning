package main

import (
	"fmt"
	"time"
)

func sampleErr() {
	data := make(map[int]int, 10)
	for i := 0; i < 10; i++ {
		data[i] = i
	}
	fmt.Printf("data: %v\n", data)

	for k, v := range data {
		// goroutine 栈没有刷新导致的，导致随机值
		go func() {
			fmt.Printf("k: %d, v: %d\n", k, v)
		}()
	}
	time.Sleep(time.Second * 3)
}

func sampleOk() {
	data := make(map[int]int, 10)
	for i := 0; i < 10; i++ {
		data[i] = i
	}
	fmt.Printf("data: %v\n", data)

	for k, v := range data {
		go func(k, v int) {
			fmt.Printf("k: %d, v: %d\n", k, v)
		}(k, v)
	}
	time.Sleep(time.Second * 3)
}

func main() {
	sampleErr()
	sampleOk()
}
