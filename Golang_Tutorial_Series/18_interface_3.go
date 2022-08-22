package main

import "fmt"

/**
 * @description 接口内部由 tuple 表示 (type, value)，所有类型都实现了空接口
 * @author chengzw
 * @since 2022/8/15
 * @link
 */
func describe(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

func main() {
	s := "Hello World"
	describe(s)
	i := 55
	describe(i)
	strt := struct {
		name string
	}{
		name: "Naveen R",
	}
	describe(strt)
}
