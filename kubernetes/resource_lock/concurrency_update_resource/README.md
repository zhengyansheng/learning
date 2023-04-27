# 乐观锁

## 1. 模拟 并发修改同一个资源时会存在锁的问题

```go
	w := sync.WaitGroup{}
	for i := 0; i < ConcurrentQuantity; i++ {
		w.Add(1)
		go func(clientSetA *kubernetes.Clientset, index int) {
			defer w.Done()

			// 方式1
			// 并发修改资源 会存在只有一个修改成功，其它都失败
			err := updateByGetAndUpdate(clientSetA, DpName, index)

			//// 方式2：
			//// 冲突时重试
			//var retryParam = wait.Backoff{
			//	Steps:    20,
			//	Duration: 10 * time.Millisecond,
			//	Factor:   1.0,
			//	Jitter:   0.1,
			//}

			err := retry.RetryOnConflict(retryParam, func() error {
				return updateByGetAndUpdate(clientSet, DpName, index)
			})

			if err != nil {
				klog.Infof("goroutine-%d update err: %v", index, err)
				return
			}
			klog.Infof("goroutine-%d update success", index)

		}(clientSet, i)
	}

	w.Wait()
```

```go
➜  concurrency_update_resource git:(master) ✗ go run main.go
I0427 22:58:13.366280   44704 operator.go:27] 开始创建deployment
I0427 22:58:13.366400   44704 operator.go:124] Creating deployment...
I0427 22:58:13.502635   44704 operator.go:130] Created deployment "nginx-deployment"
I0427 22:58:13.502671   44704 operator.go:36] 在协程中并发更新自定义标签
I0427 22:58:13.537382   44704 operator.go:180] goroutine-4, current label: 101, resource version: 8869535
I0427 22:58:13.537438   44704 operator.go:180] goroutine-3, current label: 101, resource version: 8869535
I0427 22:58:13.549107   44704 operator.go:180] goroutine-2, current label: 101, resource version: 8869535
I0427 22:58:13.549143   44704 operator.go:180] goroutine-1, current label: 101, resource version: 8869535
I0427 22:58:13.853234   44704 operator.go:180] goroutine-0, current label: 101, resource version: 8869535
I0427 22:58:13.909604   44704 operator.go:60] goroutine-2 update err: Operation cannot be fulfilled on deployments.apps "nginx-deployment": the object has been modified; please apply your changes to the latest version and try again
I0427 22:58:13.909598   44704 operator.go:60] goroutine-1 update err: Operation cannot be fulfilled on deployments.apps "nginx-deployment": the object has been modified; please apply your changes to the latest version and try again
I0427 22:58:13.909673   44704 operator.go:187] ---> [success] goroutine-3, current label: 101, resource version old: 8869535, new: 8869537
I0427 22:58:13.909608   44704 operator.go:60] goroutine-4 update err: Operation cannot be fulfilled on deployments.apps "nginx-deployment": the object has been modified; please apply your changes to the latest version and try again
I0427 22:58:13.909709   44704 operator.go:63] goroutine-3 update success
I0427 22:58:13.909984   44704 operator.go:60] goroutine-0 update err: Operation cannot be fulfilled on deployments.apps "nginx-deployment": the object has been modified; please apply your changes to the latest version and try again
I0427 22:58:13.944094   44704 operator.go:77] 自定义标签的最终值为: 101，耗时442毫秒

```

## 2. 模拟 通过冲突重试 

```go
	w := sync.WaitGroup{}
	for i := 0; i < ConcurrentQuantity; i++ {
		w.Add(1)
		go func(clientSetA *kubernetes.Clientset, index int) {
			defer w.Done()

			// 方式1
			// 并发修改资源 会存在只有一个修改成功，其它都失败
			//err := updateByGetAndUpdate(clientSetA, DpName, index)

			// 方式2：
			// 冲突时重试
			var retryParam = wait.Backoff{
				Steps:    10,
				Duration: 10 * time.Millisecond,
				Factor:   1.0,
				Jitter:   0.1,
			}

			err := retry.RetryOnConflict(retryParam, func() error {
				return updateByGetAndUpdate(clientSet, DpName, index)
			})

			if err != nil {
				klog.Infof("goroutine-%d update err: %v", index, err)
				return
			}
			klog.Infof("goroutine-%d update success", index)

		}(clientSet, i)
	}

	w.Wait()
```

```go
➜  concurrency_update_resource git:(master) ✗ go run main.go
I0427 23:00:01.905045   44962 operator.go:28] 开始创建deployment
I0427 23:00:01.905195   44962 operator.go:127] Creating deployment...
I0427 23:00:01.982857   44962 operator.go:133] Created deployment "nginx-deployment"
I0427 23:00:01.982900   44962 operator.go:37] 在协程中并发更新自定义标签
I0427 23:00:02.045604   44962 operator.go:183] goroutine-1, current label: 101, resource version: 8869764
I0427 23:00:02.058775   44962 operator.go:183] goroutine-0, current label: 101, resource version: 8869764
I0427 23:00:02.080647   44962 operator.go:183] goroutine-4, current label: 101, resource version: 8869764
I0427 23:00:02.090303   44962 operator.go:183] goroutine-2, current label: 101, resource version: 8869764
I0427 23:00:02.090320   44962 operator.go:183] goroutine-3, current label: 101, resource version: 8869764
I0427 23:00:02.120550   44962 operator.go:190] ---> [success] goroutine-1, current label: 101, resource version old: 8869764, new: 8869765
I0427 23:00:02.120593   44962 operator.go:66] goroutine-1 update success
I0427 23:00:02.327474   44962 operator.go:183] goroutine-0, current label: 102, resource version: 8869765
I0427 23:00:02.525980   44962 operator.go:183] goroutine-3, current label: 102, resource version: 8869765
I0427 23:00:02.945693   44962 operator.go:183] goroutine-2, current label: 102, resource version: 8869775
I0427 23:00:02.957103   44962 operator.go:183] goroutine-4, current label: 102, resource version: 8869775
I0427 23:00:03.543688   44962 operator.go:190] ---> [success] goroutine-2, current label: 102, resource version old: 8869775, new: 8869778
I0427 23:00:03.543738   44962 operator.go:66] goroutine-2 update success
I0427 23:00:03.923817   44962 operator.go:183] goroutine-0, current label: 103, resource version: 8869778
I0427 23:00:04.137089   44962 operator.go:183] goroutine-3, current label: 103, resource version: 8869778
I0427 23:00:04.324088   44962 operator.go:183] goroutine-4, current label: 103, resource version: 8869778
I0427 23:00:04.572844   44962 operator.go:190] ---> [success] goroutine-0, current label: 103, resource version old: 8869778, new: 8869782
I0427 23:00:04.572866   44962 operator.go:66] goroutine-0 update success
I0427 23:00:05.308574   44962 operator.go:183] goroutine-3, current label: 104, resource version: 8869782
I0427 23:00:05.323007   44962 operator.go:183] goroutine-4, current label: 104, resource version: 8869782
I0427 23:00:05.848317   44962 operator.go:190] ---> [success] goroutine-3, current label: 104, resource version old: 8869782, new: 8869786
I0427 23:00:05.848366   44962 operator.go:66] goroutine-3 update success
I0427 23:00:06.205742   44962 operator.go:183] goroutine-4, current label: 105, resource version: 8869786
I0427 23:00:07.043547   44962 operator.go:190] ---> [success] goroutine-4, current label: 105, resource version old: 8869786, new: 8869789
I0427 23:00:07.043598   44962 operator.go:66] goroutine-4 update success
I0427 23:00:07.390286   44962 operator.go:80] 自定义标签的最终值为: 105，耗时5408毫秒
```

## 参考
- [文章来源](https://blog.csdn.net/boling_cavalry/article/details/128745382)
- [代码来源](https://github.com/zq2599/blog_demos/tree/master/tutorials/client-go-tutorials)