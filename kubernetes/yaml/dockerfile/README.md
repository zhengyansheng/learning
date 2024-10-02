# Dockerfile

- Vue
- Go
- Java

## 构建
```bash
// Go
# docker build -t echo-httpserver:v1.0.0 -f golang/Dockerfile golang

// Vue
# docker build -t vue-nginx:v1.0.0 -f vue/Dockerfile vue

// Java
# docker build -t vue-nginx:v1.0.0 -f spring-boot/Dockerfile spring-boot
```

_参数_
- -f 指定Dockerfile文件 golang参数指定目录的上下文路径    

## 部署
```bash
# docker run -p 8080:8080 echo-httpserver:v1.0.0

# docker run -p 8080:80 vue-nginx:v1.0.0

# docker run -p 8080:80 spring-boot:v1.0.0
```
