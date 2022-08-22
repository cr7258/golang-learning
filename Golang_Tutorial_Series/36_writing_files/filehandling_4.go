package main

import (
	"fmt"
	"os"
)

/**
 * @description 往文件中追加内容
 * @author chengzw
 * @since 2022/8/22
 * @link
 */
func main() {
	f, err := os.OpenFile("test.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	newLine := "File handling is easy."
	_, err = fmt.Fprintln(f, newLine)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file written successfully")

}
