# 使用 centos:7 作为基础镜像
FROM centos:7

COPY CentOS-Base.repo /etc/yum.repos.d/CentOS-Base.repo

# 更新 yum 源，并安装 vim、wget 和 bind-utils（提供 dig 工具）
RUN yum clean all && \
    yum makecache &&\
    yum -y install vim wget bind-utils

# 设置容器启动时的默认命令
CMD ["/bin/bash"]
