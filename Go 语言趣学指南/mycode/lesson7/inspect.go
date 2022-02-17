package main

import "fmt"

func main() {
	year := 2018
	fmt.Printf("Type %T for %v\n", year, year)

	// 使用 [1] 可以复用第一个格式化的变量值
	fmt.Printf("Type %T for %[1]v\n", year)
}
