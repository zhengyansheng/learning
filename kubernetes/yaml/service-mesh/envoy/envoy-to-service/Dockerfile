# 使用 Envoy 官方镜像作为基础镜像
FROM envoyproxy/envoy:v1.22.2

# 将本地的 envoy.yaml 文件复制到容器中
COPY envoy.yaml /etc/envoy/envoy.yaml

# 设置容器启动时执行的命令
CMD ["envoy", "-c", "/etc/envoy/envoy.yaml"]

