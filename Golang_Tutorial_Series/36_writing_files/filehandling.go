package main

import (
	"fmt"
	"os"
)

/**
 * @description 往文件中写入字符串
 * @author chengzw
 * @since 2022/8/22
 * @link
 */
func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString("Hello World")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
