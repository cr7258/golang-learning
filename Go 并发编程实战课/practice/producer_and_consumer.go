package main

import (
	"fmt"
)

/**
 * @description 生产者和消费者，无缓冲区
 * @author chengzw
 * @since 2022/7/27
 * @link
 */

func produce(ch chan<- int) {
	i := 0
	for {
		ch <- i
		fmt.Println("Produce: ", i)
		i++
	}
}

func consume(ch <-chan int) {
	for {
		v := <-ch
		fmt.Println("Consume: ", v)
	}
}

func main() {
	ch := make(chan int)
	go produce(ch)
	go consume(ch)
	select {}
}
