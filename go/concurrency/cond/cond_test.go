package cond

import (
	"math/rand/v2"
	"sync"
	"testing"
	"time"
)

func TestCond(t *testing.T) {

	// 创建cond
	c := sync.NewCond(&sync.Mutex{})

	var ready int

	for i := 0; i < 10; i++ {
		go func(i int) {

			time.Sleep(time.Duration(rand.Int64N(10)) * time.Second)

			// 加锁 更改等待条件ready
			c.L.Lock()
			ready++
			c.L.Unlock()

			t.Logf("运动员 %v 已准备就绪", i)

			time.Sleep(time.Second * 5)
			// 运动员准备就绪，通知裁判员
			c.Broadcast()
		}(i)
	}

	c.L.Lock()
	for ready != 10 { //检查条件是否满足
		c.Wait()
		t.Logf("裁判员被唤醒一次, %v", ready)
	}
	c.L.Unlock()

	t.Logf("所有运行员准备就绪，比赛即将开始......")
}
