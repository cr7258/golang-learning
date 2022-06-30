package main

import "fmt"

func main() {
	var ch = make(chan int, 10)
	for i := 0; i < 10; i++ {
		select { // 随机选择一个
		case ch <- i:
		case v := <-ch:
			fmt.Println(v)
		}
	}
}
