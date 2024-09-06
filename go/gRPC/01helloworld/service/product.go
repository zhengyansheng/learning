package service

import (
	"context"
	"google.golang.org/protobuf/types/known/anypb"
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

func (p *productService) getStockById(id int32) int32 {
	log.Printf("id: %v\n", id)
	return id
}
