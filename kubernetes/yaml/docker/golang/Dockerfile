FROM registry.cn-beijing.aliyuncs.com/zhengyansheng/golang:1.23 AS builder
WORKDIR /build
COPY go.* .
COPY . .
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go

FROM registry.cn-beijing.aliyuncs.com/zhengyansheng/alpine:latest
COPY --from=builder /build/main /bin/main
EXPOSE 8080
CMD ["/bin/main"]