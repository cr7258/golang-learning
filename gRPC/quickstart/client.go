package main

import (
	"52lu/go-study-example/grpc/server/hello"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	// 建立链接
	conn, err := grpc.Dial("127.0.0.1:1234", grpc.WithInsecure())
	if err != nil {
		fmt.Println("Dial Error ", err)
		return
	}
	// 延迟关闭链接
	defer conn.Close()
	// 实例化客户端
	client := hello.NewUserServiceClient(conn)
	// 发起请求
	reply, err := client.Say(context.TODO(), &hello.Request{Name: "张三"})
	if err != nil {
		return
	}
	fmt.Println("返回:", reply.Result)
}
