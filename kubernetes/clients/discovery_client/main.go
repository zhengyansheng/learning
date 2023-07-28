package main

import (
	"fmt"
	"path/filepath"

	"k8s.io/client-go/discovery"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var (
	defaultName = "terraform-example"
)

func getClient() (*discovery.DiscoveryClient, error) {
	// kubeconfig
	kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")
	// config
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}
	// discovery client
	discoveryClient, err := discovery.NewDiscoveryClientForConfig(config)
	if err != nil {
		return nil, err
	}
	return discoveryClient, nil
}

func ShowAllResources() {
	// discovery client
	discoveryClient, err := getClient()
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

func GcControllerExample() {
	discoveryClient, err := getClient()

	preferredResources, err := discoveryClient.ServerPreferredResources()
	if err != nil {
		panic(err)
	}
	//fmt.Println(preferredResources)
	for _, rl := range preferredResources {
		//fmt.Println(rl.GroupVersion)
		for i := range rl.APIResources {
			fmt.Println(rl.GroupVersion, rl.APIResources[i].Name, rl.APIResources[i].Verbs)
		}
		fmt.Println("-----------------")

	}
	//deletableResources := discovery.FilteredBy(discovery.SupportsAllVerbs{Verbs: []string{"delete", "list", "watch"}}, preferredResources)
	//deletableGroupVersionResources := map[schema.GroupVersionResource]struct{}{}
	//
	//for _, rl := range deletableResources {
	//	gv, err := schema.ParseGroupVersion(rl.GroupVersion)
	//	if err != nil {
	//		klog.Warningf("ignoring invalid discovered resource %q: %v", rl.GroupVersion, err)
	//		continue
	//	}
	//	for i := range rl.APIResources {
	//		deletableGroupVersionResources[schema.GroupVersionResource{Group: gv.Group, Version: gv.Version, Resource: rl.APIResources[i].Name}] = struct{}{}
	//	}
	//}
	//
	//for k, _ := range deletableGroupVersionResources {
	//	fmt.Println(k.Group, k.Version, k.Resource)
	//}
}

func main() {
	GcControllerExample()
}
