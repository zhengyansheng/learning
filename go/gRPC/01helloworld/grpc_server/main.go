package main

import (
	"github.com/zhengyansheng/helloworld/grpc_server/middleware"
	"github.com/zhengyansheng/helloworld/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	// 实现拦截器token认证
	rpcServer := grpc.NewServer(grpc.UnaryInterceptor(middleware.AuthMiddleware()))

	productService := service.NewProduct()
	service.RegisterProdServiceServer(rpcServer, productService)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	log.Println("rpc server listen :8080")
	log.Fatal(rpcServer.Serve(l))

}
