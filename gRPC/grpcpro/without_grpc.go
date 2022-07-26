package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"grpcpro/src/pbfiles"
)

func main() {
	// 不依赖 grpc 使用 protobuf
	// 编码
	prod := &pbfiles.ProdModel{Id: 123, Name: "chengzw"}
	b, err := proto.Marshal(prod)
	fmt.Println(b, err)

	// 解码
	prod2 := &pbfiles.ProdModel{}
	proto.Unmarshal(b, prod2)
	fmt.Println(prod2)
}
