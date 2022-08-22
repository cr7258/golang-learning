package main

import "fmt"

/**
 * @description 关闭通道
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
	for {
		v, ok := <-ch
		if !ok {
			fmt.Println("Channel is closed")
			break
		}
		fmt.Println("Received ", v, ok)
	}
}
