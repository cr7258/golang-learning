package main

import "fmt"

/**
 * @description 闭包
 * @author chengzw
 * @since 2022/8/22
 * @link
 */

func appendStr() func(string) string {
	t := "Hello"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}

func main() {
	a := appendStr()
	b := appendStr()
	fmt.Println(a("World"))    // Hello World
	fmt.Println(b("Everyone")) // Hello Everyone

	fmt.Println(a("Gopher")) // Hello World Gopher
	fmt.Println(b("!"))      // Hello Everyone !
}
