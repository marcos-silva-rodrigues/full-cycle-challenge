package services

import (
	"context"

	"github.com/full-cycle-challenge/grpc/pb"
)

type ProductServiceGrpc struct {
	pb.UnimplementedProductServiceServer
}

func (p *ProductServiceGrpc) CreateProduct(ctx context.Context, input *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	return &pb.CreateProductResponse{
		Product: &pb.Product{
			Id:          1,
			Name:        input.Name,
			Description: input.Description,
			Price:       input.Price,
		},
	}, nil
}

func (p *ProductServiceGrpc) FindProducts(ctx context.Context, input *pb.FindProductsRequest) (*pb.FindProductsResponse, error) {
	list := make([]*pb.Product, 1)
	return &pb.FindProductsResponse{
		Products: list,
	}, nil
}

func NewProductServiceGrpc() *ProductServiceGrpc {
	return &ProductServiceGrpc{}
}
