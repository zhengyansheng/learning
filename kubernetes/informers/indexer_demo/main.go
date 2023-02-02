package main

import (
	"fmt"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/cache"
)

const (
	NamespaceIndexName = "namespace"
	NodeNameIndexName  = "node"
)

// MetaNodeIndexFunc is a default index function that indexes based on an object's namespace
func MetaNodeIndexFunc(obj interface{}) ([]string, error) {
	meta, ok := obj.(*v1.Pod)
	if !ok {
		return []string{}, fmt.Errorf("object has no meta: %v", "Pod")
	}
	return []string{meta.Spec.NodeName}, nil
}

func main() {
	indexer := cache.NewIndexer(cache.MetaNamespaceKeyFunc, cache.Indexers{
		NamespaceIndexName: cache.MetaNamespaceIndexFunc,
		NodeNameIndexName:  MetaNodeIndexFunc,
	})

	pod1 := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "pod1",
			Namespace: "default",
		},
		Spec: v1.PodSpec{NodeName: "node1"},
	}

	pod2 := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "pod2",
			Namespace: "default",
		},
		Spec: v1.PodSpec{NodeName: "node1"},
	}

	pod3 := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "pod3",
			Namespace: "kube-system",
		},
		Spec: v1.PodSpec{NodeName: "node2"},
	}

	_ = indexer.Add(pod1)
	_ = indexer.Add(pod2)
	_ = indexer.Add(pod3)

	fmt.Println("当前所有object key", indexer.ListKeys())             // ["default/pod1", "default/pod2", "kube-system/pod3"]
	fmt.Println(indexer.IndexKeys(NamespaceIndexName, "default")) // ["default/pod1", "default/pod2"]

	fmt.Println("==============ByIndex============")
	fmt.Printf("%s 索引器中生成的所有索引 %v\n", NodeNameIndexName, indexer.ListIndexFuncValues(NodeNameIndexName))

	pods, err := indexer.ByIndex(NodeNameIndexName, "node1")
	if err != nil {
		panic(err)
	}
	for _, pod := range pods {
		fmt.Println(pod.(*v1.Pod).Name)
	}
}

// Indexers: {
//		namespace: NamespaceIndexFunc,
//		nodeName:  NodeNameIndexFunc
//}

// Indices: {
//	  namespace: {
//	 	default: set["default/pod1","default/pod2"],
//		kube-system: set["kube-system/pod3"]
//	  },
//	  nodeName: {
//	 	node1: set["default/pod4","kube-system/pod5"],
//		node2: set["ingress-nginx/pod6","default/pod7"],
//	  }
//}

//Index: {
//	default: set["default/pod1","default/pod2"],
//	kube-system: set["kube-system/pod3"]
//}
