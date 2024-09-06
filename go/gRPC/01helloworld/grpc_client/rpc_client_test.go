package grpc_client

import (
	"context"
	"github.com/zhengyansheng/helloworld/grpc_client/auth"
	"github.com/zhengyansheng/helloworld/service"
	"log"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestClient(t *testing.T) {
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

func TestSteamClient(t *testing.T) {
	rpcAuth := auth.Authentication{User: "admin", Password: "admin"}
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithPerRPCCredentials(rpcAuth))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := service.NewProdServiceClient(conn)

	var counter int

	done := make(chan struct{})
	defer close(done)

	stream, err := client.UpdateProductStockStream(context.Background(), grpc.EmptyCallOption{})
	if err != nil {
		panic(err)
	}

	go func(done chan struct{}) {
		for {
			r := &service.ProductRequest{
				ProdId: 101,
			}

			if err := stream.SendMsg(r); err != nil {
				t.Logf("err: %v", err)
			}
			time.Sleep(time.Second)
			counter++
			if counter > 10 {
				done <- struct{}{}
			}

		}
	}(done)

	<-done
	recv, err := stream.CloseAndRecv()
	if err != nil {
		t.Failed()
	}
	t.Log(recv)
}
