package main

import (
	"flag"
	"path/filepath"
	"reflect"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/klog/v2"
)

func main() {
	stopCh := make(chan struct{})
	defer close(stopCh)

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

	informerType := reflect.TypeOf(appsv1.Deployment{})
	klog.Infof("informerType: %+v\n", informerType)

	// 定义事件回调函数
	handlers := cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			switch obj.(type) {
			case *appsv1.Deployment:
				deploy := obj.(*appsv1.Deployment)
				klog.Infof("add deploy: %v/%v\n", deploy.Namespace, deploy.Name)
			}
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			switch oldObj.(type) {
			case *appsv1.Deployment:
				deploy := newObj.(*appsv1.Deployment)
				klog.Infof("update deploy: %v/%v\n", deploy.Namespace, deploy.Name)
			}
		},
		DeleteFunc: func(obj interface{}) {
			switch obj.(type) {
			case *appsv1.Deployment:
				deploy := obj.(*appsv1.Deployment)
				klog.Infof("delete deploy: %v/%v\n", deploy.Namespace, deploy.Name)
			}
		},
	}

	// 1. 创建 informer factory
	SharedInformerFactory := informers.NewSharedInformerFactory(clientSet, time.Minute*5)

	// 2. 创建 deploymentInformer
	deploymentSharedInformer := SharedInformerFactory.Apps().V1().Deployments().Informer()
	// 注册事件回调函数
	deploymentSharedInformer.AddEventHandler(handlers)

	//3. 启动 informers ( 已经注册到factory到informer )
	SharedInformerFactory.Start(stopCh) // controller.Run()
	//deploymentSharedInformer.Run(stopCh)

	// 4. 等待 informer list 操作执行完成
	SharedInformerFactory.WaitForCacheSync(stopCh)

	<-stopCh
}
