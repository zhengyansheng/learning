package wait_util

import (
	"context"
	"fmt"
	"testing"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/klog/v2"
)

const (
	JitterFactor = 1.2
)

func TestJitterUntil(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			time.Sleep(time.Second * 2)
		}
		close(ch)
	}()

	wait.JitterUntil(func() {
		v, ok := <-ch
		if !ok {
			cancel()
			return
		}
		klog.Infof("test jitter until func, get v: %v", v)

	}, 5*time.Second, JitterFactor, true, ctx.Done())
}

func TestPollImmediateUntil(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	err := wait.PollImmediateUntil(time.Second*1, func() (bool, error) {
		fmt.Println("continue")
		return true, nil
		// return true, fmt.Errorf("xxx")
	}, ctx.Done())
	if err != nil {
		panic(err)
	}
}
