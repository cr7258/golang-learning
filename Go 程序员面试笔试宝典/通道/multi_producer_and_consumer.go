package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
 * @description 多生产者和消费者
 * @author chengzw
 * @since 2022/7/1
 * @link
 */

func main() {
	taskCh := make(chan int, 100)
	go producer(taskCh)
	go consumer(taskCh)

	time.Sleep(time.Hour)
}

// 只能接收 channel 中的数据
// 消费者
func consumer(taskCh <-chan int) {
	for i := 0; i < 5; i++ {
		go func(id int) {
			for {
				task := <-taskCh
				fmt.Printf("consume task: %d by consumer %d\n", task, id)
				time.Sleep(time.Second)
			}
		}(i)
	}
}

func producer(taskCh chan<- int) {
	for i := 0; i < 5; i++ {
		go func(id int) {
			for {
				task := rand.Intn(100)
				taskCh <- task
				fmt.Printf("produce task: %d by producer %d\n", task, id)
				time.Sleep(time.Second)
			}
		}(i)
	}
}
