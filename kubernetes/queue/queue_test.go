package queue

import (
	"testing"
	"time"

	"k8s.io/client-go/util/workqueue"
)

func TestWorkQueueRateLimiter2(t *testing.T) {
	queue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())

	queue.Add("one")
	queue.Add("two")

	// 非阻塞
	t.Log(queue.Get())
	t.Log(queue.Get())

	go func() {
		for i := 1; i < 102; i++ {
			// 限速添加
			queue.AddRateLimited(i)
		}
		t.Log("add rate limited done")
	}()
	for {
		item, shutdown := queue.Get()
		if shutdown {
			return
		}
		t.Log(time.Now(), item)
	}

}
