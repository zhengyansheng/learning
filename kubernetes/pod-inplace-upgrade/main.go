package main

import (
	"flag"
	"fmt"
	"path/filepath"

	"golang.org/x/net/context"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/klog/v2"
)

var (
	clientSet *kubernetes.Clientset
)

func init() {
	var err error
	clientSet, err = getClientSet()
	if err != nil {
		klog.Fatal(err)
	}
}

func main() {
	deploymentName := "my-deployment"
	namespace := "default"
	//image := "nginx:mainline-alpine3.18-perl" // docker pull nginx:mainline-alpine3.18-perl
	image := "nginx:latest"

	// 1. Get the pod by deployment name
	pods, err := getPodByDeploymentName(deploymentName, namespace)
	if err != nil {
		klog.Fatal(err)
	}
	klog.Infof("get %d pods", len(pods))

	// 2. Upgrade the pods
	for _, pod := range pods {
		newPod := pod.DeepCopy()
		klog.Infof("upgrade pod %s", pod.Name)
		_, err := upgradeInPlace(newPod, namespace, image)
		if err != nil {
			klog.Errorf("upgrade pod %s err: %v", newPod.Name, err)
		}
	}

}

func upgradeInPlace(pod *v1.Pod, namespace, image string) (string, error) {
	patch := fmt.Sprintf(`[{"op": "replace", "path": "/spec/containers/0/image", "value": "%s"}]`, image)
	result, err := clientSet.CoreV1().Pods(namespace).Patch(context.TODO(), pod.Name, types.JSONPatchType, []byte(patch), metav1.PatchOptions{})
	if err != nil {
		return "", err
	}
	return result.Name, nil
}

func getPodByDeploymentName(name, namespace string) (pods []*v1.Pod, err error) {
	deployment, err := clientSet.AppsV1().Deployments(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return
	}
	klog.Infof("get deployment %s", deployment.Name)

	labelSelector := metav1.FormatLabelSelector(deployment.Spec.Selector)
	podList, err := clientSet.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{LabelSelector: labelSelector})
	if err != nil {
		return
	}
	klog.Infof("get %d pods", len(podList.Items))

	for _, pod := range podList.Items {
		newPod := pod.DeepCopy()
		pods = append(pods, newPod)
	}

	return
}

func getClientSet() (*kubernetes.Clientset, error) {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(homedir.HomeDir(), ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		klog.Fatal(err)
	}

	return kubernetes.NewForConfig(config)
}
