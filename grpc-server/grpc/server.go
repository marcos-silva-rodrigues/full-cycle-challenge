package server

import (
	"fmt"
	"log"
	"net"

	"github.com/full-cycle-challenge/grpc/pb"
	"github.com/full-cycle-challenge/grpc/services"
	"github.com/full-cycle-challenge/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartGrpcServer(port int, productUseCase usecase.ProductUseCase) {
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	productServiceGrpc := services.NewProductServiceGrpc(productUseCase)
	pb.RegisterProductServiceServer(grpcServer, productServiceGrpc)

	address := fmt.Sprintf("0.0.0.0:%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot start grpc server", err)
	}

	log.Printf("gRPC server has been started on port %d", port)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start grpc server", err)
	}
}
