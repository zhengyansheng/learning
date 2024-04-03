## 构建镜像

```bash
# docker build -t my_sigterm_container:v1.2 .
```

## 运行容器

```bash
# docker run -itd my_sigterm_container:v1.2
```

## 发送 SIGTERM 信号

```bash
// 发送 SIGTERM 信号，等待 30 秒
# docker stop -s SIGTERM -t 30 my_sigterm_container  

备注
SIGTERM 信号为15
SIGKILL 信号为9
```



