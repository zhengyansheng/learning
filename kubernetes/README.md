# Kubernetes

## 基础
- [x] [client-go](./clients)
- [x] [Informer](./informers)
- [x] [Leader election](./leader-election)
- [x] [Pod原地升级](./pod-inplace-upgrade/main.go)
- [x] [cron hpa controller](https://github.com/AliyunContainerService/kubernetes-cronhpa-controller)

## Dockerfile
- [多阶段构建](./yaml/docker/README.md)

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

**CICD**
- tekton
- argo workflow
- argo CD

**渲染工具**
- kustomize
- helm

## 可观测

- 指标 Metric （prometheus)
- 日志 Log (efk,loki)
- 链路 Trace (jaeger, skywalking)







## argo CD
