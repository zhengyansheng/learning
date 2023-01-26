package main

import (
	"flag"
	"fmt"
	"path/filepath"
	
	"k8s.io/client-go/discovery"
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
	
	// discovery client
	discoveryClient, err := discovery.NewDiscoveryClientForConfig(config)
	if err != nil {
		panic(err)
	}
	// ServerGroupsAndResources returns the supported resources for all groups and versions.
	apiGroups, apiResourceList, err := discoveryClient.ServerGroupsAndResources()
	if err != nil {
		return
	}
	for i, group := range apiGroups {
		if group.Kind == "" {
			fmt.Printf("i: %d, name: %v, version: %v, apiVersion: %v\n", i, group.Name, group.Versions, group.APIVersion)
			continue
		}
		fmt.Printf("i: %d, name: %v, kind: %v, version: %v, apiVersion: %v\n", i, group.Name, group.Kind, group.Versions, group.APIVersion)
	}
	fmt.Println("============================")
	
	for _, r := range apiResourceList {
		//fmt.Println(i, r.Kind, r.APIVersion, r.APIResources, r.GroupVersion, r.String())
		//fmt.Println(i, r.Kind, r.APIVersion, r.APIResources, r.GroupVersion)
		//fmt.Println(i, r.Kind, r.APIVersion, r.GroupVersion)
		//fmt.Println(i, r.APIVersion, r.GroupVersion)
		for _, resource := range r.APIResources {
			fmt.Println(r.APIVersion, r.GroupVersion, resource.Name, resource.Group, resource.Version, resource.Kind)
		}
	}
}
