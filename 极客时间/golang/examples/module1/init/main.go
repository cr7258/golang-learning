package main

import (
	"fmt"

	_ "github.com/cncamp/golang/examples/module1/init/a"
	_ "github.com/cncamp/golang/examples/module1/init/b"
)

//Init 函数：会在包初始化时运行, 如果被多个包引用也只会执行一次
func init() {
	fmt.Println("main init")
}
func main() {
	fmt.Println("main function")
}
