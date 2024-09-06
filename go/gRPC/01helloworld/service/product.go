package service

import (
	"context"
	"log"
)

type productService struct {
}

func NewProduct() *productService {
	return &productService{}
}

func (p *productService) GetProductStock(ctx context.Context, in *ProductRequest) (*ProductResponse, error) {
	stock := p.getStockById(in.ProdId)
	return &ProductResponse{ProStock: stock}, nil
}

func (p *productService) mustEmbedUnimplementedProdServiceServer() {}

func (p *productService) getStockById(id int32) int32 {
	log.Printf("id: %v\n", id)
	return id
}
