package context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// 定义一个键的类型，避免与其他包中的键冲突
type key string

const requestIDKey = key("requestID")

func TestContext(t *testing.T) {
	t.Run("WithTimeout", func(t *testing.T) {
		// 创建一个超时context
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		// 创建一个通道来接收worker的结果
		done := make(chan error)

		go WorkerFuncTimeout(ctx, done)

		t.Logf("done: %v", <-done)
	})

	t.Run("WithCancel", func(t *testing.T) {
		// 创建一个超时context
		ctx, cancel := context.WithCancel(context.Background())

		// 启动多个worker goroutine
		for i := 0; i < 10; i++ {
			go WorkerFuncCancel(ctx, i)
		}

		// 让worker运行一段时间
		time.Sleep(2 * time.Second)

		// 取消所有worker
		fmt.Println("Cancelling all workers...")
		cancel()
		// 等待一段时间以观察所有worker被取消
		time.Sleep(1 * time.Second)

	})

	t.Run("WithValue", func(t *testing.T) {
		// 创建一个超时context
		ctx := context.WithValue(context.Background(), requestIDKey, "12345")

		// 启动一个worker goroutine
		go WorkerFuncValue(ctx)

		// 等待worker完成
		time.Sleep(1 * time.Second)
		t.Log("Done")
	})

	t.Run("WithDeadline", func(t *testing.T) {
		ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*2))
		// 超时
		<-ctx.Done()
		t.Logf("err %v", ctx.Err()) // context deadline exceeded
		cancel()

	})

	t.Run("WithDeadline2", func(t *testing.T) {
		ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*5))
		// 主动撤销
		cancel()

		<-ctx.Done()
		t.Logf("err %v", ctx.Err()) // context canceled

	})
}

func WorkerFuncTimeout(ctx context.Context, done chan error) {
	taskCh := make(chan struct{})
	go func() {
		time.Sleep(time.Second * 2)
		taskCh <- struct{}{}
	}()

	for {
		select {
		case <-ctx.Done():
			done <- ctx.Err()
			return
		case <-taskCh:
			done <- nil
			return
		}
	}
}

func WorkerFuncCancel(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			// 当上下文被取消时，终止工作
			fmt.Printf("Worker %d cancelled: %v\n", id, ctx.Err())
			return
		default:
			fmt.Printf("Worker %d is working\n", id)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func WorkerFuncValue(ctx context.Context) {
	if reqValue, ok := ctx.Value(requestIDKey).(string); ok {
		fmt.Printf("Worker started with request ID: %s\n", reqValue)
	} else {
		fmt.Println("Worker started without request ID")
	}

	// 模拟长时间运行的任务
	time.Sleep(2 * time.Second)
	fmt.Println("Worker finished")
}
