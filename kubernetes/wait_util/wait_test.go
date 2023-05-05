package wait_util

import (
	"context"
	"testing"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
)

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
