package main

import "fmt"

/**
 * @description 带有缓冲区的 channel
 * @author chengzw
 * @since 2022/8/16
 * @link
 */

func main() {
	ch := make(chan string, 1)
	ch <- "naveen"
	ch <- "paul"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
