# Kind

[kind](https://kind.sigs.k8s.io/docs/user/quick-start/)

## Install

```bash

curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.22.0/kind-linux-arm64

```

## Use

```bash

$ kind create cluster --config=config/kind-config.yaml --name=my-kind

$ kubectl cluster-info --context kind-my-kind

$ kind get clusters

$ kind  get kubeconfig --name my-kind > my-kind.kubeconfig

$ kind completion zsh > bash_completion_zsh

$ source ./bash_completion_zsh

$ kind delete clusters my-kindx
```


