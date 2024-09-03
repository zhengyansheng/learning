package channel

import (
	"sync"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {

	ch := make(chan bool)

	timeout := time.Second * 3

	go func() {
		defer close(ch)
		time.Sleep(timeout + time.Second*10)
		ch <- true
		t.Log("exit goroutine")
	}()

	select {
	case v := <-ch:
		t.Logf("v: %v", v)
	case <-time.After(timeout):
		t.Log("timeout")
	}

}

func TestChannelPrint(t *testing.T) {
	var w sync.WaitGroup
	slice := [9]int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	ch := make([]chan struct{}, len(slice))
	for i := 0; i < len(slice); i++ {
		ch[i] = make(chan struct{})
	}

	for i := 0; i < len(slice); i++ {
		w.Add(1)
		go func(i int) {
			defer w.Done()

			if i > 0 {
				<-ch[i]
			}

			t.Log(slice[i])

			if i < len(slice)-1 {
				ch[i] <- struct{}{}
			}

		}(i)
	}

	ch[0] <- struct{}{}

	w.Wait()
	t.Log("Done")

}
