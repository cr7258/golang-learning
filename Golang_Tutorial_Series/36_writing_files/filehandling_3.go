package main

import (
	"fmt"
	"os"
)

/**
 * @description 逐行写入字符串
 * @author chengzw
 * @since 2022/8/22
 * @link
 */
func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	d := []string{"Welcome to the world of Go1.", "Go is a compiled language.", "It is easy to learn Go."}
	for _, v := range d {
		fmt.Fprintln(f, v)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file written successfully")

}
