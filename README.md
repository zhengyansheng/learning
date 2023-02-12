## kubernetes


### 大纲
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


### 常见问题

- [Informer 为什么要引入 Resync 机制]