package wait_util

import (
	"context"
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
