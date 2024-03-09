package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var numGoRoutines = 10

	wg.Add(numGoRoutines)

	for i := 0; i < numGoRoutines; i++ {

		// 方式1
		go func(i int) {
			defer wg.Done()

			// 业务逻辑
			fmt.Println(i)
		}(i)

		// 方式2
		//go worker(i, &wg)
	}

	// 阻塞等待所有任务完成
	wg.Wait()

	fmt.Println("Hello, playground")
}

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done()

	// 业务逻辑
	fmt.Println(i)
}
