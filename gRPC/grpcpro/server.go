package main

import (
	"google.golang.org/grpc"
	"grpcpro/src/pbfiles"
	"grpcpro/src/service"
	"log"
	"net"
)

func main() {
	// 使用 grpc
	myserver := grpc.NewServer()
	pbfiles.RegisterProdServiceServer(myserver, service.NewProdService())
	lis, _ := net.Listen("tcp", ":8080")
	if err := myserver.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
