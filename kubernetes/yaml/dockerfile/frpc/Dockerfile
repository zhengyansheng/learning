FROM alpine:latest

# 安装必要的工具
RUN apk add --no-cache wget

# 下载 frpc
ENV FRPC_VERSION=0.49.0  # 根据需要更改版本号
RUN wget https://github.com/fatedier/frp/releases/download/v${FRPC_VERSION}/frp_${FRPC_VERSION}_linux_amd64.tar.gz && \
    tar -xzf frp_${FRPC_VERSION}_linux_amd64.tar.gz && \
    mv frp_*/frpc /usr/local/bin/ && \
    rm -rf frp_${FRPC_VERSION}_linux_amd64.tar.gz frp_* 

# 添加 frpc 配置文件
COPY frpc.ini /etc/frp/frpc.ini

# 设置工作目录
WORKDIR /etc/frp

# 启动 frpc
CMD ["frpc", "-c", "/etc/frp/frpc.ini"]

