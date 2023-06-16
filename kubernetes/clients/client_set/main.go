package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"
	"time"

	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/klog/v2"
	deploymentutil "k8s.io/kubectl/pkg/util/deployment"
)

var (
	defaultName = "nginx-deployment2"
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

func get(clientSet *kubernetes.Clientset) {
	// get
	deploy, err := clientSet.AppsV1().Deployments(apiv1.NamespaceDefault).Get(context.TODO(), defaultName, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		fmt.Printf("%v not found\n", defaultName)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Println(deploy.Name, deploy.Namespace)
}

func list(clientSet *kubernetes.Clientset) {
	// list
	namespace := apiv1.NamespaceDefault
	namespace = "sys"
	opts := metav1.ListOptions{}
	deployList, err := clientSet.AppsV1().Deployments(namespace).List(context.TODO(), opts)
	if errors.IsNotFound(err) {
		fmt.Printf("list err: %v\n", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Println(len(deployList.Items), deployList.ResourceVersion)
}

func selectorDeepEqual(clientSet *kubernetes.Clientset) {
	deployment, err := clientSet.AppsV1().Deployments(apiv1.NamespaceDefault).Get(context.TODO(), "my-dep", metav1.GetOptions{})
	if err != nil {
		return
	}
	fmt.Println(deployment.Spec.Selector)

	everything := metav1.LabelSelector{}
	fmt.Println(everything)

	//deployment.Spec.Replicas = int32Ptr(0)
	//update, err := clientSet.AppsV1().Deployments(apiv1.NamespaceDefault).Update(context.TODO(), deployment, metav1.UpdateOptions{})
	//if err != nil {
	//	return
	//}
	//fmt.Println(update)
	//return

	//d := deployment.DeepCopy()
	//deployment.Spec.Replicas = int32Ptr(0)
	//deployment.Status.ObservedGeneration = 3
	//status, err := clientSet.AppsV1().Deployments(apiv1.NamespaceDefault).UpdateStatus(context.TODO(), deployment, metav1.UpdateOptions{})
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(status.Status)

	fmt.Println(deployment.Spec.Paused)
	fmt.Println(deployment.DeletionTimestamp)

}

func int32Ptr(i int32) *int32 { return &i }

func rolloutRecreate(clientSet *kubernetes.Clientset) {
	deployment, err := clientSet.AppsV1().Deployments(apiv1.NamespaceDefault).Get(context.TODO(), "my-dep", metav1.GetOptions{})
	if err != nil {
		panic(err)
	}

	_, allOldRSs, newRS, err := deploymentutil.GetAllReplicaSets(deployment, clientSet.AppsV1())
	if err != nil {
		panic(err)
	}

	klog.Infof("new rs %v", newRS.Name)
	for _, rs := range allOldRSs {
		klog.Infof("old rs %v", rs.Name)
	}
	klog.Infof("--------------------------")

	rsCopy := newRS.DeepCopy()
	rsCopy.Spec.Replicas = int32Ptr(3)
	update, err := clientSet.AppsV1().ReplicaSets(apiv1.NamespaceDefault).Update(context.TODO(), rsCopy, metav1.UpdateOptions{})
	if err != nil {
		panic(err)
	}
	klog.Infof("new rs replicas: %d", *update.Spec.Replicas)
}

func PodListExample(clientSet *kubernetes.Clientset) {
	sharedInformerFactory := informers.NewSharedInformerFactory(clientSet, 10*time.Second)
	podInformer := sharedInformerFactory.Core().V1().Pods()
	podLister := podInformer.Lister()

	podInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			pod := obj.(*apiv1.Pod)
			fmt.Println("add pod", pod.Name)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			oldPod := oldObj.(*apiv1.Pod)
			newPod := newObj.(*apiv1.Pod)
			fmt.Println("update pod", oldPod.Name, newPod.Name)
		},
		DeleteFunc: func(obj interface{}) {
			pod := obj.(*apiv1.Pod)
			fmt.Println("delete pod", pod.Name)
		},
	})
	fmt.Println("list all pods")
	pods, err := podLister.Pods("").List(labels.Everything())
	if err != nil {
		panic(err)
	}
	fmt.Println("pod len", len(pods))
	for _, pod := range pods {
		fmt.Printf("---> pod: %v\n", pod.Name)
	}

	stopCh := make(chan struct{})
	defer close(stopCh)
	sharedInformerFactory.Start(stopCh)
	sharedInformerFactory.WaitForCacheSync(stopCh)
	<-stopCh
}

func main() {
	clientSet, err := getClientSet()
	if err != nil {
		panic(err)
	}

	//selectorDeepEqual(clientSet)
	//rolloutRecreate(clientSet)
	PodListExample(clientSet)
}
