package main

import "fmt"

/**
 * @description defer 调用栈
 * @author chengzw
 * @since 2022/8/17
 * @link
 */
func main() {
	name := "Naveen"
	fmt.Printf("Original String: %s\n", string(name))
	fmt.Printf("Reversed String: ")
	for _, v := range []rune(name) {
		defer fmt.Printf("%c", v) // neevaN
	}
}
