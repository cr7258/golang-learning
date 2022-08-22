package main

import "fmt"

/**
 * @description defer 参数确定
 * @author chengzw
 * @since 2022/8/17
 * @link
 */

func printA(a int) {
	fmt.Println("value of a in deferred function", a) // 5
}
func main() {
	a := 5
	defer printA(a) // 执行的时候就确定 a=5 了，不是在实际调用的函数的时候确定
	a = 10
	fmt.Println("value of a before deferred function call", a) // 10

}
