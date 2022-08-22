package main

import "fmt"

/**
 * @description 类型断言
 * @author chengzw
 * @since 2022/8/15
 * @link
 */
func assert(i interface{}) {
	v, ok := i.(int)
	fmt.Println(v, ok)
}
func main() {
	var s interface{} = 56
	assert(s)
	var i interface{} = "Steven Paul"
	assert(i)
}
