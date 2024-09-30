# Kind

[kind文档](https://kind.sigs.k8s.io/docs/user/quick-start/)

## Install

```bash
// 最新版本有时创建集群时会报错
# curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.24.0/kind-linux-amd64
or
# wget https://github.com/kubernetes-sigs/kind/releases/download/v0.20.0/kind-linux-amd64

# chmod +x ./kind
# mv ./kind /usr/local/bin/kind

# kind --help
```


```bash
# wget https://storage.googleapis.com/kubernetes-release/release/v1.27.3/bin/linux/amd64/kubectl
```

**版本**
```bash
# kind version
kind v0.20.0 go1.20.4 linux/amd64

# docker version
Client: Docker Engine - Community
 Version:           26.1.4
```

## 升级内核

from _3.10.0_ to _5.4.225_


## 操作

### 创建集群
```bash
# kind create cluster --config=config/kind-config.yaml --name=my-kind
# kind create cluster --config=config/kind-config.yaml --name=my-kind --image kindest/node:latest

# kubectl cluster-info --context kind-my-kind

# kind get clusters

# kind get kubeconfig --name my-kind > my-kind.kubeconfig
```

### 命令快速补全
```bash
# kind completion zsh > bash_completion_zsh

# source ./bash_completion_zsh
```


### 删除集群
```bash
# kind get clusters

# kind delete clusters my-kind
```


