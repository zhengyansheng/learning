package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
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

func main() {
	clientSet, err := getClientSet()
	if err != nil {
		panic(err)
	}

	list(clientSet)
}
