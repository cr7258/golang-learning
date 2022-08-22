package main

import "fmt"

/**
 * @description 变量类型
 * @author chengzw
 * @since 2022/8/15
 * @link
 */

func main() {
	a := 5.67
	fmt.Printf("type of a %T\n", a) // 打印变量类型

	// 类型转换
	i := 55   // int
	j := 67.8 // float64
	sum := float64(i) + j
	fmt.Println(sum)
}
