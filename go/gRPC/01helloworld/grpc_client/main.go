package main

import (
	"context"
	"github.com/zhengyansheng/helloworld/grpc_client/auth"
	"github.com/zhengyansheng/helloworld/service"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	rpcAuth := auth.Authentication{User: "admin", Password: "admin"}
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithPerRPCCredentials(rpcAuth))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := service.NewProdServiceClient(conn)

	r := &service.ProductRequest{
		ProdId: 101,
	}

	stock, err := client.GetProductStock(context.Background(), r, grpc.EmptyCallOption{})
	if err != nil {
		panic(err)
	}
	log.Printf("stock: %v", stock)
}
