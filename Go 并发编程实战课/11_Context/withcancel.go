package main

import (
	"context"
	"fmt"
	"time"
)

/**
https://geektutu.com/post/quick-go-context.html#1-%E4%B8%BA%E4%BB%80%E4%B9%88%E9%9C%80%E8%A6%81-Context
 */

func reqTask(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop", name)
			return
		default:
			fmt.Println(name, "send request")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	// 控制多个协程
	go reqTask(ctx, "worker1")
	go reqTask(ctx, "worker2")
	time.Sleep(3 * time.Second)
	cancel() // 通知子协程退出
	time.Sleep(3 * time.Second)
}
