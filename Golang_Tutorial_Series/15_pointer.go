package main

import "fmt"

/**
 * @description 指针
 * @author chengzw
 * @since 2022/8/15
 * @link
 */

func main() {
	b := 255
	var a *int = &b // & 操作符用于获取变量的地址
	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("address of b is", a)

	// 使用 new 函数创建指针
	size := new(int)
	fmt.Printf("Size value is %d, type is %T, address is %v\n", *size, size, size)
	*size = 85
	fmt.Println("New size value is", *size)
}
