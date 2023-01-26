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
	
	ds, err := clientSet.AppsV1().Deployments(apiv1.NamespaceDefault).Get(context.TODO(), defaultName, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		fmt.Printf("%v not found\n", defaultName)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Println(ds.Name, ds.Namespace)
}
