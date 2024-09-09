package goroutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGoroutineNums(t *testing.T) {
	chPool := make(chan struct{}, 8)
	done := make(chan struct{})

	var wg sync.WaitGroup

	go func() {
		for {
			select {
			case <-done:
				return
			default:
				t.Logf("num goroutine: %d", runtime.NumGoroutine())
				time.Sleep(time.Second * 1)
			}
		}
	}()

	for i := 0; i < 50; i++ {
		wg.Add(1)
		chPool <- struct{}{}
		go worker(chPool, i, &wg)
	}

	wg.Wait()
	done <- struct{}{}
	t.Log("Done")
}

func worker(ch chan struct{}, i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(i)
	time.Sleep(time.Second)
	<-ch
}
