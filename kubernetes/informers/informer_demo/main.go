package main

import (
	"flag"
	"fmt"
	"path/filepath"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var (
	defaultResync = time.Minute * 5
)

func main() {
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

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	// 1. 实例化 informer factory
	factory := informers.NewSharedInformerFactory(clientSet, defaultResync)

	// 2. 向 factory 注册 各种Informer, 比如： podInformer
	podInformer := factory.Core().V1().Pods()

	sharedIndexInformer := podInformer.Informer()
	// 添加事件回调函数 （ADD,UPDATE,DELETE）
	sharedIndexInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			switch obj.(type) {
			case *apiv1.Pod:
				pod := obj.(*apiv1.Pod)
				fmt.Printf("add pod: %v/%v\n", pod.Namespace, pod.Name)
			case *appsv1.Deployment:
				deploy := obj.(*appsv1.Deployment)
				fmt.Printf("add deploy: %v/%v\n", deploy.Namespace, deploy.Name)
			}
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			switch oldObj.(type) {
			case *apiv1.Pod:
				pod := newObj.(*apiv1.Pod)
				fmt.Printf("update pod: %v/%v\n", pod.Namespace, pod.Name)
			case *appsv1.Deployment:
				deploy := newObj.(*appsv1.Deployment)
				fmt.Printf("update deploy: %v/%v\n", deploy.Namespace, deploy.Name)
			}
		},
		DeleteFunc: func(obj interface{}) {
			switch obj.(type) {
			case *apiv1.Pod:
				pod := obj.(*apiv1.Pod)
				fmt.Printf("delete pod: %v/%v\n", pod.Namespace, pod.Name)
			case *appsv1.Deployment:
				deploy := obj.(*appsv1.Deployment)
				fmt.Printf("delete deploy: %v/%v\n", deploy.Namespace, deploy.Name)
			}
		},
	})

	deploymentInformer := factory.Apps().V1().Deployments()
	sharedIndexInformer2 := deploymentInformer.Informer()
	// 注册回调函数
	sharedIndexInformer2.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			switch obj.(type) {
			case *apiv1.Pod:
				pod := obj.(*apiv1.Pod)
				fmt.Printf("add pod: %v/%v\n", pod.Namespace, pod.Name)
			case *appsv1.Deployment:
				deploy := obj.(*appsv1.Deployment)
				fmt.Printf("add deploy: %v/%v\n", deploy.Namespace, deploy.Name)
			}
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			switch oldObj.(type) {
			case *apiv1.Pod:
				pod := newObj.(*apiv1.Pod)
				fmt.Printf("update pod: %v/%v\n", pod.Namespace, pod.Name)
			case *appsv1.Deployment:
				deploy := newObj.(*appsv1.Deployment)
				fmt.Printf("update deploy: %v/%v\n", deploy.Namespace, deploy.Name)
			}
		},
		DeleteFunc: func(obj interface{}) {
			switch obj.(type) {
			case *apiv1.Pod:
				pod := obj.(*apiv1.Pod)
				fmt.Printf("delete pod: %v/%v\n", pod.Namespace, pod.Name)
			case *appsv1.Deployment:
				deploy := obj.(*appsv1.Deployment)
				fmt.Printf("delete deploy: %v/%v\n", deploy.Namespace, deploy.Name)
			}
		},
	})

	// 3. 启动 所有注册到factory到informers ( list and watch )
	stopCh := make(chan struct{})
	defer close(stopCh)
	factory.Start(stopCh)

	// 4. 等待所有的informer同步完成
	factory.WaitForCacheSync(stopCh)

	fmt.Println("====================从Indexer缓存中读取数据=============================")
	// pod index
	podLister := podInformer.Lister()
	pods, err := podLister.Pods(apiv1.NamespaceDefault).List(labels.Everything())
	if err != nil {
		return
	}
	for i, pod := range pods {
		fmt.Printf("[i: %v -> pod], name: %v, namespace: %v\n", i, pod.Name, pod.Namespace)
	}

	// deployment
	deploymentLister := deploymentInformer.Lister()
	deployments, err := deploymentLister.Deployments(apiv1.NamespaceDefault).List(labels.Everything())
	if err != nil {
		return
	}
	if err != nil {
		return
	}
	for i, deployment := range deployments {
		fmt.Printf("[i: %v -> deployment], name: %v, namespace: %v\n", i, deployment.Name, deployment.Namespace)
	}
	select {}
}

func simple() {
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

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	handlers := cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			switch obj.(type) {
			case *apiv1.Pod:
				pod := obj.(*apiv1.Pod)
				fmt.Printf("add pod: %v/%v\n", pod.Namespace, pod.Name)
			case *appsv1.Deployment:
				deploy := obj.(*appsv1.Deployment)
				fmt.Printf("add deploy: %v/%v\n", deploy.Namespace, deploy.Name)
			}
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			switch oldObj.(type) {
			case *apiv1.Pod:
				pod := newObj.(*apiv1.Pod)
				fmt.Printf("update pod: %v/%v\n", pod.Namespace, pod.Name)
			case *appsv1.Deployment:
				deploy := newObj.(*appsv1.Deployment)
				fmt.Printf("update deploy: %v/%v\n", deploy.Namespace, deploy.Name)
			}
		},
		DeleteFunc: func(obj interface{}) {
			switch obj.(type) {
			case *apiv1.Pod:
				pod := obj.(*apiv1.Pod)
				fmt.Printf("delete pod: %v/%v\n", pod.Namespace, pod.Name)
			case *appsv1.Deployment:
				deploy := obj.(*appsv1.Deployment)
				fmt.Printf("delete deploy: %v/%v\n", deploy.Namespace, deploy.Name)
			}
		},
	}
	// 1. 实例化 informer factory
	factory := informers.NewSharedInformerFactory(clientSet, defaultResync)

	// 2. 向 factory 注册 各种Informer, 比如： podInformer
	podInformer := factory.Core().V1().Pods()
	sharedIndexInformer := podInformer.Informer()
	sharedIndexInformer.AddEventHandler(handlers)

	// 3. 启动 所有注册到factory到informers ( list and watch )
	stopCh := make(chan struct{})
	defer close(stopCh)
	factory.Start(stopCh)

	// 4. 等待所有的informer同步完成
	factory.WaitForCacheSync(stopCh)

	fmt.Println("====================从Indexer缓存中读取数据=============================")
	// pod index
	podLister := podInformer.Lister()
	pods, err := podLister.Pods(apiv1.NamespaceDefault).List(labels.Everything())
	if err != nil {
		return
	}
	for i, pod := range pods {
		fmt.Printf("[i: %v -> pod], name: %v, namespace: %v\n", i, pod.Name, pod.Namespace)
	}
	select {}
}
