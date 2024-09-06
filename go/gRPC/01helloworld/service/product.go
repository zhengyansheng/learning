package service

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/anypb"
	"io"
	"log"
)

type productService struct {
}

func NewProduct() *productService {
	return &productService{}
}

func (p *productService) GetProductStock(ctx context.Context, in *ProductRequest) (*ProductResponse, error) {
	stock := p.getStockById(in.ProdId)
	u := UserRequest{
		Username: "zhengyansheng",
		Age:      18,
		Email:    "zhengyscn@gmail.com",
	}
	data1 := Content{
		Msg: "hello ......",
	}
	dataContent, _ := anypb.New(&data1)
	return &ProductResponse{ProStock: stock, User: &u, Data: dataContent}, nil
}

func (p *productService) mustEmbedUnimplementedProdServiceServer() {}

func (p *productService) UpdateProductStockStream(stream grpc.ClientStreamingServer[ProductRequest, ProductResponse]) error {
	var counter int
	for {
		recv, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		log.Printf("server recv steam: %v, counter: %d\n", recv.ProdId, counter)
		counter++
		if counter > 10 {
			p := &ProductResponse{ProStock: recv.ProdId}
			if err := stream.SendAndClose(p); err != nil {
				return err
			}
			return nil
		}
	}

	//return status.Errorf(codes.Unimplemented, "method UpdateProductStockStream not implemented")
}

func (p *productService) getStockById(id int32) int32 {
	log.Printf("id: %v\n", id)
	return id
}
