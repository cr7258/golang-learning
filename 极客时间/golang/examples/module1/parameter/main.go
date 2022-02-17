package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// 执行时不传递 --name 参数默认值为 world
	name := flag.String("name", "world", "specify the name you want tho say hi")
	flag.Parse()
	fmt.Println("os args is:", os.Args)
	fmt.Println("input parameter is:", *name)
	fullString := fmt.Sprintf("Hello %s from Go\n", *name)
	fmt.Println(fullString)
}
