package wait_util

import (
	"context"
	"testing"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
)

// TestPollImmediateUntil 立刻轮训执行，直到函数返回true, an error or stopCh is closed
func TestPollImmediateUntil(t *testing.T) {
	// 设置超时时间
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	err := wait.PollImmediateUntil(time.Second*1, func() (bool, error) {

		t.Log("hello")

		//return true, nil // 正常，结束
		return false, nil // 继续
		//return true, fmt.Errorf("xxx") // 错误，函数执行结束
	}, ctx.Done())
	if err != nil {
		t.Fatalf("poll immediate until err: %v", err)
	}
	t.Log("Done")
}

// TestPollImmediateWithContext  立刻轮训执行，直到函数返回true, error or the specified context is cancelled or expired
func TestPollImmediateWithContext(t *testing.T) {
	// 设置超时时间
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	var i int
	err := wait.PollImmediateWithContext(ctx, time.Second*1, time.Second*15, func(context.Context) (done bool, err error) {
		i++
		t.Logf("hello %v", i)
		//return true, nil // 正常，结束
		return false, nil // 继续
		//return true, fmt.Errorf("xxx") // 错误，函数执行结束
	})
	if err != nil {
		t.Fatalf("poll immediate context err: %v", err)
	}
}
