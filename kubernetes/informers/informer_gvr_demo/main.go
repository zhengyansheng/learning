package main

import (
	"flag"
	"fmt"
	"path/filepath"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/json"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
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

	// 1. 实例化 informer 工厂
	factory := informers.NewSharedInformerFactory(clientSet, defaultResync)

	// 2. 实例化 gvr
	gvr := schema.GroupVersionResource{Group: "apps", Version: "v1", Resource: "deployments"}

	//gvr := schema.GroupVersionResource{Group: "batch", Version: "v1", Resource: "cronjobs"}
	// 注册gvr
	genericInformer, err := factory.ForResource(gvr)
	if err != nil {
		panic(err)
	}

	_ = genericInformer.Informer()

	// 3. 启动 factory ( list and watch )
	stopCh := make(chan struct{})
	defer close(stopCh)
	factory.Start(stopCh)

	// 4. 等待所有的informer同步完成
	factory.WaitForCacheSync(stopCh)

	// 从 Indexer 中获取指定的资源
	genericList, err := genericInformer.Lister().List(labels.Everything())
	if err != nil {
		panic(err)
	}
	for i, object := range genericList {
		var deployment appsv1.Deployment

		marshal, err := json.Marshal(object)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = json.Unmarshal(marshal, &deployment)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(i, deployment.Namespace, deployment.Name)
	}

	// 5. 阻塞
	<-stopCh
}
