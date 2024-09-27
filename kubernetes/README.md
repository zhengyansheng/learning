# Kubernetes

## 基础
- [x] [client-go](./clients)
- [x] [Informer](./informers)
- [x] [Leader election](./leader-election)
- [x] [Pod原地升级](./pod-inplace-upgrade/main.go)
- [x] [cron hpa controller](https://github.com/AliyunContainerService/kubernetes-cronhpa-controller)


## 组件

_**control plane**_  

- etcd
- api-server
- controller-manager
- scheduler


_**node**_

- kube-proxy
- kubelet

_**network**_

- core-dns
- calico

## 资源

- Pod
- Deployment
- StatefulSet
- DaemonSet
- Cronjob
- job

---
- EndPoint/ EndPointSlice
- Service
- Ingress

---
- Configmap
- Secret
- PV/PVC

---
- HPA


## 工具

**部署**
- kind
- kubeadm

**二次开发**
- kubebuilder
- operator-sdk



## argo CD
