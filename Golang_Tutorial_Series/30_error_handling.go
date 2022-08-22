package main

import (
	"fmt"
	"os"
)

/**
 * @description 错误处理
 * @author chengzw
 * @since 2022/8/17
 * @link
 */
func main() {
	f, err := os.Open("/test.txt")
	if err != nil {
		if pErr, ok := err.(*os.PathError); ok {
			fmt.Println("Failed to open file at path", pErr.Path)
			return
		}
		fmt.Println("Generic error", err)
	}
	fmt.Println(f.Name(), "opened successfully")
}
