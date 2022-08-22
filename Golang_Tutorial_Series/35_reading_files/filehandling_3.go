package main

import (
	"fmt"
)

/**
 * @description
 * @author chengzw
 * @since 2022/8/22
 * @link
 */

import (
	_ "embed"
)

//go:embed test.txt
var contents []byte

func main() {
	fmt.Println("Contents of file:", string(contents))
}
