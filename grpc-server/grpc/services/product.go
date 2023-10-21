package services

import (
	"context"

	"github.com/full-cycle-challenge/grpc/pb"
	"github.com/full-cycle-challenge/usecase"
)

type ProductServiceGrpc struct {
	ProductUseCase usecase.ProductUseCase
	pb.UnimplementedProductServiceServer
}

func (p *ProductServiceGrpc) CreateProduct(ctx context.Context, input *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	response, err := p.ProductUseCase.Create(
		input.Name,
		input.Description,
		input.Price,
	)

	if err != nil {
		return nil, err
	}

	return &pb.CreateProductResponse{
		Product: &pb.Product{
			Id:          response.ID,
			Name:        response.Name,
			Description: response.Description,
			Price:       response.Price,
		},
	}, nil
}

func (p *ProductServiceGrpc) FindProducts(ctx context.Context, input *pb.FindProductsRequest) (*pb.FindProductsResponse, error) {
	list, err := p.ProductUseCase.FindAll()
	if err != nil {
		return nil, err
	}

	response := make([]*pb.Product, len(list))
	for i := range list {
		response[i] = &pb.Product{
			Id:          list[i].ID,
			Name:        list[i].Name,
			Description: list[i].Description,
			Price:       list[i].Price,
		}
	}
	return &pb.FindProductsResponse{
		Products: response,
	}, nil
}

func NewProductServiceGrpc(productUseCase usecase.ProductUseCase) *ProductServiceGrpc {
	return &ProductServiceGrpc{
		ProductUseCase: productUseCase,
	}
}
