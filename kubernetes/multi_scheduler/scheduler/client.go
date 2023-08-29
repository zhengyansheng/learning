package scheduler

import (
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/klog/v2"
)

func getClientset() (*kubernetes.Clientset, error) {
	// 获取 kubeconfig 文件路径
	kubePath := filepath.Join(homedir.HomeDir(), ".kube", "config")
	klog.Infof("kubeconfig path: %s", kubePath)

	// 判断文件是否存在
	if _, err := os.Stat(kubePath); os.IsNotExist(err) {
		kubeConfig := &kubePath
		config, err := clientcmd.BuildConfigFromFlags("", *kubeConfig)
		if err != nil {
			return nil, err
		}
		return kubernetes.NewForConfig(config)
	} else {
		// 不存在，使用 in-cluster 配置
		config, err := rest.InClusterConfig()
		if err != nil {
			return nil, err
		}
		return kubernetes.NewForConfig(config)
	}
}
