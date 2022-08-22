package main

import (
	"fmt"
	"time"
)

/**
 * @description
 * @author chengzw
 * @since 2022/8/16
 * @link
 */
func hello() {
	fmt.Println("Hello world goroutine")
}
func main() {
	go hello()
	fmt.Println("main function")
	time.Sleep(1 * time.Second)
}
