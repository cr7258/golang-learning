package main

import "fmt"

/**
 * @description 条件表达式
 * @author chengzw
 * @since 2022/8/15
 * @link
 */

func main() {
	//if assignment-statement; condition
	if num := 10; num%2 == 0 { // 检查 num 是否是偶数
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}
}
