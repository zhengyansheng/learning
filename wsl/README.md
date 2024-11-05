# wsl

## 导出
```bash
# docker pull centos:7

# docker export $(docker create centos:7) -o centos-7-rootfs.tar.gz
```

## 导入
```bash
# wsl --import CentOS7 E:\wsl\local\CentOS7 E:\wsl\centos-7-rootfs.tar.gz --version 2
```

## 查看
```bash
# wsl --list --verbose
```

## 启动
```bash
# wsl -d CentOS7
```