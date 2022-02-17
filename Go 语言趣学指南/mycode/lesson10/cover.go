package main

import (
	"fmt"
	"strconv"
)

func main() {

	// 两种转换字符串的方式
	num1 := 10
	str1 := strconv.Itoa(num1)
	fmt.Println("converd: ", str1)

	num2 := 10
	str2 := fmt.Sprintf("%v", num2) // Sprintf 函数会返回格式化之后的 string 而不是打印它
	fmt.Println("converd: ", str2)
}
