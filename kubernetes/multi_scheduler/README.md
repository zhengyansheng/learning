# 自定义 kubernetes scheduler

## 1. 随机调度器

### Usage

1. 构建镜像
```bash

# docker build -t zhengyscn/random-scheduler:v1.x .

# docker push zhengyscn/random-scheduler:v1.x

```

2. 部署
```bash

# kubectl apply -f sample/rbac.yaml
# kubectl apply -f sample/random-scheduler_deployment.yaml

```

3. 测试
```bash

# kubectl apply -f sample/sleep_deployment.yaml

```




## 参考
- [官网配置多调度器](https://kubernetes.io/zh-cn/docs/tasks/extend-kubernetes/configure-multiple-schedulers/)
- [banzaicloud-random-scheduler](https://github.com/banzaicloud/random-scheduler)