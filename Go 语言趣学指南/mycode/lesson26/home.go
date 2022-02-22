package main

import "fmt"

func main() {
	canada := "Canana"

	// 将 * 放在类型前面表示要声明的指针类型
	var home *string
	fmt.Printf("home is a %T\n", home) //*string

	home = &canada
	// 将 * 放在变量前面表示解引用变量指向的值
	fmt.Printf(*home) // Canada
}
