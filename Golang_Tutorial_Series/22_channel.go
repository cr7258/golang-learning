package main

import "fmt"

/**
 * @description
 * @author chengzw
 * @since 2022/8/16
 * @link
 */

// 只能发送数据到 channel
func hello(done chan<- bool) {
	fmt.Println("Hello World Goroutine")
	done <- true
}

func main() {
	done := make(chan bool)
	go hello(done)
	<-done
	fmt.Println("main function")
}
