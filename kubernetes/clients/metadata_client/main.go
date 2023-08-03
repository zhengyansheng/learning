package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"path/filepath"

	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/metadata"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// PrintStruct 写一个函数，传递一个参数，实现golang struct格式化对其输出
func PrintStruct(obj interface{}) {
	bs, _ := json.Marshal(obj)
	var out bytes.Buffer
	json.Indent(&out, bs, "", "\t")
	fmt.Printf("%+v\n", out.String())
}

func getClient() metadata.Interface {
	kubeConfig := filepath.Join(homedir.HomeDir(), ".kube", "config")
	kubeconfig := &kubeConfig
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
	metadataClient := getClient()

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

func GetPod() {
	metadataClient := getClient()

	// 通过metadata找到资源对应的owner
	resource := schema.GroupVersionResource{Group: "", Version: "v1", Resource: "pods"}
	var subResources []string
	var podName = "nginx2"
	partialObjectMetadata, err := metadataClient.Resource(resource).Namespace("default").Get(context.TODO(), podName, metav1.GetOptions{}, subResources...)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%+v\n", partialObjectMetadata)
	PrintStruct(partialObjectMetadata)
}

func DeletePod() {
	metadataClient := getClient()
	var podName = "nginx2"

	// 通过metadata找到资源对应的owner
	resource := schema.GroupVersionResource{Group: "", Version: "v1", Resource: "pods"}
	err := metadataClient.Resource(resource).Namespace("default").Delete(context.TODO(), podName, metav1.DeleteOptions{})
	if err != nil {
		panic(err)
	}
}

func DeleteCm() {
	metadataClient := getClient()

	// 通过metadata找到资源对应的owner
	resource := schema.GroupVersionResource{Group: "", Version: "v1", Resource: "configmaps"}

	policy := metav1.DeletePropagationBackground
	var uid types.UID = "0d33e93b-1ab2-4622-8f55-2a14f03fcc76"
	preconditions := metav1.Preconditions{UID: &uid}
	deleteOptions := metav1.DeleteOptions{Preconditions: &preconditions, PropagationPolicy: &policy}
	err := metadataClient.Resource(resource).Namespace("default").Delete(context.TODO(), "cm", deleteOptions)
	if err != nil {
		panic(err)
	}
}

func main() {
	//// get replicaset
	//GetPod()
	////GetRs()
	//time.Sleep(time.Second)
	//fmt.Println("------------------")
	//
	//// delete pod from metadata client
	//DeletePod()
	//time.Sleep(time.Second)
	//fmt.Println("------------------")
	//
	//// mark
	//GetPod()
	//fmt.Println("------------------")
	DeleteCm()
}
