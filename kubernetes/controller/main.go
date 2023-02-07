package main

import (
	"flag"
	"fmt"
	"path/filepath"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/client-go/util/workqueue"
)

func getClientSet() (*kubernetes.Clientset, error) {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}

	return kubernetes.NewForConfig(config)
}

type Controller struct {
	Index    cache.Indexer
	queue    workqueue.RateLimitingInterface
	Informer cache.Controller
}

func NewController(index cache.Indexer, queue workqueue.RateLimitingInterface, informer cache.Controller) *Controller {
	return &Controller{Index: index, queue: queue, Informer: informer}
}

func (c *Controller) Run(stopCh chan struct{}) {
	// 处理 panic
	defer runtime.HandleCrash()

	// 关闭 queue
	defer c.queue.ShutDown()

	fmt.Println("Start controller")

	// 启动 informer
	go c.Informer.Run(stopCh)

	// 等待informer lw同步完成
	if !cache.WaitForCacheSync(stopCh, c.Informer.HasSynced) {
		fmt.Println("wait for cache sync failed")
		return
	}

	go wait.Until(func() {
		c.runWorker()
	}, time.Second, stopCh)

	<-stopCh

	fmt.Println("Stop controller")
}

func (c *Controller) runWorker() {
	for c.processItem() {
	}
}

func (c *Controller) processItem() bool {
	// 从 queue[0] 拿到 key
	key, shutdown := c.queue.Get()
	if shutdown {
		return false
	}
	// Done 表示这个key已经处理完成
	defer c.queue.Done(key)

	if err := c.handleObject(key.(string)); err != nil {
		// 同一个key入队次数大于5次就丢弃
		if c.queue.NumRequeues(key) < 5 {
			c.queue.Add(key)
		}
	}
	return true
}

func (c *Controller) handleObject(key string) error {
	object, exists, err := c.Index.GetByKey(key)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("key: %v does not exists in index", key)
	}
	if v, ok := object.(*corev1.Pod); ok {
		fmt.Printf("process: name: %v, name: %v\n", v.Namespace, v.Name)
	}

	return nil

}

func main() {
	clientSet, err := getClientSet()
	if err != nil {
		panic(err)
	}

	// workqueue
	queue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())

	// informer
	lw := cache.NewListWatchFromClient(clientSet.CoreV1().RESTClient(), "pods", "default", fields.Everything())
	indexer, informer := cache.NewIndexerInformer(lw, &corev1.Pod{}, time.Minute*5, cache.ResourceEventHandlerFuncs{
		// 资源事件处理函数
		AddFunc: func(obj interface{}) {

			// obj 是一个对象, 计算得到key
			key, err := cache.MetaNamespaceKeyFunc(obj)
			if err != nil {
				return
			}
			queue.Add(key)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			// obj 是一个对象, 计算得到key
			key, err := cache.MetaNamespaceKeyFunc(newObj)
			if err != nil {
				return
			}
			queue.Add(key)
		},
		DeleteFunc: func(obj interface{}) {
			// obj 是一个对象, 计算得到key
			key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
			if err != nil {
				return
			}
			queue.Add(key)
		},
	}, cache.Indexers{})

	// logic
	c := NewController(indexer, queue, informer)

	stopCh := make(chan struct{})
	defer close(stopCh)
	c.Run(stopCh)

	<-stopCh

}
