/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/google/uuid"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/leaderelection"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
	"k8s.io/klog/v2"
)

func buildConfig(kubeconfig string) (*rest.Config, error) {
	if kubeconfig != "" {
		cfg, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			return nil, err
		}
		return cfg, nil
	}

	cfg, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

/*
具体来说，通过调用k8s.io/client-go/tools/leaderelection库中的函数实现。主要过程是：
1. 首先获取kubernetes REST API的访问配置，在本代码中是通过buildConfig函数生成的rest.Config对象；
2. 然后创建一个LeaseLock对象作为锁，用于领导者选举过程中通过对租赁对象的读写来实现同步，具体来说就是通过创建一个指定名称的Lease对象（或ConfigMap、Endpoints对象）并对其进行写入操作来获得独占权限；
3. 最后使用指定的Callback执行选举循环，此处采用了RunOrDie方式，即无论选举是否成功都将阻塞程序。在选举循环中，当程序成为leader时，将调用OnStartedLeading()回调函数执行业务逻辑，当leader失去锁时，将调用OnStoppedLeading()回调函数退出程序，而当有新leader被选出时，则调用OnNewLeader()回调函数进行通知。
注意，在使用LeaseLock锁时也需要声明锁的Config参数中的ReleaseOnCancel，当程序退出时将释放锁资源。同时，需要确保在退出程序之前，已经把所有被租用的资源释放完毕，避免造成资源泄漏和竞争问题。
*/
func main() {
	klog.InitFlags(nil)

	var kubeconfig string
	var leaseLockName string
	var leaseLockNamespace string
	var id string

	// config 文件的绝对路径
	flag.StringVar(&kubeconfig, "kubeconfig", "", "absolute path to the kubeconfig file")
	// 持有者的身份标识
	flag.StringVar(&id, "id", uuid.New().String(), "the holder identity name")
	// 租约锁的名称
	flag.StringVar(&leaseLockName, "lease-lock-name", "", "the lease lock resource name")
	// 租约锁的命名空间
	flag.StringVar(&leaseLockNamespace, "lease-lock-namespace", "", "the lease lock resource namespace")
	flag.Parse()

	// 若没有指定lease-lock-name和lease-lock-namespace，则无法获取lease lock资源，程序退出
	if leaseLockName == "" {
		klog.Fatal("unable to get lease lock resource name (missing lease-lock-name flag).")
	}
	if leaseLockNamespace == "" {
		klog.Fatal("unable to get lease lock resource namespace (missing lease-lock-namespace flag).")
	}

	// leader election uses the Kubernetes API by writing to a
	// lock object, which can be a LeaseLock object (preferred),
	// a ConfigMap, or an Endpoints (deprecated) object.
	// Conflicting writes are detected and each client handles those actions
	// independently.
	// 通过rest.Config来生成clientset
	config, err := buildConfig(kubeconfig)
	if err != nil {
		klog.Fatal(err)
	}
	client := clientset.NewForConfigOrDie(config)

	// 定义运行主逻辑的函数run
	run := func(ctx context.Context) {
		// complete your controller loop here
		klog.Info("Controller loop...")

		select {}
	}

	// use a Go context so we can tell the leaderelection code when we
	// want to step down
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 监听中断和Linux SIGTERM信号，并将它们发送到channel ch；当接收到信号时，取消context
	// listen for interrupts or the Linux SIGTERM signal and cancel
	// our context, which the leader election code will observe and
	// step down
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-ch
		klog.Info("Received termination, signaling shutdown")
		cancel()
	}()

	// we use the Lease lock type since edits to Leases are less common
	// and fewer objects in the cluster watch "all Leases".
	// 使用 resource lock 锁类型
	lock := &resourcelock.LeaseLock{
		// 指定 锁的名称 和 锁的命名空间
		LeaseMeta: metav1.ObjectMeta{
			Name:      leaseLockName,
			Namespace: leaseLockNamespace,
		},
		// reset client
		Client: client.CoordinationV1(),
		// 锁的配置
		LockConfig: resourcelock.ResourceLockConfig{
			Identity: id,
		},
	}

	// start the leader election code loop
	leaderelection.RunOrDie(ctx, leaderelection.LeaderElectionConfig{
		Lock: lock, // 使用Lease lock类型
		// IMPORTANT: you MUST ensure that any code you have that
		// is protected by the lease must terminate **before**
		// you call cancel. Otherwise, you could have a background
		// loop still running and another process could
		// get elected before your background loop finished, violating
		// the stated goal of the lease.
		ReleaseOnCancel: true, // 当context被取消时释放锁
		LeaseDuration:   60 * time.Second,
		RenewDeadline:   15 * time.Second,
		RetryPeriod:     5 * time.Second,
		Callbacks: leaderelection.LeaderCallbacks{
			OnStartedLeading: func(ctx context.Context) {
				// we're notified when we start - this is where you would
				// usually put your code
				// 当成为leader时，会调用这个回调函数，可以在这里执行controller逻辑
				run(ctx)
			},
			OnStoppedLeading: func() {
				// 当不再是leader时，会调用这个回调函数，可以在这里执行清理逻辑
				// we can do cleanup here
				klog.Infof("leader lost: %s", id)
				os.Exit(0)
			},
			OnNewLeader: func(identity string) {
				// 当有新的leader选举成功时，会调用这个回调函数
				// we're notified when new leader elected
				if identity == id {
					// 如果获取到锁的identity与本身的identity一致，表示本身成为领导者
					// I just got the lock
					return
				}
				klog.Infof("new leader elected: %s", identity)
			},
		},
	})
}
