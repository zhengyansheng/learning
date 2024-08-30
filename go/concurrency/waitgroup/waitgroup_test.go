package waitgroup

import (
	"sync"
	"testing"
)

func TestWaitGroup(t *testing.T) {
	w := sync.WaitGroup{}

	var counter int
	var lock sync.Mutex
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go func(i int, w *sync.WaitGroup) {
			defer w.Done()
			lock.Lock()
			defer lock.Unlock()
			counter++
		}(i, &w)
	}

	w.Wait()
	t.Logf("Done, total: %d", counter)
}
