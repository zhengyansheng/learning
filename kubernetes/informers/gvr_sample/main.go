package main

import (
	"flag"
	"path/filepath"

	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/klog/v2"
)

func getClientSet() *kubernetes.Clientset {
	// client
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
	return clientSet
}

func main() {

	// 1. 创建clientSet
	clientSet := getClientSet()

	// 2. 创建informer factory
	sharedInformers := informers.NewSharedInformerFactory(clientSet, 0)

	// 3. 定义gvr
	gvr := schema.GroupVersionResource{Group: "apps", Version: "v1", Resource: "deployments"}

	// 4. 根据gvr创建informer
	genericInformer, err := sharedInformers.ForResource(gvr)
	if err != nil {
		return
	}

	// 5. 创建gvr对应的informer 和 注册事件处理函数
	genericInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			accessor, err := meta.Accessor(obj)
			if err != nil {
				return
			}
			klog.Infof("add %v/%v", accessor.GetNamespace(), accessor.GetName())
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			accessor, err := meta.Accessor(newObj)
			if err != nil {
				return
			}
			klog.Infof("update %v/%v", accessor.GetNamespace(), accessor.GetName())
		},
		DeleteFunc: func(obj interface{}) {
			accessor, err := meta.Accessor(obj)
			if err != nil {
				return
			}
			klog.Infof("delete %v/%v", accessor.GetNamespace(), accessor.GetName())
		},
	})

	// 6. 启动informer
	stopCh := make(chan struct{})
	sharedInformers.Start(stopCh)
	<-stopCh

}
