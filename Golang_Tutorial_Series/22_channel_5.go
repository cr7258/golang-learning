package main

import "fmt"

/**
 * @description for range 读取通道
 * @author chengzw
 * @since 2022/8/16
 * @link
 */

func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}
func main() {
	ch := make(chan int)
	go producer(ch)
	// 当通道关闭时，for range 会自动退出
	for v := range ch {
		fmt.Println("Received ", v)
	}
}
