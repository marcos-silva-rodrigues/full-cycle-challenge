package main

import (
	"fmt"

	grpc "github.com/full-cycle-challenge/grpc"
)

func main() {
	fmt.Println("Coding challenge ...")
	grpc.StartGrpcServer(50051)
}
