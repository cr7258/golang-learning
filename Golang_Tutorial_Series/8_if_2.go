package main

import "fmt"

/**
 * @description 避免不必要的分支和代码缩进，尽早返回
 * @author chengzw
 * @since 2022/8/15
 * @link
 */

func main() {
	if num := 10; num%2 == 0 { //checks if number is even
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 更好的写法
	num := 10
	if num%2 == 0 {
		fmt.Println(num, "is even")
		return
	}
	fmt.Println(num, "is odd")
}
