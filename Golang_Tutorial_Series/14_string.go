package main

import "fmt"

/**
 * @description 字符串
 * @author chengzw
 * @since 2022/8/15
 * @link
 */
func printBytes(s string) {
	fmt.Printf("Bytes: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i]) // 十六进制
	}
}

func printChars(s string) {
	fmt.Printf("Characters: ")
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i]) // 打印字符
	}
}

func main() {
	name := "Hello World"
	fmt.Printf("String: %s\n", name)
	printBytes(name)
	fmt.Printf("\n")
	printChars(name)
}
