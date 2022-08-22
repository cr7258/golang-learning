package main

import (
	"fmt"
)

/**
 * @description 关闭缓冲通道
 * @author chengzw
 * @since 2022/8/16
 * @link
 */
func main() {
	ch := make(chan int, 5)
	ch <- 5
	ch <- 6
	close(ch) // channel 里的数据需要全部读完才会被关闭
	n, open := <-ch
	fmt.Printf("Received: %d, open: %t\n", n, open)
	n, open = <-ch
	fmt.Printf("Received: %d, open: %t\n", n, open)
	n, open = <-ch
	fmt.Printf("Received: %d, open: %t\n", n, open)
}