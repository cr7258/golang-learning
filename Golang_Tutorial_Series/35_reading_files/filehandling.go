package main

import (
	"fmt"
	"os"
)

/**
 * @description 读取文件
 * @author chengzw
 * @since 2022/8/22
 * @link
 */

func main() {
	contents, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Println("File reading err", err)
		return
	}
	fmt.Println("Contents of file:", string(contents))
}
