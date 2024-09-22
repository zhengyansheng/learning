# Kubernetes

- [kubernetes 1.25.x 代码注解](https://github.com/zhengyansheng/kubernetes)
- [client-go](./kubernetes/clients)
- [Informer](./kubernetes/informers)
- [Leader election](kubernetes/leader-election)
- [Pod原地升级](./kubernetes/pod-inplace-upgrade/main.go)
- [自定义控制器cron hpa controller](https://github.com/AliyunContainerService/kubernetes-cronhpa-controller)


## 自定义调度器的多种方式

- configure with KubeSchedulerConfiguration
- add Plugins of Scheduling Framework
- add Extenders
- etc...

`https://github.com/kubernetes-sigs/kube-scheduler-simulator`

## Patch Type

- JSONPatchType
- MergePatchType
- StrategicMergePatchType
- ApplyPatchType

### JSONPatchType

```go
// Define the JSON Patch data
patchData := []byte(`[
    {"op": "replace", "path": "/spec/replicas", "value": 3}
]`)

// Apply the JSON Patch using the Patch method
patchedDeployment, err := clientset.AppsV1().Deployments(namespace).Patch(
    context.TODO(),
    deploymentName,
    metav1.JSONPatchType, // Specify the Patch type
    patchData,
    metav1.PatchOptions{},
)
```

### MergePatchType

> Merge Patch 适用于简单的合并场景，其中你只关心要更新的字段，而不需要考虑字段的内部结构。

```go
// Define the Merge Patch data
patchData := []byte(`{
    "spec": {
        "replicas": 3
    }
}`)

// Apply the Merge Patch using the Patch method
patchedDeployment, err := clientset.AppsV1().Deployments(namespace).Patch(
    context.TODO(),
    deploymentName,
    metav1.MergePatchType, // Specify the Patch type
    patchData,
    metav1.PatchOptions{},
)
```

```go
patchData := fmt.Sprintf(`{
  "spec": {
    "template": {
      "spec": {
        "containers": [
          {
			"name": "nginx",
            "image": "nginx:1.8"
          }
        ],
		"affinity":null
      }
    }
  }
}`)
_, err = clientset.AppsV1().Deployments("default").Patch(context.TODO(), "nginx-deployment", types.MergePatchType, []byte(patchData), metav1.PatchOptions{})
```

### StrategicMergePatchType

```go
// Define the Strategic Merge Patch data
	patchData := []byte(`
{
  "spec": {
    "replicas": 3
  }
}
`)

	// Apply the Strategic Merge Patch using the Patch method
	patchedDeployment, err := clientset.AppsV1().Deployments(namespace).Patch(
		context.TODO(),
		deploymentName,
		metav1.StrategicMergePatchType, // Specify the Patch type
		patchData,
		metav1.PatchOptions{},
	)
```

### ApplyPatchType

> Apply Patch 适用于你想要将部分更新应用于资源对象的场景，而不是整个对象。

```go
// Define the Apply Patch data
	patchData := []byte(`
spec:
  replicas: 3
`)

	// Apply the Apply Patch using the Patch method
	patchedDeployment, err := clientset.AppsV1().Deployments(namespace).Patch(
		context.TODO(),
		deploymentName,
		metav1.ApplyPatchType, // Specify the Patch type
		patchData,
		metav1.PatchOptions{},
	)
```

- [Informer 为什么要引入 Resync 机制](https://github.com/cloudnativeto/sig-kubernetes/issues/11)  
