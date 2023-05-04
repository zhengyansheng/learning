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

func TestGroupStart(t *testing.T) {
	var w wait.Group
	for i := 0; i < 3; i++ {
		v := i
		w.Start(func() {
			for {
				fmt.Printf("hello %d\n", v)
				time.Sleep(time.Second)
			}
		})
	}
	w.Wait()
	fmt.Println("done")
}

func TestGroupStartWithContext(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
	defer cancel()

	var w wait.Group
	for i := 0; i < 3; i++ {
		v := i
		w.StartWithContext(ctx, func(context.Context) {
			for {
				// 业务逻辑
				fmt.Printf("hello %d\n", v)

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
	fmt.Println("done")
}

func TestGroupStartWithChannel(t *testing.T) {
	stopCh := make(chan struct{})

	go func() {
		time.Sleep(time.Second * 5)
		stopCh <- struct{}{}
		close(stopCh)
	}()

	var w wait.Group
	for i := 0; i < 3; i++ {
		v := i
		w.StartWithChannel(stopCh, func(<-chan struct{}) {
			for {
				// 业务逻辑
				fmt.Printf("hello %d\n", v)

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
	fmt.Println("done")
}
