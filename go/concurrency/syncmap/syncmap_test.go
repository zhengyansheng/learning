package syncmap

import (
	"sync"
	"testing"
)

func TestSyncMap(t *testing.T) {
	m := sync.Map{}
	var w sync.WaitGroup

	for i := 0; i < 1000; i++ {
		w.Add(1)
		go func(i int) {
			defer w.Done()
			m.Store(i, i*10)
		}(i)
	}

	for i := 0; i < 1000; i++ {
		w.Add(1)
		go func(i int) {
			defer w.Done()
			if v, ok := m.Load(i); ok {
				t.Logf("m[%d] = %d\n", i, v)
			}
		}(i)
	}

	w.Wait()
	t.Log("Done")
}

func TestMapRWLock(t *testing.T) {
	m := make(map[int]int, 10)
	var w sync.WaitGroup
	var l sync.RWMutex

	for i := 0; i < 1000; i++ {
		w.Add(1)
		go func(i int) {
			defer w.Done()
			l.Lock()
			m[i] = i * 10
			l.Unlock()
		}(i)
	}

	for i := 0; i < 1000; i++ {
		w.Add(1)
		go func(i int) {
			defer w.Done()
			l.RLock()
			t.Logf("m[%d] = %d\n", i, m[i])
			l.RUnlock()
		}(i)
	}

	w.Wait()
	t.Log("Done")
}

func TestMap(t *testing.T) {
	m := make(map[int]int, 10)
	var w sync.WaitGroup

	for i := 0; i < 1000; i++ {
		w.Add(1)
		go func(i int) {
			defer w.Done()
			m[i] = i * 10
		}(i)
	}

	for i := 0; i < 1000; i++ {
		w.Add(1)
		go func(i int) {
			defer w.Done()
			t.Logf("m[%d] = %d\n", i, m[i])
		}(i)
	}

	w.Wait()
	t.Log("Done")
}
