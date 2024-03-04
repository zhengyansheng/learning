# istio sample

## v1

VirtualService直接路由到Service，然后Service进行路由转发到后端Pod

## v2

VirtualService直接路由到DestinationRule，然后DestinationRule进行分组路由


## v3

按权重进行流量分发

## v4

故障注入
- 延迟
- 中止

## v5
