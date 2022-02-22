package main

import "fmt"

func main() {
	answer := 42
	fmt.Println(&answer) // 内存地址
	fmt.Printf("%T\n", answer)

	address := &answer
	fmt.Println(*address) // 解引用
	fmt.Printf("%T\n", address)
}
