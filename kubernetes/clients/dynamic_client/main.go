package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/json"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var (
	defaultName = "terraform-example"
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
	// Dynamic Client
	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		return
	}
	gvr := schema.GroupVersionResource{Group: "apps", Version: "v1", Resource: "deployments"}
	unstructured, err := dynamicClient.Resource(gvr).Namespace(apiv1.NamespaceDefault).Get(context.TODO(), defaultName, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		fmt.Printf("%v not found\n", defaultName)
		return
	}
	if err != nil {
		panic(err)
	}

	// 序列化 方式1
	var ds appsv1.Deployment
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(unstructured.Object, &ds)
	if err != nil {
		panic(err)
	}
	fmt.Printf("方式1 name: %v, namespace: %v\n", ds.Name, ds.Namespace)

	// 序列化 方式2
	var ds2 appsv1.Deployment
	bs, err := json.Marshal(unstructured)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(bs, &ds2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("方式2 name: %v, namespace: %v\n", ds2.Name, ds2.Namespace)

}
