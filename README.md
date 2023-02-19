## kubernetes

> [kubernetes代码注解](https://github.com/zhengyansheng/kubernetes)

`client-go 常见客户端`
- ClientSet
- DynamicClient
- DiscoveryClient
- ScaleClient
- MetricClient


`Informer`
- ListAndWatch
- Informer factory
  - gvc
  - ...
- Indexer

`Controller`
- controller(workqueue, indexer, informer)

`operators`
- nginx operator
- elasticweb operator


## Go

`sync`


### 工具
- [反序列化etcd数据 kube-etcd-helper](https://github.com/yamamoto-febc/kube-etcd-helper)
- [json diff](https://jsondiff.com/)

### 常见问题

- [Informer 为什么要引入 Resync 机制](https://github.com/cloudnativeto/sig-kubernetes/issues/11)



