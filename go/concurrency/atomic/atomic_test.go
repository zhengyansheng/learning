package atomic

import (
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomicAdd(t *testing.T) {
	var cnt int64 = 0
	var wg sync.WaitGroup

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt64(&cnt, 1)
			wg.Done()
		}()
	}

	wg.Wait()
	t.Logf("Final cnt: %v", cnt)
}

func TestAtomicSwap(t *testing.T) {

	var x int64 = 100
	t.Log(x)
	atomic.SwapInt64(&x, 200)
	t.Log(x)

	var y int64 = 1000
	v := atomic.LoadInt64(&y)
	t.Log(v)
}

func TestAtomicCompareAndSwapInt64(t *testing.T) {

	var wg sync.WaitGroup
	var counter int64

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int, w *sync.WaitGroup) {
			defer wg.Done()

			for {
				v := atomic.LoadInt64(&counter)
				val := v + 1
				// 无锁更新: 使用 atomic.CompareAndSwapInt64 尝试将 counter 的值从 old 更新为 new
				if atomic.CompareAndSwapInt64(&counter, v, val) {
					break
				}
			}

		}(i, &wg)

	}
	wg.Wait()
	t.Logf("Done: %v", counter)
}
