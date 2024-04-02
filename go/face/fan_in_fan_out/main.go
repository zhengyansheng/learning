package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup, jobs <-chan int, resultCh chan<- int) {
	defer wg.Done()
	for job := range jobs {
		resultCh <- job * 2
	}
}

func main() {
	/*
		有100个任务，10个worker，每个worker处理5个任务，最后将结果输出
	*/
	var startTime = time.Now()
	var wg sync.WaitGroup
	tasksNumber := 100
	workersNumber := 5

	taskCh := make(chan int, 5)
	resultCh := make(chan int, workersNumber)
	//defer close(resultCh)

	// fan out
	for i := 0; i < workersNumber; i++ {
		wg.Add(1)
		go worker(i, &wg, taskCh, resultCh)
	}

	// fan in
	go func() {
		// 100个任务
		for i := 0; i < tasksNumber; i++ {
			taskCh <- i
		}
		close(taskCh)
	}()

	go func() {
		for ch := range resultCh {
			fmt.Printf("Result: %d\n", ch)
		}
		fmt.Println("result channel done!")
	}()

	// 等待所有worker处理完毕
	wg.Wait()

	close(resultCh)

	fmt.Printf("Time: %v\n", time.Since(startTime))
	time.Sleep(1 * time.Second)

}
