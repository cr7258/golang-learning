package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpcpro/src/pbfiles"
	"log"
)

func main() {
	client, err := grpc.DialContext(context.Background(),
		"localhost:8080",
		grpc.WithInsecure())

	if err != nil {
		log.Fatal()
	}

	req := &pbfiles.ProdRequest{ProdId: 123}
	rsp := &pbfiles.ProdResponse{}

	_ = client.Invoke(context.Background(),
		"/ProdService/GetProd", req, rsp)

	fmt.Println(rsp.Result)
}
