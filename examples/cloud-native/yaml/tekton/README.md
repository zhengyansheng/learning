# Tekton

- task
  - step
- taskrun
- pipeline
- pipelinerun


## 部署
```bash
kubectl apply --filename https://storage.googleapis.com/tekton-releases/pipeline/latest/release.yaml
```

## 查看
```bash
kubectl api-resources | grep tekton
```

```bash
kubectl get pods -n tekton-pipelines
```

## 安装CLI
[tekton-cli](https://github.com/tektoncd/cli/releases)
```bash
brew install tektoncd-cli
```

## 运行
```bash
kubectl apply -f hello.yaml
```

```bash
tkn taskrun describe echo-hello-world-task-run
tkn taskrun logs echo-hello-world-task-run
```