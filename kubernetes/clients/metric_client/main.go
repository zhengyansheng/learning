package main

import (
	"flag"
	"fmt"
	"path/filepath"
	
	"github.com/zhengyansheng/sample-operator/metrics"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	cacheddiscovery "k8s.io/client-go/discovery/cached"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/restmapper"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/controller-manager/pkg/clientbuilder"
	resourceclient "k8s.io/metrics/pkg/client/clientset/versioned/typed/metrics/v1beta1"
	"k8s.io/metrics/pkg/client/custom_metrics"
	"k8s.io/metrics/pkg/client/external_metrics"
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
	
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	
	builder := clientbuilder.SimpleControllerClientBuilder{ClientConfig: config}
	discoveryClientOrDie := builder.DiscoveryClientOrDie("controller-discovery")
	cacheClient := cacheddiscovery.NewMemCacheClient(discoveryClientOrDie)
	restMapper := restmapper.NewDeferredDiscoveryRESTMapper(cacheClient)
	
	availableAPIsGetter := custom_metrics.NewAvailableAPIsGetter(clientSet.DiscoveryClient)
	metricsClient := metrics.NewRESTMetricsClient(
		resourceclient.NewForConfigOrDie(config),
		custom_metrics.NewForConfig(config, restMapper, availableAPIsGetter),
		external_metrics.NewForConfigOrDie(config),
	)
	
	var set labels.Set = map[string]string{"test": "test"}
	selector := labels.SelectorFromSet(set)
	// 获取 metrice cpu
	podMetricsInfo, time, err := metricsClient.GetResourceMetric(v1.ResourceMemory, "default", selector)
	if err != nil {
		return
	}
	fmt.Println("time", time)
	for name, metric := range podMetricsInfo {
		//fmt.Printf("name: %v, metric: %+v\n", name, metric)
		fmt.Printf("name: %v, value: %+v\n", name, metric.Value/1024/1024/1024)
	}
}
