package main

import (
	"flag"
	"fmt"
	"os"
)

/**
 * @description 从输入参数读取文件路径
 * @author chengzw
 * @since 2022/8/22
 * @link
 */

func main() {
	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()
	contents, err := os.ReadFile(*fptr)
	if err != nil {
		fmt.Println("File reading err", err)
		return
	}
	fmt.Println("Contents of file:", string(contents))
}
