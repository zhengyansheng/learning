# istio sample

- Traffic Management 流量管理
- Security 安全
- Policy Enforcement 策略执行
- Observability 可观测


## Traffic Management 流量管理

### v1

VirtualService直接路由到Service，然后Service进行路由转发到后端Pod

### v2

VirtualService直接路由到DestinationRule，然后DestinationRule进行分组路由


### v3

按权重进行流量分发

### v4

故障注入
- 延迟
- 中止

### v5

mirror 流量镜像


### v6

Circuit Breaking 熔断器 (没有生效）

### v7

Timeout 超时


