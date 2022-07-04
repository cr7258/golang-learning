package main

import (
	"fmt"
	"time"
)

/**
 * @description 生产者和消费者
 * @author chengzw
 * @since 2022/7/1
 * @link
 */

func main() {
	taskCh := make(chan int, 100)
	go consumer(taskCh)

	// 生产 10 条数据
	for i := 0; i < 10; i++ {
		taskCh <- i
	}

	time.Sleep(time.Hour)
}

// 只能接收 channel 中的数据
// 消费者
func consumer(taskCh <-chan int) {
	for i := 0; i < 5; i++ {
		go func(id int) {
			for {
				task := <-taskCh
				fmt.Printf("finish task: %d by worker %d\n", task, id)
				time.Sleep(time.Second)
			}
		}(i)
	}
}
