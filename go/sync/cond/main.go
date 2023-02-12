package main

import (
	"fmt"
	"sync"
	"time"
)

type Fifo struct {
	Queue []string
	Cond  *sync.Cond
}

func NewFifo() *Fifo {
	return &Fifo{
		Queue: []string{},
		Cond:  sync.NewCond(&sync.Mutex{}),
	}
}

func (f *Fifo) Enqueue(item string) {
	f.Cond.L.Lock()
	defer f.Cond.L.Unlock()
	f.Queue = append(f.Queue, item)
	f.Cond.Broadcast() // 唤醒

}

func (f *Fifo) Dequeue() string {
	f.Cond.L.Lock()
	defer f.Cond.L.Unlock()
	if f.Size() == 0 {
		f.Cond.Wait() // 堵塞
	}
	var item string
	item, f.Queue = f.Queue[0], f.Queue[1:]
	return item
}

func (f *Fifo) Size() int {
	return len(f.Queue)
}

func main() {
	f := NewFifo()

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for {
			f.Enqueue("a")
			time.Sleep(time.Second)
		}
	}()

	go func() {
		defer wg.Done()
		for {
			item := f.Dequeue()
			fmt.Printf("pop item: %v\n", item)
			time.Sleep(time.Second * 3)
		}
	}()

	wg.Wait()
}
