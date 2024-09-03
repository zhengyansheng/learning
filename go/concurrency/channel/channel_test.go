package channel

import (
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
