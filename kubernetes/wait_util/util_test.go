package wait_util

import (
	"context"
	"testing"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/utils/clock"
)

const (
	JitterFactor = 1.2
)

// TestForever 永远不会退出
func TestForever(t *testing.T) {
	wait.Forever(func() {
		t.Log("hello world")
	}, time.Second*2)
}

// TestUntil 直到stopCh关闭函数才会结束
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

// TestJitterUntil 直到stopCh关闭函数才会结束并附加间隔的抖动因子
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

// TestBackoffUntil 退避stopCh关闭函数才会结束并附加间隔的抖动因子和滑动窗口
func TestBackoffUntil(t *testing.T) {
	stopCh := make(chan struct{})
	go func() {
		<-time.After(time.Second * 3)
		defer close(stopCh)
	}()
	wait.BackoffUntil(func() {
		t.Log(time.Now().Format("2006-01-02 15:04:05"), "call backoff until")
	}, wait.NewJitteredBackoffManager(1*time.Second, 1.1, &clock.RealClock{}), true, stopCh)
}

// TestBackoffUntilWithExponential 退避stopCh关闭函数才会结束并附加间隔的抖动因子和滑动窗口，指数退避	1s, 2s, 4s, 8s, 10s, 1s, 2s, 4s
func TestBackoffUntilWithExponential(t *testing.T) {
	stopCh := make(chan struct{})
	defer close(stopCh)
	wait.BackoffUntil(func() {
		t.Log(time.Now().Format("2006-01-02 15:04:05"), "call backoff until")
	}, wait.NewExponentialBackoffManager(1*time.Second, 10*time.Second, 10*time.Second, 2.0, 0.0, &clock.RealClock{}), true, stopCh)
}
