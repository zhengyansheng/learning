package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/metadata"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var (
	defaultName = "terraform-example"
)

func getMetadataClient() metadata.Interface {
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
	// metadata Client
	metadataClient, err := metadata.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	return metadataClient
}

func GetRs() {
	metadataClient := getMetadataClient()

	// 通过metadata找到资源对应的owner
	resource := schema.GroupVersionResource{Group: "apps", Version: "v1", Resource: "replicasets"}
	owner, err := metadataClient.Resource(resource).Namespace("default").Get(context.TODO(), "nginx-dep2-695468c58d", metav1.GetOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("owner: %+v\n", owner)
	fmt.Printf("owner references: %+v\n", owner.GetOwnerReferences())
	ownerAccessor, err := meta.Accessor(owner)
	fmt.Println(ownerAccessor.GetDeletionTimestamp())

}

func DeletePod() {
	metadataClient := getMetadataClient()

	// 通过metadata找到资源对应的owner
	resource := schema.GroupVersionResource{Group: "", Version: "v1", Resource: "pods"}
	err := metadataClient.Resource(resource).Namespace("default").Delete(context.TODO(), "nginx2", metav1.DeleteOptions{})
	if err != nil {
		panic(err)
	}
}

func main() {
	DeletePod()
}
