package main

import (
	"context"
	"errors"
	"flag"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/client-go/util/retry"
	"k8s.io/klog/v2"
)

const (
	DpName              string = "nginx-deployment"
	LabelName           string = "biz-version"
	InitResourceVersion string = "100"
	ConcurrentQuantity  int    = 10 // 5， 10
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

	klog.Info("开始创建deployment")

	// create deployment
	if err := create(clientSet); err != nil {
		panic(err)
	}

	defer delete(clientSet, DpName)

	klog.Info("在协程中并发更新自定义标签")

	startTime := time.Now().UnixMilli()

	w := sync.WaitGroup{}
	for i := 0; i < ConcurrentQuantity; i++ {
		w.Add(1)
		go func(clientSetA *kubernetes.Clientset, index int) {
			defer w.Done()

			// 方式1
			// 并发修改资源 会存在只有一个修改成功，其它都失败
			//err := updateByGetAndUpdate(clientSetA, DpName, index)

			// 方式2：
			// 冲突时重试
			var retryParam = wait.Backoff{
				Steps:    20,
				Duration: 10 * time.Millisecond,
				Factor:   1.0,
				Jitter:   0.1,
			}

			err := retry.RetryOnConflict(retryParam, func() error {
				return updateByGetAndUpdate(clientSet, DpName, index)
			})

			if err != nil {
				klog.Infof("goroutine-%d update err: %v", index, err)
				return
			}
			klog.Infof("goroutine-%d update success", index)

		}(clientSet, i)
	}

	w.Wait()

	// 再查一下，自定义标签的最终值
	deployment, err := get(clientSet, DpName)
	if err != nil {
		klog.Warningf("查询deployment发生异常: %v", err)
		panic(err)
	}

	klog.Infof("自定义标签的最终值为: %v，耗时%v毫秒", deployment.Labels[LabelName], time.Now().UnixMilli()-startTime)
}

func int32Ptr(i int32) *int32 { return &i }

// 创建deployment
func create(clientSet *kubernetes.Clientset) error {
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:   DpName,
			Labels: map[string]string{LabelName: InitResourceVersion},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": DpName,
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": DpName,
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "web",
							Image: "nginx:1.12",
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}

	// Create Deployment
	klog.Info("Creating deployment...")
	result, err := clientSet.AppsV1().Deployments(apiv1.NamespaceDefault).Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		return err
	}

	klog.Infof("Created deployment %q", result.GetObjectMeta().GetName())

	return nil
}

// 按照名称查找deployment
func get(clientSet *kubernetes.Clientset, name string) (*v1.Deployment, error) {
	deployment, err := clientSet.AppsV1().Deployments(apiv1.NamespaceDefault).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return deployment, nil
}

// 按照名称删除
func delete(clientSet *kubernetes.Clientset, name string) error {
	deletePolicy := metav1.DeletePropagationBackground
	if err := clientSet.AppsV1().Deployments(apiv1.NamespaceDefault).Delete(context.TODO(), name, metav1.DeleteOptions{PropagationPolicy: &deletePolicy}); err != nil {
		return err
	}
	return nil
}

// 查询指定名称的deployment对象，得到其名为biz-version的label，加一后保存
func updateByGetAndUpdate(clientSet *kubernetes.Clientset, name string, index int) error {
	deployment, err := clientSet.AppsV1().Deployments(apiv1.NamespaceDefault).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return err
	}

	// 取出当前值
	currentVal, ok := deployment.Labels[LabelName]

	if !ok {
		return errors.New("未取得自定义标签")
	}

	// 将字符串类型转为int型
	val, err := strconv.Atoi(currentVal)
	if err != nil {
		klog.Warning("取得了无效的标签，重新赋初值")
		currentVal = "101"
	}

	// 将int型的label加一，再转为字符串
	deployment.Labels[LabelName] = strconv.Itoa(val + 1)

	oldResourceVersion := deployment.ObjectMeta.ResourceVersion
	oldLabelValue := deployment.Labels[LabelName]
	klog.Infof("goroutine-%d, current label: %v, resource version: %v", index, oldLabelValue, oldResourceVersion)
	updateDeployment, err := clientSet.AppsV1().Deployments(apiv1.NamespaceDefault).Update(context.TODO(), deployment, metav1.UpdateOptions{})
	if err != nil {
		return err
	}
	newResourceVersion := updateDeployment.ObjectMeta.ResourceVersion
	newLabelValue := updateDeployment.ObjectMeta.Labels[LabelName]
	klog.Infof("---> [success] goroutine-%d, current label: %v, resource version old: %v, new: %v", index, newLabelValue, oldResourceVersion, newResourceVersion)
	return nil
}

func updateByRetryOnConflict(clientSet *kubernetes.Clientset, name string, index int) error {
	return retry.RetryOnConflict(retry.DefaultRetry, func() error {
		return updateByGetAndUpdate(clientSet, name, index)
	})
}
