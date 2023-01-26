package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var (
		counter int64
		//mutex   sync.Mutex
		wg sync.WaitGroup
	)

	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func(x int, wg *sync.WaitGroup) {
			defer wg.Done()
			/*
				defer mutex.Unlock()
				mutex.Lock()

			*/
			atomic.AddInt64(&counter, 1)

		}(i, &wg)
	}
	wg.Wait()
	fmt.Printf("counter=%d\n", counter)
}
