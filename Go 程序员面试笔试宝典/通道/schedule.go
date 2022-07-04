package main

import (
	"fmt"
	"time"
)

/**
 * @description channel 实现定时任务
 * @author chengzw
 * @since 2022/7/1
 * @link https://golang.design/go-questions/channel/application/
 */

func main() {
	ticker := time.Tick(3 * time.Second)

	for {
		select {
		case <-ticker:
			// 执行定时任务
			fmt.Println("执行 3s 定时任务")
		}
	}
}
