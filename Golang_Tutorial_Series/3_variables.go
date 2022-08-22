package main

import "fmt"

/**
 * @description 声明变量
 * @author chengzw
 * @since 2022/8/15
 * @link
 */

func main() {
	// 声明初始值时，类型可以省略
	var width, height int = 100, 50
	fmt.Println("width is ", width, ", height is ", height)

	// 同时声明不同类型的变量
	var (
		name = "tom"
		age  = 27
	)
	fmt.Printf("%s is %d years old", name, age)
}
