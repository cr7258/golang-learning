package main

import "fmt"

/**
 * @description 闭包是匿名函数的一种特殊情况。闭包是匿名函数，可以访问在函数体之外定义的变量。
 * @author chengzw
 * @since 2022/8/22
 * @link
 */

func main() {
	a := 5
	func() {
		fmt.Println("a =", a)
	}()
}
