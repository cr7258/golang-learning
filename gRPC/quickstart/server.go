package main

import (
	"52lu/go-study-example/grpc/server/hello"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

func main() {
	// 创建grpc服务
	grpcServer := grpc.NewServer()
	// 注册服务
	hello.RegisterUserServiceServer(grpcServer, new(hello.UnimplementedUserServiceServer))
	// 监听端口
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("服务启动失败", err)
		return
	}
	grpcServer.Serve(listen)
}
