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
		t.Logf("test jitter until func, get v: %v", v)

	}, 5*time.Second, JitterFactor, true, ctx.Done())
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

func TestGroupStart(t *testing.T) {
	var w wait.Group
	for i := 0; i < 3; i++ {
		idx := i
		w.Start(func() {
			for {
				t.Logf("hello goroutine-%d", idx)
				<-time.After(time.Second)
			}
		})
	}
	t.Log("wait all goroutine finish")
	w.Wait()
	t.Log("Done")
}

func TestGroupStartWithContext(t *testing.T) {
	// 设置超时时间
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var w wait.Group
	for i := 0; i < 3; i++ {
		idx := i
		w.StartWithContext(ctx, func(context.Context) {
			for {
				// 业务逻辑
				t.Logf("hello goroutine-%d", idx)

				select {
				case <-ctx.Done():
					return
				default:
					<-time.After(time.Second)
				}
			}
		})
	}
	w.Wait()
	t.Log("Done")
}

func TestGroupStartWithChannel(t *testing.T) {
	var w wait.Group
	stopCh := make(chan struct{})

	// TODO:
	w.Start(func() {
		t.Log("main goroutine sleep 5s")
		<-time.After(time.Second * 5)
		stopCh <- struct{}{}
		close(stopCh)
	})

	for i := 0; i < 3; i++ {
		idx := i
		w.StartWithChannel(stopCh, func(<-chan struct{}) {
			for {
				// 业务逻辑
				t.Logf("hello goroutine-%d", idx)

				select {
				case <-stopCh:
					return
				default:
					<-time.After(time.Second)
				}
			}
		})
	}
	w.Wait()
	t.Log("Done")
}
