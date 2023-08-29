package main

import (
	"context"
	"math/rand"
	"time"

	"github.com/zhengyansheng/multi_scheduler/scheduler"
	"k8s.io/klog/v2"
)

const schedulerName = "random-scheduler"

func main() {
	klog.Info("I'm a scheduler!")
	rand.Seed(time.Now().Unix())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 创建一个 Scheduler 对象
	sched := scheduler.NewScheduler(ctx, schedulerName, 0)

	// 启动 Scheduler
	sched.Run(ctx)

	// 阻塞主 goroutine
	select {}
}
