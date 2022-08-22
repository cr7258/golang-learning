package main

import "fmt"

/**
 * @description 数组
 * @author chengzw
 * @since 2022/8/15
 * @link
 */
func main() {
	a := [...]string{"USA", "China", "India", "Germany", "France"}
	b := a             // a copy of a is assigned to b
	b[0] = "Singapore" // array 是值类型，不是引用类型，改变 array b 的元素不会影响 array a
	fmt.Println("a is ", a)
	fmt.Println("b is ", b)
}
