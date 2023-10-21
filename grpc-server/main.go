package main

import (
	"fmt"

	"github.com/full-cycle-challenge/database"
	grpc "github.com/full-cycle-challenge/grpc"
	"github.com/full-cycle-challenge/usecase"
)

func main() {
	fmt.Println("Coding challenge ...")

	db := database.ConnectDB()
	productUseCase := usecase.ProductUseCase{
		DB: db,
	}
	grpc.StartGrpcServer(50051, productUseCase)
}
