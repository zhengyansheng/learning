package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
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

	// list
	dl, err := clientSet.AppsV1().Deployments(apiv1.NamespaceDefault).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	if len(dl.Items) != 0 {
		for _, item := range dl.Items {
			fmt.Printf("List deployment: %v", item.Name)
		}
	}

	// watch
	watch, err := clientSet.AppsV1().Deployments(apiv1.NamespaceDefault).Watch(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	defer watch.Stop()

	for {
		select {
		case e := <-watch.ResultChan():
			switch e.Object.(type) {
			case *appsv1.Deployment:
				deploy, ok := e.Object.(*appsv1.Deployment)
				if !ok {
					return
				}
				fmt.Printf("Watch 到 %s 变化,EventType: %s\n", deploy.Name, e.Type)
			}
		}
	}

}
