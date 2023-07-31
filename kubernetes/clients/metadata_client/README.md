# metadata client

通过分析，发现metadata client是一个用于获取kubernetes集群中的资源信息的客户端，它的主要功能是通过调用apiserver的RESTful API来获取资源信息。

核心是获取资源的 metadata 信息，包括： name, namespace, labels, annotations, ownerReferences, finalizers, deletionTimestamp, deletionGracePeriodSeconds, generation, uid, resourceVersion, selfLink, creationTimestamp, initializers, managedFields, generateName, ownerReferences, clus terName, 
clusterUID, controller, blockOwnerDeletion, and orphanDependents. 

## 1. 查询 一个 Pod 的 metadata 资源

```json
{
	"metadata": {
		"name": "nginx",
		"namespace": "default",
		"uid": "5b631df0-8559-4943-931a-43a63ad2c198",
		"resourceVersion": "20985182",
		"creationTimestamp": "2023-07-31T05:39:02Z",
		"labels": {
			"run": "nginx"
		},
		"annotations": {
			"cni.projectcalico.org/containerID": "a0c2612fee2572c438f0e7615994ae587f420e4bb6fa012cca9abb1943657ab4",
			"cni.projectcalico.org/podIP": "10.244.166.171/32",
			"cni.projectcalico.org/podIPs": "10.244.166.171/32"
		},
		"managedFields": [
			{
				"manager": "calico",
				"operation": "Update",
				"apiVersion": "v1",
				"time": "2023-07-31T05:39:02Z",
				"fieldsType": "FieldsV1",
				"fieldsV1": {
					"f:metadata": {
						"f:annotations": {
							".": {},
							"f:cni.projectcalico.org/containerID": {},
							"f:cni.projectcalico.org/podIP": {},
							"f:cni.projectcalico.org/podIPs": {}
						}
					}
				},
				"subresource": "status"
			},
			{
				"manager": "kubectl-run",
				"operation": "Update",
				"apiVersion": "v1",
				"time": "2023-07-31T05:39:02Z",
				"fieldsType": "FieldsV1",
				"fieldsV1": {
					"f:metadata": {
						"f:labels": {
							".": {},
							"f:run": {}
						}
					},
					"f:spec": {
						"f:containers": {
							"k:{\"name\":\"nginx\"}": {
								".": {},
								"f:image": {},
								"f:imagePullPolicy": {},
								"f:name": {},
								"f:resources": {},
								"f:terminationMessagePath": {},
								"f:terminationMessagePolicy": {}
							}
						},
						"f:dnsPolicy": {},
						"f:enableServiceLinks": {},
						"f:restartPolicy": {},
						"f:schedulerName": {},
						"f:securityContext": {},
						"f:terminationGracePeriodSeconds": {}
					}
				}
			},
			{
				"manager": "kubelet",
				"operation": "Update",
				"apiVersion": "v1",
				"time": "2023-07-31T05:39:19Z",
				"fieldsType": "FieldsV1",
				"fieldsV1": {
					"f:status": {
						"f:conditions": {
							"k:{\"type\":\"ContainersReady\"}": {
								".": {},
								"f:lastProbeTime": {},
								"f:lastTransitionTime": {},
								"f:status": {},
								"f:type": {}
							},
							"k:{\"type\":\"Initialized\"}": {
								".": {},
								"f:lastProbeTime": {},
								"f:lastTransitionTime": {},
								"f:status": {},
								"f:type": {}
							},
							"k:{\"type\":\"Ready\"}": {
								".": {},
								"f:lastProbeTime": {},
								"f:lastTransitionTime": {},
								"f:status": {},
								"f:type": {}
							}
						},
						"f:containerStatuses": {},
						"f:hostIP": {},
						"f:phase": {},
						"f:podIP": {},
						"f:podIPs": {
							".": {},
							"k:{\"ip\":\"10.244.166.171\"}": {
								".": {},
								"f:ip": {}
							}
						},
						"f:startTime": {}
					}
				},
				"subresource": "status"
			}
		]
	}
}
```

## 2. 删除 一个 Pod 的 metadata 资源

> 重点看 deletionTimestamp 和 deletionGracePeriodSeconds 2个字段；

```json
{
	"metadata": {
		"name": "nginx",
		"namespace": "default",
		"uid": "5b631df0-8559-4943-931a-43a63ad2c198",
		"resourceVersion": "20985203",
		"creationTimestamp": "2023-07-31T05:39:02Z",
		"deletionTimestamp": "2023-07-31T05:40:06Z", # 重点看这个字段
		"deletionGracePeriodSeconds": 30,            # 重点看这个字段
		"labels": {
			"run": "nginx"
		},
		"annotations": {
			"cni.projectcalico.org/containerID": "a0c2612fee2572c438f0e7615994ae587f420e4bb6fa012cca9abb1943657ab4",
			"cni.projectcalico.org/podIP": "10.244.166.171/32",
			"cni.projectcalico.org/podIPs": "10.244.166.171/32"
		},
		"managedFields": [
			{
				"manager": "calico",
				"operation": "Update",
				"apiVersion": "v1",
				"time": "2023-07-31T05:39:02Z",
				"fieldsType": "FieldsV1",
				"fieldsV1": {
					"f:metadata": {
						"f:annotations": {
							".": {},
							"f:cni.projectcalico.org/containerID": {},
							"f:cni.projectcalico.org/podIP": {},
							"f:cni.projectcalico.org/podIPs": {}
						}
					}
				},
				"subresource": "status"
			},
			{
				"manager": "kubectl-run",
				"operation": "Update",
				"apiVersion": "v1",
				"time": "2023-07-31T05:39:02Z",
				"fieldsType": "FieldsV1",
				"fieldsV1": {
					"f:metadata": {
						"f:labels": {
							".": {},
							"f:run": {}
						}
					},
					"f:spec": {
						"f:containers": {
							"k:{\"name\":\"nginx\"}": {
								".": {},
								"f:image": {},
								"f:imagePullPolicy": {},
								"f:name": {},
								"f:resources": {},
								"f:terminationMessagePath": {},
								"f:terminationMessagePolicy": {}
							}
						},
						"f:dnsPolicy": {},
						"f:enableServiceLinks": {},
						"f:restartPolicy": {},
						"f:schedulerName": {},
						"f:securityContext": {},
						"f:terminationGracePeriodSeconds": {}
					}
				}
			},
			{
				"manager": "kubelet",
				"operation": "Update",
				"apiVersion": "v1",
				"time": "2023-07-31T05:39:19Z",
				"fieldsType": "FieldsV1",
				"fieldsV1": {
					"f:status": {
						"f:conditions": {
							"k:{\"type\":\"ContainersReady\"}": {
								".": {},
								"f:lastProbeTime": {},
								"f:lastTransitionTime": {},
								"f:status": {},
								"f:type": {}
							},
							"k:{\"type\":\"Initialized\"}": {
								".": {},
								"f:lastProbeTime": {},
								"f:lastTransitionTime": {},
								"f:status": {},
								"f:type": {}
							},
							"k:{\"type\":\"Ready\"}": {
								".": {},
								"f:lastProbeTime": {},
								"f:lastTransitionTime": {},
								"f:status": {},
								"f:type": {}
							}
						},
						"f:containerStatuses": {},
						"f:hostIP": {},
						"f:phase": {},
						"f:podIP": {},
						"f:podIPs": {
							".": {},
							"k:{\"ip\":\"10.244.166.171\"}": {
								".": {},
								"f:ip": {}
							}
						},
						"f:startTime": {}
					}
				},
				"subresource": "status"
			}
		]
	}
}
```

`结论`

```text

当 通过 metadata client 删除 Pod 时，本质上就是在metadata属性中添加了 deletionTimestamp 和 deletionGracePeriodSeconds 2个字段；

```