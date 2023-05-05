package wait_util

import (
	"context"
	"testing"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
)

const (
	JitterFactor = 1.2
)

func TestForever(t *testing.T) {
	wait.Forever(func() {
		t.Log("hello world")
	}, time.Second*2)
}

func TestUntil(t *testing.T) {
	stopCh := make(chan struct{})

	go func() {
		defer close(stopCh)
		<-time.After(time.Second * 5)
	}()

	var i int
	wait.Until(func() {
		t.Logf("hello world, %v", i)
		i++
	}, time.Second*1, stopCh)
}

func TestJitterUntil(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := make(chan struct{})
	go func() {
		defer close(ch)
		for i := 0; i < 5; i++ {
			ch <- struct{}{}
			<-time.After(time.Second)
		}
	}()

	wait.JitterUntil(func() {
		_, ok := <-ch
		if !ok {
			cancel()
			return
		}
		t.Logf("call jitter until")

	}, 1*time.Second, JitterFactor, true, ctx.Done())
}

func TestPollImmediateUntil(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	err := wait.PollImmediateUntil(time.Second*1, func() (bool, error) {
		t.Log("continue")
		return true, nil
		// return true, fmt.Errorf("xxx")
	}, ctx.Done())
	if err != nil {
		panic(err)
	}
}
